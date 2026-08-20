package state

type X32StateDcaElement struct {
	Fon     X32EnumOnType
	Ffader  X32Level
	Fconfig X32StateUniConfig
}

func (s *X32StateDcaElement) Get() []any {
	return []any{s.Fon, s.Ffader}
}

func (s *X32StateDcaElement) Set(values ...any) (int, error) {
	return 2, setAll([]X32StateValue{&s.Fon, &s.Ffader}, values)
}

type X32StateDca struct {
	Index [8]X32StateDcaElement `start:"1"`
}
