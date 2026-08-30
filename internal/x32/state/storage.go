package state

import (
	"aumix/internal/osc"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type ShowScene struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Movement int    `json:"movement"`
	Measure  int    `json:"measure"`
	X32Scene int    `json:"sceneId"`
}

type Show struct {
	Id           int          `json:"id"`
	Name         string       `json:"name"`
	CurrentScene int          `json:"_"`
	Scenes       []*ShowScene `json:"scenes"`
}

type Config struct {
	MixerIp     string
	MixerPort   int
	Shows       map[int]*Show
	CurrentShow int
	NextShowId  int
	NextSceneId int
	State       X32State
}

type ConfigWrapper struct {
	Config *Config
	lock   sync.Mutex
}

func (c *ConfigWrapper) Save(filePath string) error {
	err := os.MkdirAll("data", 0755)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)

	{
		c.lock.Lock()
		defer c.lock.Unlock()
		err = encoder.Encode(c.Config)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *ConfigWrapper) Lock() *Config {
	c.lock.Lock()
	return c.Config
}

func (c *ConfigWrapper) Unlock() {
	c.lock.Unlock()
}

func (c *ConfigWrapper) resolvePath(path string) (reflect.Value, error) {
	parts := strings.Split(path, "/")
	if parts[0] == "" {
		parts = parts[1:]
	}
	for i, part := range parts {
		parts[i] = strings.ReplaceAll(part, "-", "_")
	}

	state := &c.Config.State
	currentValue := reflect.ValueOf(state).Elem()

	for _, part := range parts {
		nextValue := currentValue.FieldByName("F" + part)
		if !nextValue.IsValid() {
			nextValue = currentValue.FieldByName("Index")
			indexField, _ := currentValue.Type().FieldByName("Index")
			if !nextValue.IsValid() {
				return nextValue, fmt.Errorf("Field \"%s\" does not exist", part)
			}
			index, err := strconv.Atoi(part)
			if err != nil {
				pairString, ok := indexField.Tag.Lookup("pairs")
				if !ok {
					return nextValue, fmt.Errorf("Fragment \"%s\" is not an integer", part)
				}
				split := strings.Split(part, "-")
				index, err = strconv.Atoi(split[0])
				if err != nil {
					return nextValue, fmt.Errorf("Field \"%s\" does not exist", part)
				}
				pair, err := strconv.Atoi(pairString)
				if err != nil {
					return nextValue, fmt.Errorf("Error parsing field")
				}
				index /= pair
			}
			startString, ok := indexField.Tag.Lookup("start")
			if ok {
				start, err := strconv.Atoi(startString)
				if err != nil {
					return nextValue, fmt.Errorf("Error parsing field")
				}
				index -= start
			}
			if index >= nextValue.Len() {
				return nextValue, fmt.Errorf("%d is out of range", index)
			}
			nextValue = nextValue.Index(index)
		}
		currentValue = nextValue
	}

	return currentValue, nil
}

func (c *ConfigWrapper) GetByPath(path string) (any, error) {
	c.Lock()
	defer c.Unlock()

	value, err := c.resolvePath(path)
	return value, err
}

