package state

// universal
type X32StateUniConfig struct {
	Name  string
	Icon  int32
	Color int32 `enum:"Color"`
}
type X32StateUniConfigSource struct {
	Name   string
	Icon   int32
	Color  int32 `enum:"Color"`
	Source int32 `enum:"Source"`
}
type X32StateUniEq4 struct {
	On    int32                   `enum:"On"`
	Index [4]X32StateUniEqElement `start:"1"`
}
type X32StateUniEq6 struct {
	On    int32                   `enum:"On"`
	Index [6]X32StateUniEqElement `start:"1"`
}
type X32StateUniEqElement struct {
	Type int32 `enum:"EqType"`
	F    float32
	G    float32
	Q    float32
}
type X32StateUniMix16 struct {
	On     int32 `enum:"On"`
	Fader  float32
	St     int32 `enum:"On"`
	Pan    float32
	Mono   int32 `enum:"On"`
	Mlevel float32
	Index  [16]X32StateUniMixElement `start:"1"`
}
type X32StateUniMix6 struct {
	On     int32 `enum:"On"`
	Fader  float32
	St     int32 `enum:"On"`
	Pan    float32
	Mono   int32 `enum:"On"`
	Mlevel float32
	Index  [6]X32StateUniMixElement `start:"1"`
}
type X32StateUniMixElement struct {
	On    int32 `enum:"On"`
	Level float32
	Pan   float32
	Type  int32 `enum:"MixType"`
}
type X32StateUniGrp struct {
	Dca  int32
	Mute int32
}
type X32StateUniDyn struct {
	On      int32 `enum:"On"`
	Mode    int32 `enum:"DynMode"`
	Det     int32 `enum:"Det"`
	Env     int32 `enum:"Env"`
	Thr     float32
	Ratio   int32 `enum:"Ratio"`
	Knee    float32
	Mgain   float32
	Attack  float32
	Hold    float32
	Release float32
	Pos     int32 `enum:"Pos"`
	Keysrc  int32 `enum:"Source"`
	Mix     float32
	Auto    int32 `enum:"On"`
	Filter  X32StateUniEffectFilter
}
type X32StateUniEffectFilter struct {
	On   int32 `enum:"On"`
	Type int32 `enum:"FilterType"`
	F    float32
}
type X32StateUniInsert struct {
	On  int32 `enum:"On"`
	Pos int32 `enum:"Pos"`
	Sel int32 `enum:"Sel"`
}

// config
type X32StateConfigChlink struct {
	Index [16]int32 `enum:"On" pairs:"2"`
}
type X32StateConfigAuxlink struct {
	Index [4]int32 `enum:"On" pairs:"2"`
}
type X32StateConfigFxlink struct {
	Index [4]int32 `enum:"On" pairs:"2"`
}
type X32StateConfigBuslink struct {
	Index [8]int32 `enum:"On" pairs:"2"`
}
type X32StateConfigMtxlink struct {
	Index [3]int32 `enum:"On" pairs:"2"`
}
type X32StateConfigMute struct {
	Index [6]int32 `enum:"On" start:"1"`
}
type X32StateConfigLinkcfg struct {
	Hadly   int32 `enum:"On"`
	Eq      int32 `enum:"On"`
	Dyn     int32 `enum:"On"`
	Fdrmute int32 `enum:"On"`
}
type X32StateConfigMono struct {
	Mode int32 `enum:"MonoMode"`
	Link int32 `enum:"On"`
}
type X32StateConfigSolo struct {
	Level      float32
	Source     int32 `enum:"SoloSource"`
	Sourcetrim float32
	Chmode     int32 `enum:"SoloMode"`
	Busmode    int32 `enum:"SoloMode"`
	Dcamode    int32 `enum:"SoloMode"`
	Exclusive  int32 `enum:"On"`
	Followsel  int32 `enum:"On"`
	Followsolo int32 `enum:"On"`
	Dimatt     float32
	Dim        int32 `enum:"On"`
	Mono       int32 `enum:"On"`
	Delay      int32 `enum:"On"`
	Delaytime  float32
	Masterctrl int32 `enum:"On"`
	Mute       int32 `enum:"On"`
	Dimpfl     int32 `enum:"On"`
}
type X32StateConfigTalkElement struct {
	Level   float32
	Dim     int32 `enum:"On"`
	Latch   int32 `enum:"On"`
	Destmap int32
}
type X32StateConfigTalk struct {
	Enable int32 `enum:"On"`
	Source int32 `enum:"TalkSource"`
	A      X32StateConfigTalkElement
	B      X32StateConfigTalkElement
}
type X32StateConfigOsc struct {
	Level float32
	F1    float32
	F2    float32
	Fsel  int32 `enum:"Fsel"`
	Type  int32 `enum:"OscType"`
	Dest  int32 `enum:"Dest"`
}
type X32StateConfigRoutingIN struct {
	Index [4]int32 `enum:"InRouting" pairs:"8"`
	AUX   int32    `enum:"InAux"`
}
type X32StateConfigRoutingAES50A struct {
	Index [6]int32 `enum:"Routing" pairs:"8"`
}
type X32StateConfigRoutingAES50B struct {
	Index [6]int32 `enum:"Routing" pairs:"8"`
}
type X32StateConfigRoutingCARD struct {
	Index [4]int32 `enum:"Routing" pairs:"8"`
}
type X32StateConfigRoutingOUT struct {
	Index [4]int32 `enum:"OutRouting" pairs:"4"`
}
type X32StateConfigRouting struct {
	In     X32StateConfigRoutingIN
	Aes50a X32StateConfigRoutingAES50A
	Aes50b X32StateConfigRoutingAES50B
	Card   X32StateConfigRoutingCARD
	Out    X32StateConfigRoutingOUT
}
type X32StateConfigUserctrlEnc struct {
	Index [4]string `start:"1"`
}
type X32StateConfigUserctrlBtn struct {
	Index [8]string `start:"5"`
}
type X32StateConfigUserctrlElement struct {
	Color int32 `enum:"Color"`
	Enc   X32StateConfigUserctrlEnc
	Btn   X32StateConfigUserctrlBtn
}
type X32StateConfigUserctrl struct {
	A X32StateConfigUserctrlElement
	B X32StateConfigUserctrlElement
	C X32StateConfigUserctrlElement
}
type X32StateConfigTape struct {
	GainL    float32
	GainR    float32
	Autoplay int32 `enum:"On"`
}

