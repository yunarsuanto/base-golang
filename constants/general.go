package constants

import "time"

type contextKeyType int
type permissionKeyType int
type schemaTypeKeyType int

const (
	AppName = "umkm-go"
)

const (
	MigrationDir = "migrations"
	LogDir       = "./log/"
)

const (
	HeaderAuthorization = "Authorization"
	// HeaderHospitalId                       = "X-Hospital-ID"
	Bearer                                 = "Bearer"
	ClaimsContextKey     contextKeyType    = iota
	PermissionContextKey permissionKeyType = iota
	schemaTypeContextKey schemaTypeKeyType = iota
)

const (
	SuperUserRoleName = "Super User"
)

const (
	MaximumSendBatch = 300
)

const (
	WebPlatform    = "web"
	MobilePlatform = "mobile"
)

const (
	APlusBloodType   = "A+"
	BPlusBloodType   = "B+"
	ABPlusBloodType  = "AB+"
	OPlusBloodType   = "O+"
	AMinusBloodType  = "A-"
	BMinusBloodType  = "B-"
	ABMinusBloodType = "AB-"
	OMinusBloodType  = "O-"
)

const (
	IndonesiaCountry   = "Indonesia"
	IndonesiaAlphaCode = "ID"
	JakartaTimezone    = "Asia/Jakarta"
	Oneday             = time.Second * 3600 * 24
)

const (
	UserSexMale   = "M"
	UserSexFemale = "F"

	IndonesianMale   = "Laki-laki"
	IndonesianFemale = "Perempuan"
)

const (
	AllPathVariable           = "all"
	IdPathVariable            = "id"
	FormIdPathVariable        = "formId"
	FormVersionIdPathVariable = "formVersionId"
	CodePathVariable          = "code"
)

const (
	Assignee = "Dikerjakan"
)

const (
	WibTimezone  = "wib"
	WitaTimezone = "wita"
	WitTimezone  = "wit"
)

const (
	CalendarGroupByDate  = "date"
	CalendarGroupByMonth = "month"
)
