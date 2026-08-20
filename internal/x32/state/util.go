package state

import "fmt"

type X32StateValue interface {
	Get() []any
	Set(values ...any) (int, error)
}

type X32Int int32

func (i *X32Int) Get() []any {
	return []any{*i}
}

func (i *X32Int) Set(values ...any) (int, error) {
	if len(values) < 1 {
		return 0, fmt.Errorf("Values must be of at least size 1")
	}
	switch v := values[0].(type) {
	case int32:
		*i = X32Int(v)
	case float32:
		*i = X32Int(v)
	default:
		return 0, fmt.Errorf("Can't convert value to int")
	}
	return 1, nil
}

type X32Float float32

func (f *X32Float) Get() []any {
	return []any{*f}
}

func (f *X32Float) Set(values ...any) (int, error) {
	if len(values) < 1 {
		return 0, fmt.Errorf("Values must be of at least size 1")
	}
	v, ok := values[0].(float32)
	if !ok {
		return 0, fmt.Errorf("Can't convert value to float")
	}
	*f = X32Float(v)
	return 1, nil
}

type X32Level float32

func (l *X32Level) Get() []any {
	return []any{*l}
}

func (l *X32Level) Set(values ...any) (int, error) {
	if len(values) < 1 {
		return 0, fmt.Errorf("Values must be of at least size 1")
	}
	switch v := values[0].(type) {
	case float32:
		*l = X32Level(v)
	case string:

	default:
		return 0, fmt.Errorf("Provided type is not of type float or string")
	}
	return 1, nil
}

type X32String string

func (s *X32String) Get() []any {
	return []any{*s}
}

func (s *X32String) Set(values ...any) (int, error) {
	if len(values) < 1 {
		return 0, fmt.Errorf("Values must be of at least size 1")
	}
	v, ok := values[0].(string)
	if !ok {
		return 0, fmt.Errorf("Can't convert value to string")
	}
	*s = X32String(v)
	return 1, nil
}

func setAll(fields []X32StateValue, values []any) error {
	if len(values) < len(fields) {
		return fmt.Errorf("Not enough values provided")
	}

	index := 0
	for _, f := range fields {
		consumed, err := f.Set(values[index:]...)
		if err != nil {
			return err
		}
		index += consumed
	}
	return nil
}
