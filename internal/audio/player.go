package audio

import (
	"os"
	"path/filepath"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

func GetSamples() ([]string, error) {
	dir, err := os.ReadDir("data")
	if err != nil {
		return nil, err
	}
	sampleFiles := make([]string, 0)
	for _, file := range dir {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".mp3" {
			sampleFiles = append(sampleFiles, file.Name())
		}
	}
	return sampleFiles, nil
}

type Player struct {
	ctrl       *beep.Ctrl
	sample     string
	sampleRate beep.SampleRate
}

var player *Player

func InitPlayer(sampleRate int) error {
	targetSampleRate := beep.SampleRate(sampleRate)
	err := speaker.Init(targetSampleRate, targetSampleRate.N(time.Second/10))
	if err != nil {
		return err
	}
	player = &Player{
		sampleRate: targetSampleRate,
	}

	err = os.MkdirAll("data", 0755)
	if err != nil {
		return err
	}

	return nil
}

func PlaySample(file string) error {
	Stop()

	f, err := os.Open(filepath.Join("data", file))
	if err != nil {
		return err
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return err
	}

	resampledStreamer := beep.Resample(4, format.SampleRate, player.sampleRate, streamer)

	player.ctrl = &beep.Ctrl{
		Streamer: resampledStreamer,
		Paused:   false,
	}
	player.sample = file

	speaker.Play(player.ctrl)

	return nil
}

func Stop() {
	if player.ctrl == nil {
		return
	}
	speaker.Lock()
	defer speaker.Unlock()

	player.ctrl.Streamer = nil
}

func Pause() {
	if player.ctrl == nil {
		return
	}
	speaker.Lock()
	defer speaker.Unlock()

	player.ctrl.Paused = true
}

func Close() {
	speaker.Close()
}
