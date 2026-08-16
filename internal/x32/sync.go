package x32

import (
	"aumix/internal/osc"
	"time"
)

var NODES = []string{
	"config/chlink",
	"config/auxlink",
	"config/fxlink",
	"config/buslink",
	"config/mtxlink",
	"config/mute",
	"config/linkcfg",
	"config/mono",
	"config/solo",
	"config/talk",
	"config/talk/A",
	"config/talk/B",
	"config/osc",
	"config/routing/IN",
	"config/routing/AES50A",
	"config/routing/AES50B",
	"config/routing/CARD",
	"config/routing/OUT",
	"config/userctl/A",
	"config/userctl/B",
	"config/userctl/C",
	"config/userctl/A/enc",
	"config/userctl/B/enc",
	"config/userctl/C/enc",
	"config/userctl/A/btn",
	"config/userctl/B/btn",
	"config/userctl/C/btn",
}

func syncX32(send func(osc.Message) error) {
	for index, node := range NODES {
		if index%10 == 0 {
			time.Sleep(10 * time.Millisecond)
		}
		send(getNodeMessage(node))
	}
}

func getNodeMessage(node string) osc.Message {
	return osc.Message{
		Address: "/node",
		Parameters: []osc.Parameter{
			osc.StringParam(node),
		},
	}
}
