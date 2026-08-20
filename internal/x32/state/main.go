package state

type X32StateMainStMix struct {
	Fon    X32EnumOnType
	Ffader X32Level
	Fpan   X32Float
	Index  [6]X32StateUniMixElement `start:"1"`
}

func (s *X32StateMainStMix) Get() []any {
	return []any{s.Fon, s.Ffader, s.Fpan}
}

func (s *X32StateMainStMix) Set(values ...any) (int, error) {
	return 3, setAll([]X32StateValue{&s.Fon, &s.Ffader, &s.Fpan}, values)
}

type X32StateMainStElement struct {
	Fconfig X32StateUniConfig
	Fdyn    X32StateUniDyn
	Finsert X32StateUniInsert
	Feq     X32StateUniEq6
	Fmix    X32StateMainStMix
}

type X32StateMainMMix struct {
	Fon    X32EnumOnType
	Ffader X32Level
	Index  [6]X32StateUniMixElement `start:"1"`
}

func (s *X32StateMainMMix) Get() []any {
	return []any{s.Fon, s.Ffader}
}

func (s *X32StateMainMMix) Set(values ...any) (int, error) {
	return 3, setAll([]X32StateValue{&s.Fon, &s.Ffader}, values)
}

type X32StateMainMElement struct {
	Fconfig X32StateUniConfig
	Fdyn    X32StateUniDyn
	Finsert X32StateUniInsert
	Feq     X32StateUniEq6
	Fmix    X32StateMainMMix
}

type X32StateMain struct {
	Fst X32StateMainStElement
	Fm  X32StateMainMElement
}
