package x32

import (
	"aumix/internal/osc"
	"fmt"
	"strconv"
	"time"
)

var CONFIG_NODES = []string{
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
	// "config/userctl/A",
	// "config/userctl/B",
	// "config/userctl/C",
	// "config/userctl/A/enc",
	// "config/userctl/B/enc",
	// "config/userctl/C/enc",
	// "config/userctl/A/btn",
	// "config/userctl/B/btn",
	// "config/userctl/C/btn",
}

func syncX32Config(send func(string)) {
	for _, node := range CONFIG_NODES {
		send(node)
	}
}

var CH_NODES = []string{
	"config",
	"delay",
	"preamp",
	"gate",
	"gate/filter",
	"dyn",
	"dyn/filter",
	"insert",
	"eq",
	"mix",
	"grp",
}

func syncX32Ch(send func(string)) {
	for index := range 32 {
		prefix := fmt.Sprintf("ch/%02d/", index+1)
		for _, node := range CH_NODES {
			send(prefix + node)
		}
		for index := range 4 {
			send(fmt.Sprintf("%seq/%d", prefix, index+1))
		}
		for index := range 16 {
			send(fmt.Sprintf("%smix/%02d", prefix, index+1))
		}
	}
}

var AUXIN_NODES = []string{
	"config",
	"preamp",
	"eq",
	"mix",
	"grp",
}

func syncX32Auxin(send func(string)) {
	for index := range 8 {
		prefix := fmt.Sprintf("auxin/%02d/", index+1)
		for _, node := range AUXIN_NODES {
			send(prefix + node)
		}
		for index := range 4 {
			send(fmt.Sprintf("%seq/%d", prefix, index+1))
		}
		for index := range 16 {
			send(fmt.Sprintf("%smix/%02d", prefix, index+1))
		}
	}
}

var FXRTN_NODES = []string{
	"config",
	"eq",
	"mix",
	"grp",
}

func syncX32Fxrtn(send func(string)) {
	for index := range 8 {
		prefix := fmt.Sprintf("fxrtn/%02d/", index+1)
		for _, node := range FXRTN_NODES {
			send(prefix + node)
		}
		for index := range 4 {
			send(fmt.Sprintf("%seq/%d", prefix, index+1))
		}
		for index := range 16 {
			send(fmt.Sprintf("%smix/%02d", prefix, index+1))
		}
	}
}

var BUS_NODES = []string{
	"config",
	"dyn",
	"dyn/filter",
	"insert",
	"eq",
	"mix",
	"grp",
}

func syncX32Bus(send func(string)) {
	for index := range 16 {
		prefix := fmt.Sprintf("bus/%02d/", index+1)
		for _, node := range BUS_NODES {
			send(prefix + node)
		}
		for index := range 6 {
			send(fmt.Sprintf("%seq/%d", prefix, index+1))
		}
		for index := range 6 {
			send(fmt.Sprintf("%smix/%02d", prefix, index+1))
		}
	}
}

var MTX_NODES = []string{
	"config",
	"preamp",
	"dyn",
	"dyn/filter",
	"insert",
	"eq",
	"mix",
}

func syncX32Mtx(send func(string)) {
	for index := range 6 {
		prefix := fmt.Sprintf("mtx/%02d/", index+1)
		for _, node := range MTX_NODES {
			send(prefix + node)
		}
		for index := range 6 {
			send(fmt.Sprintf("%seq/%d", prefix, index+1))
		}
	}
}

var MAIN_NODES = []string{
	"config",
	"dyn",
	"dyn/filter",
	"insert",
	"eq",
	"mix",
}

func syncX32Main(send func(string)) {
	for _, s := range []string{"st", "m"} {
		prefix := fmt.Sprintf("main/%s/", s)
		for _, node := range MAIN_NODES {
			send(prefix + node)
		}
		for index := range 6 {
			send(fmt.Sprintf("%seq/%d", prefix, index+1))
		}
		for index := range 6 {
			send(fmt.Sprintf("%smix/%02d", prefix, index+1))
		}
	}
}

func syncX32Dca(send func(string)) {
	for index := range 8 {
		send("dca/" + strconv.Itoa(index+1))
		send("dca/" + strconv.Itoa(index+1) + "/config")
	}
}

func syncX32Fx(send func(string)) {
	for index := range 8 {
		send("fx/" + strconv.Itoa(index+1))
		send("fx/" + strconv.Itoa(index+1) + "/config")
		send("fx/" + strconv.Itoa(index+1) + "/par")
	}
}

func syncX32Outputs(send func(string)) {
	for index := range 16 {
		prefix := fmt.Sprintf("outputs/main/%02d", index+1)
		send(prefix)
		send(prefix + "/delay")
	}
	for index := range 6 {
		send(fmt.Sprintf("outputs/aux/%02d", index+1))
	}
	for index := range 16 {
		prefix := fmt.Sprintf("outputs/p16/%02d", index+1)
		send(prefix)
		send(prefix + "/iQ")
	}
	for index := range 2 {
		send(fmt.Sprintf("outputs/aes/%02d", index+1))
		send(fmt.Sprintf("outputs/rec/%02d", index+1))
	}
}

func syncX32Headamp(send func(string)) {
	for index := range 128 {
		send(fmt.Sprintf("headamp/%03d", index))
	}
}

func syncX32Insert(send func(string)) {
	send("-insert")
}

func syncX32(send func(osc.Message) error) {
	index := 0
	throttledSend := func(node string) {
		if index%10 == 0 {
			time.Sleep(10 * time.Millisecond)
		}
		send(getNodeMessage(node))
		index += 1
	}

	syncX32Config(throttledSend)
	syncX32Ch(throttledSend)
	syncX32Auxin(throttledSend)
	syncX32Fxrtn(throttledSend)
	syncX32Bus(throttledSend)
	syncX32Mtx(throttledSend)
	syncX32Main(throttledSend)
	syncX32Dca(throttledSend)
	syncX32Fx(throttledSend)
	syncX32Outputs(throttledSend)
	syncX32Headamp(throttledSend)
	syncX32Insert(throttledSend)
	throttledSend("-show/showfile/cue")
}

func getNodeMessage(node string) osc.Message {
	return osc.Message{
		Address: "/node",
		Parameters: []osc.Parameter{
			osc.StringParam(node),
		},
	}
}
