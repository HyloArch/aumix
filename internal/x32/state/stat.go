package state

type X32StateStatScreenPage struct {
	Fpage X32Int
}

type X32StateStatScreen struct {
	Fscreen  X32EnumScreenType
	Fmutegrp X32EnumOnType
	Futils   X32EnumOnType
	FCHAN    X32StateStatScreenPage
	FMETER   X32StateStatScreenPage
	FROUTE   X32StateStatScreenPage
	FSETUP   X32StateStatScreenPage
	FLIB     X32StateStatScreenPage
	FFX      X32StateStatScreenPage
	FMON     X32StateStatScreenPage
	FUSB     X32StateStatScreenPage
	FSCENE   X32StateStatScreenPage
	FASSIGN  X32StateStatScreenPage
}

func (s *X32StateStatScreen) Get() []any {
	return []any{s.Fscreen, s.Fmutegrp, s.Futils}
}

func (s *X32StateStatScreen) Set(values ...any) (int, error) {
	return 3, setAll([]X32StateValue{&s.Fscreen, &s.Fmutegrp, &s.Futils}, values)
}

type X32StateStatSolosw struct {
	Index [80]X32EnumOnType `start:"1"`
}

type X32StateStatAes50 struct {
	FA     X32String
	FB     X32String
	Fstate X32Int
}

type X32StateStatTape struct {
	Fstate X32EnumTapeStateType
	Ffile  X32String
	Fetime X32Int
	Frtime X32Int
}

type X32StateStat struct {
	Fselidx       X32EnumStatSelidxType
	Fchfaderbank  X32Int
	Fgrpfaderbank X32Int
	Fsendsonfader X32EnumOnType
	Fbussendbank  X32Int
	Feqband       X32Int
	Fsolo         X32EnumOnType
	Fkeysolo      X32EnumOnType
	Fuserbank     X32Int
	Fautosave     X32EnumOnType
	Flock         X32EnumOnType
	Fusbmounted   X32EnumOnType
	Fremote       X32EnumOnType
	Frtamodeeq    X32EnumRTAModeType
	Frtamodegeq   X32EnumRTAModeType
	Frtaeqpre     X32EnumOnType
	Frtaeqpost    X32EnumOnType
	Frtasource    X32Int
	Fxcardtype    X32Int
	Fxcardsync    X32EnumOnType
	Fgeqonfdr     X32EnumOnType
	Fgeqpos       X32Int
	Fscreen       X32StateStatScreen
	Fsolosw       X32StateStatSolosw
	Faes50        X32StateStatAes50
	Ftape         X32StateStatTape
	Fosc          X32EnumOnType
}

func (s *X32StateStat) Get() []any {
	return []any{s.Fselidx, s.Fchfaderbank, s.Fgrpfaderbank, s.Fsendsonfader, s.Fbussendbank, s.Feqband, s.Fsolo, s.Fkeysolo, s.Fuserbank, s.Fautosave, s.Flock, s.Fusbmounted, s.Fremote, s.Frtamodeeq, s.Frtamodegeq, s.Frtaeqpre, s.Frtaeqpost, s.Frtasource, s.Fxcardtype, s.Fxcardsync, s.Fgeqonfdr, s.Fgeqpos}
}

func (s *X32StateStat) Set(values ...any) (int, error) {
	return 22, setAll([]X32StateValue{&s.Fselidx, &s.Fchfaderbank, &s.Fgrpfaderbank, &s.Fsendsonfader, &s.Fbussendbank, &s.Feqband, &s.Fsolo, &s.Fkeysolo, &s.Fuserbank, &s.Fautosave, &s.Flock, &s.Fusbmounted, &s.Fremote, &s.Frtamodeeq, &s.Frtamodegeq, &s.Frtaeqpre, &s.Frtaeqpost, &s.Frtasource, &s.Fxcardtype, &s.Fxcardsync, &s.Fgeqonfdr, &s.Fgeqpos}, values)
}
