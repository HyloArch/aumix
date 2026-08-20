package state

type X32StateAuxinPreamp struct {
	Ftrim   X32Float
	Finvert X32EnumOnType
}
type X32StateAuxinElement struct {
	Fconfig X32StateUniConfigSource
	Fpreamp X32StateAuxinPreamp
	Feq     X32StateUniEq4
	Fmix    X32StateUniMix16
	Fgrp    X32StateUniGrp
}

type X32StateAuxin struct {
	Index [8]X32StateAuxinElement `start:"1"`
}
