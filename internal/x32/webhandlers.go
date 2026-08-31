package x32

import (
	"aumix/internal/audio"
	"aumix/internal/osc"
	"aumix/internal/webserver"
	"aumix/internal/x32/state"
	"fmt"
	"log"
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

func show(manager *Manager, message webserver.Message) {
	switch message.Op {
	case webserver.MessageOpGET:
		manager.ConfigWrapper.Lock()
		defer manager.ConfigWrapper.Unlock()

		show, _ := manager.ConfigWrapper.GetCurrentShow()
		message.Sender.Send(webserver.Message{
			Op:    webserver.MessageOpSET,
			Key:   "show",
			Value: show,
		})

		if show.CurrentScene != -1 {
			message.Sender.Send(webserver.Message{
				Op:    webserver.MessageOpSET,
				Key:   "go-scene",
				Value: show.CurrentScene,
			})
		}
	case webserver.MessageOpSET:
		manager.ConfigWrapper.Lock()
		defer manager.ConfigWrapper.Unlock()

		value := message.Value.(map[string]any)
		idValue, ok := value["id"].(float64)
		if ok {
			id := int(idValue)
			show, ok := manager.ConfigWrapper.GetCurrentShow()
			if !ok || show.Id != id {
				manager.ConfigWrapper.SetCurrentShow(id)
				show, ok = manager.ConfigWrapper.GetCurrentShow()
				if !ok {
					return
				}
				show.CurrentScene = -1
			}

			remove, ok := value["remove"].(bool)
			if ok && remove {
				manager.ConfigWrapper.RemoveShow(id)
				manager.Webserver.Broadcast(webserver.Message{
					Op:    webserver.MessageOpSET,
					Key:   "show",
					Value: nil,
				})
				return
			}

			name, ok := value["name"].(string)
			if ok {
				show.Name = name
			}
			manager.Webserver.Broadcast(webserver.Message{
				Op:    webserver.MessageOpSET,
				Key:   "show",
				Value: show,
			})
		} else {
			name := value["name"].(string)
			show := manager.ConfigWrapper.CreateShow(name)
			manager.Webserver.Broadcast(webserver.Message{
				Op:    webserver.MessageOpSET,
				Key:   "show",
				Value: show,
			})
			manager.ConfigWrapper.SetCurrentShow(show.Id)
		}
	}
}

func scene(manager *Manager, message webserver.Message) {
	switch message.Op {
	case webserver.MessageOpSET:
		manager.ConfigWrapper.Lock()
		defer manager.ConfigWrapper.Unlock()

		value := message.Value.(map[string]any)
		idValue, ok := value["id"].(float64)
		if ok {
			id := int(idValue)

			remove, ok := value["remove"].(bool)
			if ok && remove {
				manager.ConfigWrapper.RemoveScene(id)
				break
			}

			newIndex, ok := value["newIndex"].(float64)
			if ok {
				manager.ConfigWrapper.MoveScene(id, int(newIndex))
				break
			}

			scene, _ := manager.ConfigWrapper.GetSceneById(id)

			name, ok := value["name"].(string)
			if ok {
				scene.Name = name
			}

			movement, ok := value["movement"].(float64)
			if ok {
				scene.Movement = int(movement)
			}

			measure, ok := value["measure"].(float64)
			if ok {
				scene.Measure = int(measure)
			}

			sceneId, ok := value["sceneId"].(float64)
			if ok {
				scene.X32Scene = int(sceneId)
			}

			samples, ok := value["samples"].([]any)
			if ok {
				scene.Samples = make([]state.Sample, len(samples))
				for index, sample := range samples {
					s := sample.(map[string]any)
					scene.Samples[index] = state.Sample{
						Name: s["name"].(string),
						File: s["file"].(string),
					}
				}
			}
		} else {
			name := value["name"].(string)
			movement := int(value["movement"].(float64))
			measure := int(value["measure"].(float64))
			sceneId := int(value["sceneId"].(float64))
			manager.ConfigWrapper.CreateScene(name, movement, measure, sceneId)
		}
	default:
		return
	}

	show, _ := manager.ConfigWrapper.GetCurrentShow()
	manager.Webserver.Broadcast(webserver.Message{
		Op:    webserver.MessageOpSET,
		Key:   "show",
		Value: show,
	})
}

func showList(manager *Manager, message webserver.Message) {
	if message.Op == webserver.MessageOpGET {
		shows := manager.ConfigWrapper.GetShowList()
		message.Sender.Send(webserver.Message{
			Op:    webserver.MessageOpSET,
			Key:   "show-list",
			Value: shows,
		})
	}
}

func goSceneWeb(manager *Manager, message webserver.Message) {
	if message.Op == webserver.MessageOpSET {
		manager.ConfigWrapper.Lock()
		defer manager.ConfigWrapper.Unlock()

		sceneId := int(message.Value.(float64))
		scene, _ := manager.ConfigWrapper.GetSceneById(sceneId)

		manager.ConfigWrapper.SetCurrentScene(sceneId)
		manager.Webserver.Broadcast(webserver.Message{
			Op:    webserver.MessageOpSET,
			Key:   "go-scene",
			Value: sceneId,
		})

		manager.OscClient.Send(osc.Message{
			Address:    "/-action/goscene",
			Parameters: []osc.Parameter{osc.IntParam(scene.X32Scene)},
		})
	}
}

func samples(manager *Manager, message webserver.Message) {
	switch message.Op {
	case webserver.MessageOpGET:
		sampleFiles, err := audio.GetSamples()
		if err != nil {
			log.Printf("Error getting samples: %v\n", err)
			return
		}
		message.Sender.Send(webserver.Message{
			Op:    webserver.MessageOpSET,
			Key:   "samples",
			Value: sampleFiles,
		})
	case webserver.MessageOpSET:
		sample, ok := message.Value.(string)
		if ok {
			log.Println(audio.PlaySample(sample))
		} else {
			audio.Pause()
		}
	}
}

func RegisterWebHandlers(manager *Manager) {
	manager.RegisterWebHandler("status", statusWeb)
	manager.RegisterWebHandler("mixer-address", mixerAddress)
	manager.RegisterWebHandler("sync", syncWeb)
	manager.RegisterWebHandler("/ch/01/mix/fader", fader1Web)
	manager.RegisterWebHandler("mix-fader", mixFader)
	manager.RegisterWebHandler("show", show)
	manager.RegisterWebHandler("scene", scene)
	manager.RegisterWebHandler("show-list", showList)
	manager.RegisterWebHandler("go-scene", goSceneWeb)
	manager.RegisterWebHandler("samples", samples)
}
