package state

type X32StateConfigChlink struct {
	Index [16]X32EnumOnType `pairs:"2"`
}

type X32StateConfigAuxlink struct {
	Index [4]X32EnumOnType `pairs:"2"`
}

type X32StateConfigFxlink struct {
	Index [4]X32EnumOnType `pairs:"2"`
}

type X32StateConfigBuslink struct {
	Index [8]X32EnumOnType `pairs:"2"`
}

type X32StateConfigMtxlink struct {
	Index [3]X32EnumOnType `pairs:"2"`
}

type X32StateConfigMute struct {
	Index [6]X32EnumOnType `start:"1"`
}

type X32StateConfigLinkcfg struct {
	Fhadly   X32EnumOnType
	Feq      X32EnumOnType
	Fdyn     X32EnumOnType
	Ffdrmute X32EnumOnType
}

type X32StateConfigMono struct {
	Fmode X32EnumMonoModeType
	Flink X32EnumOnType
}

type X32StateConfigSolo struct {
	Flevel      X32Level
	Fsource     X32EnumSoloSourceType
	Fsourcetrim X32Float
	Fchmode     X32EnumSoloModeType
	Fbusmode    X32EnumSoloModeType
	Fdcamode    X32EnumSoloModeType
	Fexclusive  X32EnumOnType
	Ffollowsel  X32EnumOnType
	Ffollowsolo X32EnumOnType
	Fdimatt     X32Float
	Fdim        X32EnumOnType
	Fmono       X32EnumOnType
	Fdelay      X32EnumOnType
	Fdelaytime  X32Float
	Fmasterctrl X32EnumOnType
	Fmute       X32EnumOnType
	Fdimpfl     X32EnumOnType
}

type X32StateConfigTalkElement struct {
	Flevel   X32Level
	Fdim     X32EnumOnType
	Flatch   X32EnumOnType
	Fdestmap X32Int
}

type X32StateConfigTalk struct {
	Fenable X32EnumOnType
	Fsource X32EnumTalkSourceType
	FA      X32StateConfigTalkElement
	FB      X32StateConfigTalkElement
}

func (s *X32StateConfigTalk) Get() []any {
	return []any{s.Fenable, s.Fsource}
}

func (s *X32StateConfigTalk) Set(values ...any) (int, error) {
	return 2, setAll([]X32StateValue{&s.Fenable, &s.Fsource}, values)
}

type X32StateConfigOsc struct {
	Flevel X32Level
	Ff1    X32Float
	Ff2    X32Float
	Ffsel  X32EnumFselType
	Ftype  X32EnumOscTypeType
	Fdest  X32EnumDestType
}

type X32StateConfigRoutingIN struct {
	Index [4]X32EnumInRoutingType `pairs:"8"`
	FAUX  X32EnumInAuxType
}

type X32StateConfigRoutingAES50A struct {
	Index [6]X32EnumRoutingType `pairs:"8"`
}

type X32StateConfigRoutingAES50B struct {
	Index [6]X32EnumRoutingType `pairs:"8"`
}

type X32StateConfigRoutingCARD struct {
	Index [4]X32EnumRoutingType `pairs:"8"`
}

type X32StateConfigRoutingOUT struct {
	F1_4   X32EnumOutRoutingLowType
	F9_12  X32EnumOutRoutingLowType
	F5_8   X32EnumOutRoutingHighType
	F13_16 X32EnumOutRoutingHighType
}

type X32StateConfigRouting struct {
	FIN     X32StateConfigRoutingIN
	FAES50A X32StateConfigRoutingAES50A
	FAES50B X32StateConfigRoutingAES50B
	FCARD   X32StateConfigRoutingCARD
	FOUT    X32StateConfigRoutingOUT
}

type X32StateConfigUserctrlEnc struct {
	Index [4]X32String `start:"1"`
}

type X32StateConfigUserctrlBtn struct {
	Index [8]X32String `start:"5"`
}

type X32StateConfigUserctrlElement struct {
	Fcolor X32EnumColorType
	Fenc   X32StateConfigUserctrlEnc
	Fbtn   X32StateConfigUserctrlBtn
}

type X32StateConfigUserctrl struct {
	FA X32StateConfigUserctrlElement
	FB X32StateConfigUserctrlElement
	FC X32StateConfigUserctrlElement
}

type X32StateConfigTape struct {
	FgainL    X32Float
	FgainR    X32Float
	Fautoplay X32EnumOnType
}

type X32StateConfig struct {
	Fchlink   X32StateConfigChlink
	Fauxlink  X32StateConfigAuxlink
	Ffxlink   X32StateConfigFxlink
	Fbuslink  X32StateConfigBuslink
	Fmtxlink  X32StateConfigMtxlink
	Fmute     X32StateConfigMute
	Flinkcfg  X32StateConfigLinkcfg
	Fmono     X32StateConfigMono
	Fsolo     X32StateConfigSolo
	Ftalk     X32StateConfigTalk
	Fosc      X32StateConfigOsc
	Frouting  X32StateConfigRouting
	FuserCtrl X32StateConfigUserctrl
	Ftape     X32StateConfigTape
}