func (c *ConfigWrapper) setValue(currentValue reflect.Value, values []any, index *int) error {
	field, ok := currentValue.Addr().Interface().(X32StateValue)
	if ok {
		if *index >= len(values) {
			return fmt.Errorf("Not enough parameters provided")
		}
		consumed, err := field.Set(values[*index:]...)
		*index += consumed
		return err
	}

	for f, v := range currentValue.Fields() {
		if f.Name == "Index" {
			for i := range f.Type.Len() {
				err := c.setValue(v.Index(i), values, index)
				if err != nil {
					return err
				}
			}
		} else {
			err := c.setValue(v, values, index)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *ConfigWrapper) SetByPath(path string, values []any) error {
	c.Lock()
	defer c.Unlock()

	currentValue, err := c.resolvePath(path)
	if err != nil {
		return err
	}

	index := 0
	return c.setValue(currentValue, values, &index)
}

func (c *ConfigWrapper) SetParametersByPath(path string, values []osc.Parameter) error {
	extracted := make([]any, len(values))
	for index, value := range values {
		switch v := value.(type) {
		case osc.IntParam:
			extracted[index] = int32(v)
		case osc.FloatParam:
			extracted[index] = float32(v)
		case osc.StringParam:
			extracted[index] = string(v)
		default:
			return fmt.Errorf("Invalid parameter type provided")
		}
	}

	return c.SetByPath(path, extracted)
}

func (c *ConfigWrapper) SetFader(index int, level float32) error {
	config := c.Lock()
	defer c.Unlock()

	channels := &config.State.Fch.Index
	if index >= len(channels) {
		return fmt.Errorf("Fader index %d is out of range %d", index, len(channels))
	}
	channels[index].Fmix.Ffader = X32Level(level)

	return nil
}

func LoadX32Config(filePath string) (*ConfigWrapper, error) {
	var configWrapper *ConfigWrapper
	file, err := os.Open(filePath)
	if err != nil {
		return configWrapper, err
	}
	defer file.Close()

	config := &Config{}
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return configWrapper, err
	}
	configWrapper = &ConfigWrapper{
		Config: config,
	}
	return configWrapper, nil
}

func NewX32Config() *ConfigWrapper {
	configWrapper := &ConfigWrapper{
		Config: &Config{
			Shows: make(map[int]*Show),
		},
	}
	return configWrapper
}

func (c *ConfigWrapper) GetShowList() map[int]string {
	config := c.Lock()
	defer c.Unlock()

	shows := make(map[int]string)
	for _, show := range config.Shows {
		shows[show.Id] = show.Name
	}

	return shows
}

func (c *ConfigWrapper) GetCurrentShow() (*Show, bool) {
	show, ok := c.Config.Shows[c.Config.CurrentShow]
	return show, ok
}

func (c *ConfigWrapper) SetCurrentShow(id int) {
	c.Config.CurrentShow = id
}

func (c *ConfigWrapper) SetCurrentScene(id int) {
	show, ok := c.GetCurrentShow()
	if !ok {
		return
	}
	show.CurrentScene = id
	log.Println(id)
}

func (c *ConfigWrapper) CreateShow(name string) *Show {
	id := c.Config.NextShowId
	c.Config.NextShowId++

	show := &Show{
		Id:           id,
		Name:         name,
		Scenes:       make([]*ShowScene, 0),
		CurrentScene: -1,
	}

	c.Config.Shows[id] = show

	return show
}

func (c *ConfigWrapper) RemoveShow(id int) {
	delete(c.Config.Shows, id)
}

func (c *ConfigWrapper) CreateScene(name string, movement int, measure int, sceneId int) *ShowScene {
	id := c.Config.NextSceneId
	c.Config.NextSceneId++

	scene := &ShowScene{
		Id:       id,
		Name:     name,
		Movement: movement,
		Measure:  measure,
		X32Scene: sceneId,
	}

	show, ok := c.GetCurrentShow()
	if !ok {
		return nil
	}
	show.Scenes = append(show.Scenes, scene)

	return scene
}

func (c *ConfigWrapper) GetSceneById(id int) (*ShowScene, int) {
	show, ok := c.GetCurrentShow()
	if !ok {
		return nil, 0
	}

	scenes := show.Scenes
	var (
		scene *ShowScene
		index int
	)
	for i, s := range scenes {
		if s.Id == id {
			scene = s
			index = i
			break
		}
	}
	return scene, index
}

func (c *ConfigWrapper) MoveScene(id int, newIndex int) {
	show, ok := c.GetCurrentShow()
	if !ok {
		return
	}
	scenes := show.Scenes
	scene, index := c.GetSceneById(id)
	if index == newIndex {
		return
	} else if index < newIndex {
		copy(scenes[index:newIndex], scenes[index+1:newIndex+1])
	} else {
		copy(scenes[newIndex+1:index+1], scenes[newIndex:index])
	}
	scenes[newIndex] = scene
}

func (c *ConfigWrapper) RemoveScene(id int) {
	show, ok := c.GetCurrentShow()
	if !ok {
		return
	}
	show.Scenes = slices.DeleteFunc(show.Scenes, func(scene *ShowScene) bool { return scene.Id == id })
}
