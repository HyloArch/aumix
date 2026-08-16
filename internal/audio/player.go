package audio

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
)

func GetSamples() ([]string, error) {
	dir, err := os.ReadDir("data/samples")
	if err != nil {
		return nil, err
	}
	sampleFiles := make([]string, 0)
	for _, file := range dir {
		if !file.IsDir() {
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

	err = os.MkdirAll("data/samples", 0755)
	if err != nil {
		return err
	}

	return nil
}

func PlaySample(file string) error {
	Stop()

	f, err := os.Open(filepath.Join("data/samples", file))
	if err != nil {
		return err
	}

	ty := filepath.Ext(file)

	var (
		streamer beep.StreamSeekCloser
		format   beep.Format
	)

	switch ty {
	case ".mp3":
		streamer, format, err = mp3.Decode(f)
	case ".wav":
		streamer, format, err = wav.Decode(f)
	default:
		return fmt.Errorf("File format not supported.")
	}
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
