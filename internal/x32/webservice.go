package x32

import (
	"aumix/internal/osc"
	"aumix/internal/webserver"
	"fmt"
)

func getStatusMessage(manager *Manager) webserver.Message {
	return webserver.Message{
		Op:    webserver.MessageOpSET,
		Key:   "status",
		Value: manager.IsOSCConnected(),
	}
}

func getMixerAddressMessage(manager *Manager) webserver.Message {
	return webserver.Message{
		Op:  webserver.MessageOpSET,
		Key: "mixer-address",
		Value: mixerAddressMessage{
			Ip:   manager.OscClient.Ip,
			Port: manager.OscClient.Port,
		},
	}
}

func requestFaderValue(manager *Manager, faderId int, sender *webserver.Socket) {
	address := fmt.Sprintf("/ch/%02d/mix/fader", faderId)
	manager.ReceiveQueue.Enqueue(queueElement{
		sender: sender,
		key:    address,
	})
	manager.OscClient.Send(osc.Message{
		Address: address,
	})
}

// func getFaderMessage(manager *Manager, faderId int, value float32) webserver.Message {
// 	return webserver.Message{
// 		Op:    webserver.MessageOpSET_OSC,
// 		Key:   fmt.Sprintf("/ch/%02d/mix/fader", faderId),
// 		Value: value,
// 	}
// }
