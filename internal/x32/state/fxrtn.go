package state

type X32StateFxrtnElement struct {
	Fconfig X32StateUniConfig
	Feq     X32StateUniEq4
	Fmix    X32StateUniMix16
	Fgrp    X32StateUniGrp
}

type X32StateFxrtn struct {
	Index [8]X32StateFxrtnElement `start:"1"`
}
