package state

type X32StatePrefsRemote struct {
	Enable   X32EnumOnType
	Protocol X32EnumProtocolType
	Port     X32EnumPortType
	Ioenable X32Int
}

type X32StatePrefsiQElement struct {
	FiQmodel X32EnumIQModelType
	FiQeqset X32EnumIQEqType
	FiQsound X32Int
}

type X32StatePrefsiQ struct {
	Index [16]X32StatePrefsiQElement `start:"1"`
}

type X32StatePrefsCard struct {
	FUFifc    X32EnumUFifcType
	FUFmode   X32EnumUFmodeType
	FUSBmode  X32EnumUSBmodeType
	FADATwc   X32EnumADATwcType
	FADATsync X32EnumADATsyncType
	FMADImode X32EnumMADImodeType
	FMADIin   X32EnumMADIType
	// FMADIout  X32EnumMADIType
	FMADI_   X32EnumOnType
	FMADIsrc X32EnumMADIsrcType
}

type X32StatePrefsRta struct {
	Fvisibility X32EnumRTAVisibilityType
	Fgain       X32Float
	Fautogain   X32EnumOnType
	Fsource     X32EnumRTASourceType
	Fpos        X32EnumPosType
	Fmode       X32EnumRTAModeType
	Foption     X32Int
	Fdet        X32EnumDetType
	Fdecay      X32Float
	Fpeakhold   X32EnumRTAPeakholdType
}

type X32StatePrefsIPAddr struct {
	Index [4]X32Int
}

type X32StatePrefsIp struct {
	Fdhcp    X32EnumOnType
	Faddr    X32StatePrefsIPAddr
	Fmask    X32StatePrefsIPAddr
	Fgateway X32StatePrefsIPAddr
}

func (s *X32StatePrefsIp) Get() []any {
	return []any{s.Fdhcp}
}

func (s *X32StatePrefsIp) Set(values ...any) (int, error) {
	return 1, setAll([]X32StateValue{&s.Fdhcp}, values)
}

type X32StatePrefs struct {
	Fstyle             X32String
	Fbright            X32Float
	Flcdcont           X32Float
	Fledbright         X32Float
	Flamp              X32Float
	Flampon            X32EnumOnType
	Fclockrate         X32EnumClockrateType
	Fclocksource       X32EnumClocksourceType
	Fconfirm_general   X32EnumOnType
	Fconfirm_overwrite X32EnumOnType
	Fconfirm_sceneload X32EnumOnType
	Fviewrtn           X32EnumOnType
	Fselfollowbank     X32EnumOnType
	Fsceneadvance      X32EnumOnType
	Fsafe_masterlevels X32EnumOnType
	Fhaflags           X32Int
	Fautosel           X32EnumOnType
	Fshow_control      X32EnumShowControlType
	Fclockmode         X32EnumClockModeType
	Fhardmute          X32EnumOnType
	Fdcamute           X32EnumOnType
	Finvertmutes       X32EnumInvType
	Fremote            X32StatePrefsRemote
	FiQ                X32StatePrefsiQ
	Fcard              X32StatePrefsCard
	Frta               X32StatePrefsRta
	Fip                X32StatePrefsIp
}

func (s *X32StatePrefs) Get() []any {
	return []any{s.Fstyle, s.Fbright, s.Flcdcont, s.Fledbright, s.Flamp, s.Flampon, s.Fclockrate, s.Fclocksource, s.Fconfirm_general, s.Fconfirm_overwrite, s.Fconfirm_sceneload, s.Fviewrtn, s.Fselfollowbank, s.Fsceneadvance, s.Fsafe_masterlevels, s.Fhaflags, s.Fautosel, s.Fshow_control, s.Fclockmode, s.Fhardmute, s.Fdcamute, s.Finvertmutes}
}

func (s *X32StatePrefs) Set(values ...any) (int, error) {
	return 22, setAll([]X32StateValue{&s.Fstyle, &s.Fbright, &s.Flcdcont, &s.Fledbright, &s.Flamp, &s.Flampon, &s.Fclockrate, &s.Fclocksource, &s.Fconfirm_general, &s.Fconfirm_overwrite, &s.Fconfirm_sceneload, &s.Fviewrtn, &s.Fselfollowbank, &s.Fsceneadvance, &s.Fsafe_masterlevels, &s.Fhaflags, &s.Fautosel, &s.Fshow_control, &s.Fclockmode, &s.Fhardmute, &s.Fdcamute, &s.Finvertmutes}, values)
}
