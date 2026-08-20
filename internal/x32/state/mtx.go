package state

type X32StateMtxConfigPreamp struct {
	Finvert X32EnumOnType
}

type X32StateMtxConfig struct {
	Fname  X32String
	Ficon  X32Int
	Fcolor X32EnumColorType
}

type X32StateMtxMix struct {
	Fon    X32EnumOnType
	Ffader X32Level
}

type X32StateMtxElement struct {
	Fconfig X32StateMtxConfig
	Fpreamp X32StateMtxConfigPreamp
	Fdyn    X32StateUniDyn
	Finsert X32StateUniInsert
	Feq     X32StateUniEq6
	Fmix    X32StateMtxMix
}

type X32StateMtx struct {
	Index [6]X32StateMtxElement `start:"1"`
}
