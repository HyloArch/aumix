package state

type X32StatePrefsRemote struct {
	Enable   X32EnumOnType
	Protocol int32 `enum:""`
	Port     int32
	Ioenable int32 `enum:""`
}
type X32StatePrefsiQ struct {
}

type X32StatePrefs struct {
	Style             X32String
	Bright            X32Float
	Lcdcont           X32Float
	Ledbright         X32Float
	Lamp              X32Float
	Lampon            X32EnumOnType
	Clockrate         X32EnumClockrateType
	Clocksource       X32EnumClocksourceType
	Confirm_general   X32EnumOnType
	Comfirm_overwrite X32EnumOnType
	Confirm_sceneload X32EnumOnType
	Viewrtn           X32EnumOnType
	Selfollowbank     X32EnumOnType
	Sceneadvance      X32EnumOnType
	Safe_masterlevels X32EnumOnType
	Haflags           X32Int
	Autosel           X32EnumOnType
	Show_control      X32EnumShowControlType
	Clockmode         X32EnumClockModeType
	Hardmute          X32EnumOnType
	Dcamute           X32EnumOnType
	Invertmutes       X32EnumInvType
	Remote            X32StatePrefsRemote
	IQ                X32StatePrefsiQ
}
