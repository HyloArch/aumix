package state

type X32StateChPreamp struct {
	Ftrim    X32Float
	Finvert  X32EnumOnType
	Fhpon    X32EnumOnType
	Fhpslope X32EnumHpSlopeType
	Fhpf     X32Float
}

type X32StateChGate struct {
	Fon      X32EnumOnType
	Fmode    X32EnumGateModeType
	Fthr     X32Float
	Frange   X32Float
	Fattack  X32Float
	Fhold    X32Float
	Frelease X32Float
	Fkeysrc  X32EnumSourceType
	Ffilter  X32StateUniEffectFilter
}

func (s *X32StateChGate) Get() []any {
	return []any{s.Fon, s.Fmode, s.Fthr, s.Frange, s.Fattack, s.Fhold, s.Frelease}
}

func (s *X32StateChGate) Set(values ...any) (int, error) {
	return 7, setAll([]X32StateValue{&s.Fon, &s.Fmode, &s.Fthr, &s.Frange, &s.Fattack, &s.Fhold, &s.Frelease}, values)
}

type X32StateChElement struct {
	Fconfig X32StateUniConfigSource
	Fdelay  X32StateUniDelay
	Fpreamp X32StateChPreamp
	Fgate   X32StateChGate
	Fdyn    X32StateUniDynKeysource
	Finsert X32StateUniInsert
	Feq     X32StateUniEq4
	Fmix    X32StateUniMix16
	Fgrp    X32StateUniGrp
}

type X32StateCh struct {
	Index [32]X32StateChElement `start:"1"`
}
