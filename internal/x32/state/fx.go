package state

import "fmt"

type X32StateFxSource struct {
	Fl X32EnumFxSourceType
	Fr X32EnumFxSourceType
}

type X32StateFxPar struct {
	Index [64]X32Float `start:"1"`
}

func (s *X32StateFxPar) Get() []any {
	out := make([]any, 64)
	for i, v := range s.Index {
		out[i] = v
	}
	return out
}

func (s *X32StateFxPar) Set(values ...any) (int, error) {
	for i, v := range values {
		floatV, ok := v.(float32)
		if !ok {
			return 0, fmt.Errorf("Provided value isn't of type float")
		}
		s.Index[i] = X32Float(floatV)
	}
	return max(len(values), 64), nil
}

type X32StateFxSourceElement struct {
	Ftype   X32EnumFxTypeSourcedType
	Fsource X32StateFxSource
	Fpar    X32StateFxPar
}

func (s *X32StateFxSourceElement) Get() []any {
	return []any{s.Ftype}
}

func (s *X32StateFxSourceElement) Set(values ...any) (int, error) {
	return 3, setAll([]X32StateValue{&s.Ftype}, values)
}

type X32StateFxElement struct {
	Ftype X32EnumFxTypeType
	Fpar  X32StateFxPar
}

func (s *X32StateFxElement) Get() []any {
	return []any{s.Ftype}
}

func (s *X32StateFxElement) Set(values ...any) (int, error) {
	return 3, setAll([]X32StateValue{&s.Ftype}, values)
}

type X32StateFx struct {
	F1 X32StateFxSourceElement
	F2 X32StateFxSourceElement
	F3 X32StateFxSourceElement
	F4 X32StateFxSourceElement
	F5 X32StateFxElement
	F6 X32StateFxElement
	F7 X32StateFxElement
	F8 X32StateFxElement
}
