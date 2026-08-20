package state

// universal

type X32StateUniConfig struct {
	Fname  X32String
	Ficon  X32Int
	Fcolor X32EnumColorType
}

type X32StateUniConfigSource struct {
	Fname   X32String
	Ficon   X32Int
	Fcolor  X32EnumColorType
	Fsource X32EnumSourceType
}

type X32StateUniEq4 struct {
	Fon   X32EnumOnType
	Index [4]X32StateUniEqElement `start:"1"`
}

func (s *X32StateUniEq4) Get() []any {
	return []any{s.Fon}
}

func (s *X32StateUniEq4) Set(values ...any) (int, error) {
	return 1, setAll([]X32StateValue{&s.Fon}, values)
}

type X32StateUniEq6 struct {
	Fon   X32EnumOnType
	Index [6]X32StateUniEqElement `start:"1"`
}

func (s *X32StateUniEq6) Get() []any {
	return []any{s.Fon}
}

func (s *X32StateUniEq6) Set(values ...any) (int, error) {
	return 1, setAll([]X32StateValue{&s.Fon}, values)
}

type X32StateUniEqElement struct {
	Ftype X32EnumEqTypeType
	Ff    X32Float
	Fg    X32Float
	Fq    X32Float
}

type X32StateUniMix16 struct {
	Fon     X32EnumOnType
	Ffader  X32Level
	Fst     X32EnumOnType
	Fpan    X32Float
	Fmono   X32EnumOnType
	Fmlevel X32Level
	F01     X32StateUniMixPanElement
	F02     X32StateUniMixElement
	F03     X32StateUniMixPanElement
	F04     X32StateUniMixElement
	F05     X32StateUniMixPanElement
	F06     X32StateUniMixElement
	F07     X32StateUniMixPanElement
	F08     X32StateUniMixElement
	F09     X32StateUniMixPanElement
	F10     X32StateUniMixElement
	F11     X32StateUniMixPanElement
	F12     X32StateUniMixElement
	F13     X32StateUniMixPanElement
	F14     X32StateUniMixElement
	F15     X32StateUniMixPanElement
	F16     X32StateUniMixElement
}

func (s *X32StateUniMix16) Get() []any {
	return []any{s.Fon, s.Ffader, s.Fst, s.Fpan, s.Fmono, s.Fmlevel}
}

func (s *X32StateUniMix16) Set(values ...any) (int, error) {
	return 6, setAll([]X32StateValue{&s.Fon, &s.Ffader, &s.Fst, &s.Fpan, &s.Fmono, &s.Fmlevel}, values)
}

type X32StateUniMix6 struct {
	Fon     X32EnumOnType
	Ffader  X32Level
	Fst     X32EnumOnType
	Fpan    X32Float
	Fmono   X32EnumOnType
	Fmlevel X32Level
	F01     X32StateUniMixPanElement
	F02     X32StateUniMixElement
	F03     X32StateUniMixPanElement
	F04     X32StateUniMixElement
	F05     X32StateUniMixPanElement
	F06     X32StateUniMixElement
}

func (s *X32StateUniMix6) Get() []any {
	return []any{s.Fon, s.Ffader, s.Fst, s.Fpan, s.Fmono, s.Fmlevel}
}

func (s *X32StateUniMix6) Set(values ...any) (int, error) {
	return 6, setAll([]X32StateValue{&s.Fon, &s.Ffader, &s.Fst, &s.Fpan, &s.Fmono, &s.Fmlevel}, values)
}

type X32StateUniMixElement struct {
	Fon    X32EnumOnType
	Flevel X32Level
}

type X32StateUniMixPanElement struct {
	Fon    X32EnumOnType
	Flevel X32Level
	Fpan   X32Float
	Ftype  X32EnumMixTypeType
}

type X32StateUniGrp struct {
	Fdca  X32Int
	Fmute X32Int
}

type X32StateUniDynKeysource struct {
	Fon      X32EnumOnType
	Fmode    X32EnumDynModeType
	Fdet     X32EnumDetType
	Fenv     X32EnumEnvType
	Fthr     X32Float
	Fratio   X32EnumRatioType
	Fknee    X32Float
	Fmgain   X32Float
	Fattack  X32Float
	Fhold    X32Float
	Frelease X32Float
	Fpos     X32EnumPosType
	Fkeysrc  X32EnumSourceType
	Fmix     X32Float
	Fauto    X32EnumOnType
	Ffilter  X32StateUniEffectFilter
}

func (s *X32StateUniDynKeysource) Get() []any {
	return []any{s.Fon, s.Fmode, s.Fdet, s.Fenv, s.Fthr, s.Fratio, s.Fknee, s.Fmgain, s.Fattack, s.Fhold, s.Frelease, s.Fpos, s.Fkeysrc, s.Fmix, s.Fauto}
}

func (s *X32StateUniDynKeysource) Set(values ...any) (int, error) {
	return 15, setAll([]X32StateValue{&s.Fon, &s.Fmode, &s.Fdet, &s.Fenv, &s.Fthr, &s.Fratio, &s.Fknee, &s.Fmgain, &s.Fattack, &s.Fhold, &s.Frelease, &s.Fpos, &s.Fkeysrc, &s.Fmix, &s.Fauto}, values)
}

type X32StateUniDyn struct {
	Fon      X32EnumOnType
	Fmode    X32EnumDynModeType
	Fdet     X32EnumDetType
	Fenv     X32EnumEnvType
	Fthr     X32Float
	Fratio   X32EnumRatioType
	Fknee    X32Float
	Fmgain   X32Float
	Fattack  X32Float
	Fhold    X32Float
	Frelease X32Float
	Fpos     X32EnumPosType
	Fmix     X32Float
	Fauto    X32EnumOnType
	Ffilter  X32StateUniEffectFilter
}

func (s *X32StateUniDyn) Get() []any {
	return []any{s.Fon, s.Fmode, s.Fdet, s.Fenv, s.Fthr, s.Fratio, s.Fknee, s.Fmgain, s.Fattack, s.Fhold, s.Frelease, s.Fpos, s.Fmix, s.Fauto}
}

func (s *X32StateUniDyn) Set(values ...any) (int, error) {
	return 14, setAll([]X32StateValue{&s.Fon, &s.Fmode, &s.Fdet, &s.Fenv, &s.Fthr, &s.Fratio, &s.Fknee, &s.Fmgain, &s.Fattack, &s.Fhold, &s.Frelease, &s.Fpos, &s.Fmix, &s.Fauto}, values)
}

type X32StateUniEffectFilter struct {
	Fon   X32EnumOnType
	Ftype X32EnumFilterTypeType
	Ff    X32Float
}

type X32StateUniInsert struct {
	Fon  X32EnumOnType
	Fpos X32EnumPosType
	Fsel X32EnumSelType
}

type X32StateUniDelay struct {
	Fon   X32EnumOnType
	Ftime X32Float
}

// state

type X32State struct {
	Fconfig  X32StateConfig
	Fch      X32StateCh
	Fauxin   X32StateAuxin
	Ffxrtn   X32StateFxrtn
	Fbus     X32StateBus
	Fmtx     X32StateMtx
	Fmain    X32StateMain
	Fdca     X32StateDca
	Ffx      X32StateFx
	Foutputs X32StateOutputs
	Fheadamp X32StateHeadamp
	F_insert X32StateInsert
	F_show   X32StateShow
	F_libs   X32StateLibs
	F_prefs  X32StatePrefs
}

type Config struct {
	MixerIp   string
	MixerPort int
	State     X32State
}
