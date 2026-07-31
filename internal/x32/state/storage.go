package state

import (
	"aumix/internal/osc"
	"encoding/gob"
	"fmt"
	"iter"
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type ConfigWrapper struct {
	Config *Config
	lock   sync.Mutex
}

func (c *ConfigWrapper) Save(filePath string) error {
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

var capitalizer = cases.Title(language.Und)

func (c *ConfigWrapper) resolvePath(path string) (reflect.Value, reflect.StructField, error) {
	parts := strings.Split(path, "/")
	if parts[0] == "" {
		parts = parts[1:]
	}
	for i := range parts {
		parts[i] = capitalizer.String(strings.TrimPrefix(parts[i], "-"))
	}

	state := &c.Config.State
	var parentValue reflect.Value
	currentValue := reflect.ValueOf(state).Elem()
	var field reflect.StructField
	var fieldName string

	for _, part := range parts {
		parentValue = currentValue
		fieldName = part
		currentValue = currentValue.FieldByName(fieldName)
		if !currentValue.IsValid() {
			fieldName = "Index"
			currentValue = parentValue.FieldByName(fieldName)
			indexField, _ := parentValue.Type().FieldByName(fieldName)
			if !currentValue.IsValid() {
				return currentValue, field, fmt.Errorf("Field \"%s\" does not exist", part)
			}
			index, err := strconv.Atoi(part)
			if err != nil {
				pairString, ok := indexField.Tag.Lookup("pairs")
				if !ok {
					return currentValue, field, fmt.Errorf("Fragment \"%s\" is not an integer", part)
				}
				split := strings.Split(part, "-")
				index, err = strconv.Atoi(split[0])
				if err != nil {
					return currentValue, field, fmt.Errorf("Field \"%s\" does not exist", part)
				}
				pair, err := strconv.Atoi(pairString)
				if err != nil {
					return currentValue, field, fmt.Errorf("Error parsing field")
				}
				index /= pair
			}
			startString, ok := indexField.Tag.Lookup("start")
			if ok {
				start, err := strconv.Atoi(startString)
				if err != nil {
					return currentValue, field, fmt.Errorf("Error parsing field")
				}
				index -= start
			}
			if index >= currentValue.Len() {
				return currentValue, field, fmt.Errorf("%d is out of range", index)
			}
			currentValue = currentValue.Index(index)
		}
	}
	field, _ = parentValue.Type().FieldByName(fieldName)

	return currentValue, field, nil
}

func (c *ConfigWrapper) GetByPath(path string) (any, error) {
	c.Lock()
	defer c.Unlock()

	value, _, err := c.resolvePath(path)
	return value, err
}

func (c *ConfigWrapper) setValue(toSet reflect.Value, field reflect.StructField, next func() (any, bool)) error {
	if toSet.Kind() == reflect.Struct {
		for f, v := range toSet.Fields() {
			err := c.setValue(v, f, next)
			if err != nil {
				return err
			}
		}
		return nil
	}

	value, ok := next()
	if !ok {
		return fmt.Errorf("Not enough parameters provided")
	}

	enumType, enumOk := field.Tag.Lookup("enum")
	if enumOk {
		if stringValue, stringOk := value.(string); stringOk {
			value, ok = X32Enums[enumType][stringValue]
			if !ok {
				return fmt.Errorf("Provided string %s is not an element of enum %v", stringValue, enumType)
			}
		} else if intValue, intOk := value.(int); intOk {
			if intValue >= len(X32Enums[enumType]) {
				return fmt.Errorf("Int %v is out of enum range", intValue)
			}
		}
	}
	newValue := reflect.ValueOf(value)
	ok = newValue.CanConvert(toSet.Type())
	if !ok {
		return fmt.Errorf("Can't convert type %v to type %v", newValue.Type(), toSet.Type())
	}
	newValue = newValue.Convert(toSet.Type())
	toSet.Set(newValue)

	return nil
}

func (c *ConfigWrapper) SetByPath(path string, values []any) error {
	c.Lock()
	defer c.Unlock()

	currentValue, field, err := c.resolvePath(path)
	if err != nil {
		return err
	}

	next, stop := iter.Pull(slices.Values(values))
	defer stop()

	return c.setValue(currentValue, field, next)
}

func (c *ConfigWrapper) setParameterValue(toSet reflect.Value, field reflect.StructField, next func() (osc.Parameter, bool)) error {
	if toSet.Kind() == reflect.Struct {
		for f, v := range toSet.Fields() {
			err := c.setParameterValue(v, f, next)
			if err != nil {
				return err
			}
		}
		return nil
	}

	value, ok := next()
	if !ok {
		return fmt.Errorf("Not enough parameters provided")
	}

	enumType, enumOk := field.Tag.Lookup("enum")
	if enumOk {
		if stringValue, stringOk := value.(string); stringOk {
			value, ok = X32Enums[enumType][stringValue]
			if !ok {
				return fmt.Errorf("Provided string %s is not an element of enum %v", stringValue, enumType)
			}
		} else if intValue, intOk := value.(int); intOk {
			if intValue >= len(X32Enums[enumType]) {
				return fmt.Errorf("Int %v is out of enum range", intValue)
			}
		}
	}
	newValue := reflect.ValueOf(value)
	ok = newValue.CanConvert(toSet.Type())
	if !ok {
		return fmt.Errorf("Can't convert type %v to type %v", newValue.Type(), toSet.Type())
	}
	newValue = newValue.Convert(toSet.Type())
	toSet.Set(newValue)

	return nil
}

func (c *ConfigWrapper) SetParametersByPath(path string, values []osc.Parameter) error {
	c.Lock()
	defer c.Unlock()

	currentValue, field, err := c.resolvePath(path)
	if err != nil {
		return err
	}

	next, stop := iter.Pull(slices.Values(values))
	defer stop()

	return c.setParameterValue(currentValue, field, next)
}

func (c *ConfigWrapper) SetFader(index int, level float32) error {
	config := c.Lock()
	defer c.Unlock()

	channels := &config.State.Ch.Index
	if index >= len(channels) {
		return fmt.Errorf("Fader index %d is out of range %d", index, len(channels))
	}
	channels[index].Mix.Fader = level

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
		Config: &Config{},
	}
	return configWrapper
}
