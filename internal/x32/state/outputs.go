package state

type X32StateOutputsMainElement struct {
	Fsrc    X32EnumOutputSourceType
	Fpos    X32EnumOutputPosType
	Finvert X32EnumOnType
	Fdelay  X32StateUniDelay
}

func (s *X32StateOutputsMainElement) Get() []any {
	return []any{s.Fsrc, s.Fpos, s.Finvert}
}

func (s *X32StateOutputsMainElement) Set(values ...any) (int, error) {
	return 3, setAll([]X32StateValue{&s.Fsrc, &s.Fpos, &s.Finvert}, values)
}

type X32StateOutputsMain struct {
	Index [16]X32StateOutputsMainElement `start:"1"`
}

type X32StateOutputsElement struct {
	Fsrc    X32EnumOutputSourceType
	Fpos    X32EnumOutputPosType
	Finvert X32EnumOnType
}

type X32StateOutputsAux struct {
	Index [6]X32StateOutputsElement `start:"1"`
}

type X32StateOutputsP16IQ struct {
	Fgroup   X32EnumIQGroupType
	Fspeaker X32EnumIQSpeakerType
	Feq      X32EnumIQEqType
	Fmodel   X32Int
}

type X32StateOutputsP16Element struct {
	Fsrc    X32EnumOutputSourceType
	Fpos    X32EnumOutputPosType
	Finvert X32EnumOnType
	FiQ     X32StateOutputsP16IQ
}

func (s *X32StateOutputsP16Element) Get() []any {
	return []any{s.Fsrc, s.Fpos, s.Finvert}
}

func (s *X32StateOutputsP16Element) Set(values ...any) (int, error) {
	return 3, setAll([]X32StateValue{&s.Fsrc, &s.Fpos, &s.Finvert}, values)
}

type X32StateOutputsP16 struct {
	Index [16]X32StateOutputsP16Element `start:"1"`
}

type X32StateOutputsAes struct {
	F01 X32StateOutputsElement
	F02 X32StateOutputsElement
}

type X32StateOutputsRecElement struct {
	Fsrc X32EnumOutputSourceType
	Fpos X32EnumOutputPosType
}

type X32StateOutputsRec struct {
	F01 X32StateOutputsRecElement
	F02 X32StateOutputsRecElement
}

type X32StateOutputs struct {
	Fmain X32StateOutputsMain
	Faux  X32StateOutputsAux
	Fp16  X32StateOutputsP16
	Faes  X32StateOutputsAes
	Frec  X32StateOutputsRec
}
