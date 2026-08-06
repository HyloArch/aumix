package main

import (
	"aumix/internal/osc"
	"aumix/internal/webserver"
	"aumix/internal/x32"
	"aumix/internal/x32/state"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var xremote = osc.Message{Address: "/xremote"}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	configWrapper, err := state.LoadX32Config("x32state.gob")
	if err != nil {
		log.Printf("Error reading config from file: %v", err)
		configWrapper = state.NewX32Config()
	}
	defer configWrapper.Save("x32state.gob")

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

	server := webserver.NewServer(8080)
	defer server.Shutdown()
	server.InitRoutes()

	manager := x32.NewManager(1, configWrapper, client, server)
	defer manager.Shutdown()
	x32.RegisterOscHandlers(manager)
	x32.RegisterWebHandlers(manager)
	manager.StartServices()
	manager.Run()

	server.Broadcast(webserver.Message{
		Op:  webserver.MessageOpSET,
		Key: "meters",
		Value: map[string]any{
			"index":  0,
			"levels": []float32{0.5, 0, 0},
		},
	})

	<-sigChan
}
