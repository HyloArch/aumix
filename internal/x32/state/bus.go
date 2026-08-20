package state

type X32StateBusElement struct {
	Fconfig X32StateUniConfig
	Fdyn    X32StateUniDynKeysource
	Finsert X32StateUniInsert
	Feq     X32StateUniEq6
	Fmix    X32StateUniMix6
	Fgrp    X32StateUniGrp
}

type X32StateBus struct {
	Index [16]X32StateBusElement `start:"1"`
}
