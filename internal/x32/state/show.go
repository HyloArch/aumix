package state

type X32StateShowPrepos struct {
	Fcurrent X32Int
}
type X32StateShowfileShow struct {
	Fname     X32String
	Finputs   X32Int
	Fmxsends  X32Int
	Fmxbuses  X32Int
	Fconsole  X32Int
	Fchan16   X32Int
	Fchan32   X32Int
	Freturn   X32Int
	Fbuses    X32Int
	Flrmtxdca X32Int
	Feffects  X32Int
}
type X32StateShowfileCueElement struct {
	Fnumb      X32Int
	Fname      X32String
	Fskip      X32Int
	Fscene     X32Int
	Fbit       X32Int
	Fmiditype  X32Int
	Fmidichan  X32Int
	Fmidipara1 X32Int
	Fmidipara2 X32Int
}
type X32StateShowfileCue struct {
	Index [100]X32StateShowfileCueElement
}
type X32StateShowfileSceneElement struct {
	Fname    X32String
	Fnotes   X32String
	Fsafes   X32Int
	Fhasdata X32Int
}
type X32StateShowfileScene struct {
	Index [100]X32StateShowfileSceneElement
}
type X32StateShowfileSnippetElement struct {
	Fname     X32String
	Feventtyp X32Int
	Fchannels X32Int
	Fauxbuses X32Int
	Fmaingrps X32Int
	Fhasdata  X32Int
}
type X32StateShowfileSnippet struct {
	Index [100]X32StateShowfileSnippetElement
}
type X32StateShowfile struct {
	Fshow    X32StateShowfileShow
	Fcue     X32StateShowfileCue
	Fscene   X32StateShowfileScene
	Fsnippet X32StateShowfileSnippet
}

type X32StateShow struct {
	Fprepos   X32StateShowPrepos
	Fshowfile X32StateShowfile
}
