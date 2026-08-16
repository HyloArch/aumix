package x32

import (
	"aumix/internal/osc"
	"aumix/internal/webserver"
	"encoding/binary"
	"fmt"
	"regexp"
	"strconv"
)

func statusOsc(manager *Manager, _ osc.Message, replyFunc func(webserver.Message) error) {
	response := webserver.Message{
		Op:    webserver.MessageOpSET,
		Key:   "status",
		Value: true,
	}
	if !manager.IsOSCConnected() {
		manager.Webserver.Broadcast(response)
	} else if replyFunc != nil {
		replyFunc(response)
	}
	if manager.oscConnectStatus < 1 {
		manager.syncToMixer()
	}
	manager.oscConnectStatus = 2
}

func stopped(manager *Manager, _ osc.Message, _ func(webserver.Message) error) {
	if manager.IsOSCConnected() {
		manager.Webserver.Broadcast(webserver.Message{
			Op:    webserver.MessageOpSET,
			Key:   "status",
			Value: false,
		})
	}
	manager.oscConnectStatus = 0
}

func defaultHandler(manager *Manager, message osc.Message, replyFunc func(webserver.Message) error) {
	fmt.Println("Default: ", message)

	err := manager.ConfigWrapper.SetParametersByPath(message.Address, message.Parameters)
	if err != nil {
		fmt.Printf("Error setting values for node: %v\n", err)
	}
}

func node(manager *Manager, message osc.Message, replyFunc func(webserver.Message) error) {
	parameterString, ok := message.Parameters[0].(osc.StringParam)
	if !ok {
		fmt.Println("Node parameter isn't a string")
		return
	}
	address, values := DecodeNode(string(parameterString))
	fmt.Println(values...)
	err := manager.ConfigWrapper.SetByPath(address, values)
	if err != nil {
		fmt.Printf("Error setting values from node: %v\n", err)
	}
}

func faderOsc(manager *Manager, message osc.Message, replyFunc func(webserver.Message) error) {
	index, err := strconv.Atoi(message.Address[4:6])
	if err != nil {
		fmt.Printf("Error get fader index: %v", err)
		return
	}
	value, ok := message.Parameters[0].(osc.FloatParam)
	if !ok {
		fmt.Println("Fader parameter is not a float")
		return
	}
	manager.ConfigWrapper.SetFader(index, float32(value))
	response := webserver.Message{
		Op:  webserver.MessageOpSET,
		Key: "mix-fader",
		Value: map[string]any{
			"type":  "ch",
			"index": index,
			"level": value,
		},
	}
	if replyFunc != nil {
		replyFunc(response)
	} else {
		manager.Webserver.Broadcast(response)
	}
}

func meters(manager *Manager, message osc.Message, replyFunc func(webserver.Message) error) {
	index, err := strconv.Atoi(message.Address[8:])
	if err != nil {
		fmt.Printf("Error get fader index: %v", err)
		return
	}
	blob, ok := message.Parameters[0].(osc.ByteBlobParam)
	if !ok {
		fmt.Println("First value of meters response is not a byte blob")
		return
	}
	length := binary.LittleEndian.Uint32(blob[:4])

	manager.Webserver.Broadcast(webserver.Message{
		Op:  webserver.MessageOpSET,
		Key: "meters",
		Value: map[string]any{
			"index":  index,
			"length": length,
			"levels": blob[4:],
		},
	})
}

func RegisterOscHandlers(manager *Manager) {
	manager.SetDefaultOscHandler(defaultHandler)
	manager.RegisterOscHandler(regexp.MustCompile(`/meters/\d+`), meters)
	manager.RegisterOscHandler(regexp.MustCompile(`node`), node)
	manager.RegisterOscHandler(regexp.MustCompile(`stopped`), stopped)
	manager.RegisterOscHandler(regexp.MustCompile(`/status`), statusOsc)
	manager.RegisterOscHandler(regexp.MustCompile(`/ch/\d+/mix/fader`), faderOsc)
}
