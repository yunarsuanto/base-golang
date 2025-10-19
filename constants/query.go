package constants

const (
	DefaultLimit     = 20
	DefaultPage      = 1
	DefaultUnlimited = ^uint32(0)
)

const (
	Ascending          = "ASC"
	Descending         = "DESC"
	DescendingNullLast = "DESC NULLS LAST"
	Random             = "RANDOM()"
)

func ValidOrderValue() []string {
	return []string{
		Ascending,
		Descending,
		DescendingNullLast,
		Random,
	}
}

const (
	DefaultPassword = "UPDATEPASSWORDTOLOGIN"
)