type X32StateConfig struct {
	Chlink   X32StateConfigChlink
	Auxlink  X32StateConfigAuxlink
	Fxlink   X32StateConfigFxlink
	Buslink  X32StateConfigBuslink
	Mtxlink  X32StateConfigMtxlink
	Mute     X32StateConfigMute
	Linkcfg  X32StateConfigLinkcfg
	Mono     X32StateConfigMono
	Solo     X32StateConfigSolo
	Talk     X32StateConfigTalk
	Osc      X32StateConfigOsc
	Routing  X32StateConfigRouting
	UserCtrl X32StateConfigUserctrl
	Tape     X32StateConfigTape
}

// ch
type X32StateChDelay struct {
	On   int32 `enum:"On"`
	Time float32
}
type X32StateChPreamp struct {
	Trim    float32
	Invert  int32 `enum:"On"`
	Hpon    int32 `enum:"On"`
	Hpslope int32 `enum:"HpSlope"`
	Hpf     float32
}
type X32StateChGate struct {
	On      int32 `enum:"On"`
	Mode    int32 `enum:"GateMode"`
	Thr     float32
	Range   float32
	Attack  float32
	Hold    float32
	Release float32
	Keysrc  int32 `enum:"Source"`
	Filter  X32StateUniEffectFilter
}

type X32StateChElement struct {
	Config X32StateUniConfigSource
	Delay  X32StateChDelay
	Preamp X32StateChPreamp
	Gate   X32StateChGate
	Dyn    X32StateUniDyn
	Insert X32StateUniInsert
	Eq     X32StateUniEq4
	Mix    X32StateUniMix16
	Grp    X32StateUniGrp
}
type X32StateCh struct {
	Index [32]X32StateChElement `start:"1"`
}

// auxin
type X32StateAuxinPreamp struct {
	Trim   float32
	Invert int32 `enum:"On"`
}
type X32StateAuxinElement struct {
	Config X32StateUniConfigSource
	Preamp X32StateAuxinPreamp
	Eq     X32StateUniEq4
	Mix    X32StateUniMix16
	Grp    X32StateUniGrp
}

type X32StateAuxin struct {
	Index [8]X32StateAuxinElement `start:"1"`
}

// fxrtn
type X32StateFxrtnElement struct {
	Config X32StateUniConfig
	Eq     X32StateUniEq4
	Mix    X32StateUniMix16
	Grp    X32StateUniGrp
}

type X32StateFxrtn struct {
	Index [8]X32StateFxrtnElement `start:"1"`
}

// bus
type X32StateBusElement struct {
	Config X32StateUniConfig
	Dyn    X32StateUniDyn
	Insert X32StateUniInsert
	Eq     X32StateUniEq6
	Mix    X32StateUniMix6
	Grp    X32StateUniGrp
}

type X32StateBus struct {
	Index [16]X32StateBusElement `start:"1"`
}

// mtx
type X32StateMtxConfigPreamp struct {
	Invert int32 `enum:"On"`
}
type X32StateMtxConfig struct {
	Name   string
	Icon   int32
	Color  int32 `enum:"Color"`
	Preamp X32StateMtxConfigPreamp
}
type X32StateMtxMix struct {
	On    int32 `enum:"On"`
	Fader float32
}
type X32StateMtxElement struct {
	Config X32StateMtxConfig
	Dyn    X32StateUniDyn
	Insert X32StateUniInsert
	Eq     X32StateUniEq6
	Mix    X32StateMtxMix
}

