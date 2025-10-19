package constants

const (
	MorningShiftCode   = "P"
	AfternoonShiftCode = "S"
	NightShiftCode     = "M"
	OffShiftCode       = "O"
)

const (
	OffShiftName = "Off"
)

func SystemInitiatedShiftCode() []string {
	return []string{
		MorningShiftCode,
		AfternoonShiftCode,
		NightShiftCode,
		OffShiftCode,
	}
}
