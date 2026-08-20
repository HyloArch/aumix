package state

type X32StateLibsElementElement struct {
	Fpos     X32Int
	Fname    X32String
	Ftype    X32Int
	Fflags   X32Int
	Fhasdata X32Int
}
type X32StateLibsElement struct {
	Index [100]X32StateLibsElementElement `start:"1"`
}

type X32StateLibs struct {
	Fch X32StateLibsElement
	Ffx X32StateLibsElement
	Fr  X32StateLibsElement
}
