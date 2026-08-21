package state

type X32StateUSBDirElement struct {
	Ftype X32String
	Fname X32String
}

type X32StateUSBDir struct {
	Fdirpos X32Int
	Fmaxpos X32Int
	Index   [999]X32StateUSBDirElement `starts:"1"`
}

func (s *X32StateUSBDir) Get() []any {
	return []any{s.Fdirpos, s.Fmaxpos}
}

func (s *X32StateUSBDir) Set(values ...any) (int, error) {
	return 2, setAll([]X32StateValue{&s.Fdirpos, &s.Fmaxpos}, values)
}

type X32StateUSB struct {
	Fpath  X32String
	Ftitle X32String
	Fdir   X32StateUSBDir
}

func (s *X32StateUSB) Get() []any {
	return []any{s.Fpath, s.Ftitle}
}

func (s *X32StateUSB) Set(values ...any) (int, error) {
	return 2, setAll([]X32StateValue{&s.Fpath, &s.Ftitle}, values)
}
