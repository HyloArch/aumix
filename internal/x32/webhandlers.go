package x32

import (
	"aumix/internal/osc"
	"aumix/internal/webserver"
	"fmt"
)

func statusWeb(manager *Manager, message webserver.Message) {
	if message.Op == webserver.MessageOpGET {
		message.Sender.Send(getStatusMessage(manager))
	}
}

func mixerAddress(manager *Manager, message webserver.Message) {
	switch message.Op {
	case webserver.MessageOpGET:
		message.Sender.Send(getMixerAddressMessage(manager))
	case webserver.MessageOpSET:
		value := message.Value.(map[string]any)
		ip := value["ip"].(string)
		port := int(value["port"].(float64))
		if ip != manager.OscClient.Ip || port != manager.OscClient.Port {
			manager.OscClient.ConnectTo(ip, port)
			manager.oscConnectStatus = 0
			manager.Webserver.Broadcast(getStatusMessage(manager))

			func() {
				config := manager.ConfigWrapper.Lock()
				defer manager.ConfigWrapper.Unlock()
				config.MixerIp = ip
				config.MixerPort = port
			}()
		}
		manager.refreshX32Connection()
		manager.ReceiveQueue.Enqueue(queueElement{
			sender: message.Sender,
			key:    "/status",
		})
	}
}

func syncWeb(manager *Manager, message webserver.Message) {
	if message.Op == webserver.MessageOpGET {
		sendFunc := message.Sender.Send
		sendFunc(getStatusMessage(manager))
		sendFunc(getMixerAddressMessage(manager))
		for i := 1; i <= 32; i++ {
			requestFaderValue(manager, i, message.Sender)
		}
	}
}

func fader1Web(manager *Manager, message webserver.Message) {
	switch message.Op {
	case webserver.MessageOpGET_OSC:
		requestFaderValue(manager, 1, message.Sender)
	case webserver.MessageOpSET_OSC:
		value := message.Value.(float64)
		manager.OscClient.Send(osc.Message{
			Address: "/ch/01/mix/fader",
			Parameters: []osc.Parameter{
				osc.FloatParam(value),
			},
		})
		manager.Webserver.BroadcastExcept(message, message.Sender)
	}
}

func mixFader(manager *Manager, message webserver.Message) {
	switch message.Op {
	case webserver.MessageOpGET:

	case webserver.MessageOpSET:
		value := message.Value.(map[string]any)
		ty := value["type"]
		index := value["index"].(float64)
		level := value["level"].(float64)
		manager.OscClient.Send(osc.Message{
			Address: fmt.Sprintf("/%s/%02d/mix/fader", ty, int(index)),
			Parameters: []osc.Parameter{
				osc.FloatParam(level),
			},
		})
		manager.Webserver.BroadcastExcept(message, message.Sender)
	}
}

func RegisterWebHandlers(manager *Manager) {
	manager.RegisterWebHandler("status", statusWeb)
	manager.RegisterWebHandler("mixer-address", mixerAddress)
	manager.RegisterWebHandler("sync", syncWeb)
	manager.RegisterWebHandler("/ch/01/mix/fader", fader1Web)
	manager.RegisterWebHandler("mix-fader", mixFader)
}
