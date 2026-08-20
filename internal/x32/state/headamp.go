package state

type X32StateHeadampElement struct {
	Fgain    X32Float
	Fphantom X32EnumOnType
}

type X32StateHeadamp struct {
	Index [128]X32StateHeadampElement
}
