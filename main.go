package main

import (
	"aumix/internal/audio"
	"aumix/internal/osc"
	"aumix/internal/webserver"
	"aumix/internal/x32"
	"aumix/internal/x32/state"
	"embed"
	"io/fs"
	"log"
	"os"
	"os/signal"
	"syscall"
)

//go:embed frontend/dist/*
var reactAssets embed.FS

var xremote = osc.Message{Address: "/xremote"}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	configWrapper, err := state.LoadX32Config("data/x32state.gob")
	if err != nil {
		log.Printf("Error reading config from file: %v", err)
		configWrapper = state.NewX32Config()
	}

	// err = configWrapper.SetByPath("config/auxlink", []any{"OFF", "OFF", "ON", "OFF"})
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// value, err := configWrapper.GetByPath("config/auxlink")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(value)

	// return

	defer configWrapper.Save("data/x32state.gob")

	client := &osc.Client{}
	func() {
		config := configWrapper.Lock()
		defer configWrapper.Unlock()
		err = client.ConnectTo(config.MixerIp, config.MixerPort)
		if err != nil {
			log.Fatalf("Failed to create OSC client: %v", err)
		}
	}()
	defer client.Close()

	strippedFS, err := fs.Sub(reactAssets, "frontend/dist")
	if err != nil {
		log.Fatalf("Failed to get web directory: %v", err)
	}
	webserver.SetFS(strippedFS)
	server := webserver.NewServer(8080)
	defer server.Shutdown()
	server.InitRoutes()

	err = audio.InitPlayer(48000)
	if err != nil {
		log.Fatalf("Failed to initialize audio player: %v", err)
	}
	defer audio.Close()

	manager := x32.NewManager(1, configWrapper, client, server)
	defer manager.Shutdown()
	x32.RegisterOscHandlers(manager)
	x32.RegisterWebHandlers(manager)
	manager.StartServices()
	manager.Run()

	<-sigChan
}