type X32StateMtx struct {
	Index [6]X32StateMtxElement `start:"1"`
}

// main
type X32StateMainMix struct {
	On    int32 `enum:"On"`
	Fader float32
	Pan   float32
	Index [6]X32StateUniMixElement `start:"1"`
}
type X32StateMainElement struct {
	Config X32StateUniConfig
	Dyn    X32StateUniDyn
	Insert X32StateUniInsert
	Eq     X32StateUniEq6
	Mix    X32StateMainMix
}

type X32StateMain struct {
	St X32StateMainElement
	M  X32StateMainElement
}

// dca
type X32StateDcaElement struct {
	On     int32 `enum:"On"`
	Fader  float32
	Config X32StateUniConfig
}

type X32StateDca struct {
	Index [8]X32StateDcaElement
}

// fx
type X32StateFx struct {
}

// outputs
type X32StateOutputs struct {
}

// headamp
type X32StateHeadampElement struct {
	Gain    float32
	Phantom int32 `enum:"On"`
}

type X32StateHeadamp struct {
	Index [127]X32StateHeadampElement
}

// -insert
type X32StateInsert struct {
}

// -show
type X32StateShowPrepos struct {
	Current int32
}
type X32StateShowfileShow struct {
	Name     string
	Inputs   int32
	Mxsends  int32
	Mxbuses  int32
	Console  int32
	Chan16   int32
	Chan32   int32
	Return   int32
	Buses    int32
	Lrmtxdca int32
	Effects  int32
}
type X32StateShowfileCueElement struct {
	Numb      int32
	Name      string
	Skip      int32
	Scene     int32
	Bit       int32
	Miditype  int32
	Midichan  int32
	Midipara1 int32
	Midipara2 int32
}
type X32StateShowfileCue struct {
	Index [100]X32StateShowfileCueElement
}
type X32StateShowfileSceneElement struct {
	Name    string
	Notes   string
	Safes   int32
	Hasdata int32
}
type X32StateShowfileScene struct {
	Index [100]X32StateShowfileSceneElement
}
type X32StateShowfileSnippetElement struct {
	Name     string
	Eventtyp int32
	Channels int32
	Auxbuses int32
	Maingrps int32
	Hasdata  int32
}
type X32StateShowfileSnippet struct {
	Index [100]X32StateShowfileSnippetElement
}
type X32StateShowfile struct {
	Show  X32StateShowfileShow
	Cue   X32StateShowfileCue
	Scene X32StateShowfileScene
}

type X32StateShow struct {
	Prepos   X32StateShowPrepos
	Showfile X32StateShowfile
}

// -libs
type X32StateLibs struct {
}

// -prefs
type X32StatePrefsRemote struct {
	Enable   int32 `enum:"On"`
	Protocol int32 `enum:""`
	Port     int32
	Ioenable int32 `enum:""`
}
type X32StatePrefsiQ struct {
}

type X32StatePrefs struct {
	Style             string
	Bright            float32
	Lcdcont           float32
	Ledbright         float32
	Lamp              float32
	Lampon            int32 `enum:"On"`
	Clockrate         int32 `enum:"Clockrate"`
	Clocksource       int32 `enum:"Clocksource"`
	Confirm_general   int32 `enum:"On"`
	Comfirm_overwrite int32 `enum:"On"`
	Confirm_sceneload int32 `enum:"On"`
	Viewrtn           int32 `enum:"On"`
	Selfollowbank     int32 `enum:"On"`
	Sceneadvance      int32 `enum:"On"`
	Safe_masterlevels int32 `enum:"On"`
	Haflags           int32
	Autosel           int32 `enum:"On"`
	Show_control      int32 `enum:""`
	Clockmode         int32 `enum:""`
	Hardmute          int32 `enum:""`
	Dcamute           int32 `enum:""`
	Invertmutes       int32 `enum:""`
	Remote            X32StatePrefsRemote
	IQ                X32StatePrefsiQ
}

// state
type X32State struct {
	Config  X32StateConfig
	Ch      X32StateCh
	Auxin   X32StateAuxin
	Fxrtn   X32StateFxrtn
	Bus     X32StateBus
	Mtx     X32StateMtx
	Main    X32StateMain
	Dca     X32StateDca
	Fx      X32StateFx
	Outputs X32StateOutputs
	Headamp X32StateHeadamp
	Insert  X32StateInsert
	Show    X32StateShow
	Libs    X32StateLibs
	Prefs   X32StatePrefs
}

type Config struct {
	MixerIp   string
	MixerPort int
	State     X32State
}
