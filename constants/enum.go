package constants

const (
	FormConditionOperatorAnd = "and"
	FormConditionOperatorOr  = "or"
	FormConditionEqual       = "="
	FormConditionNotEqual    = "!="
	FormConditionLessThan    = "<"
	FormConditionMoreThan    = ">"
	FormConditionLessOrEqual = "<="
	FormConditionMoreOrEqual = ">="
)

func ValidConditionGroupOperators() []string {
	return []string{
		FormConditionOperatorAnd,
		FormConditionOperatorOr,
	}
}

func ValidConditionOperators() []string {
	return []string{
		FormConditionEqual,
		FormConditionNotEqual,
		FormConditionLessThan,
		FormConditionMoreThan,
		FormConditionLessOrEqual,
		FormConditionMoreOrEqual,
	}
}

const (
	FormConditionLogicActionShow    = "show"
	FormConditionLogicActionHide    = "hide"
	FormConditionLogicActionEnable  = "enable"
	FormConditionLogicActionDisable = "disable"
)

const (
	FormReportTypeDaily        = "daily"
	FormReportTypeWeekly       = "weekly"
	FormReportTypeMonthly      = "monthly"
	FormReportTypeEvery2Months = "every_2_months"
	FormReportTypeEvery6Months = "every_6_months"
	FormReportTypeAnnual       = "annual"
	FormReportTypeEvent        = "event"
)

func FormReportTypes() []string {
	return []string{
		FormReportTypeDaily,
		FormReportTypeWeekly,
		FormReportTypeMonthly,
		FormReportTypeEvery2Months,
		FormReportTypeEvery6Months,
		FormReportTypeAnnual,
		FormReportTypeEvent,
	}
}

const (
	FormVersionStatusDraft           = "draft"
	FormVersionStatusPendingApproval = "pending_approval"
	FormVersionStatusApproved        = "approved"
	FormVersionStatusDeclined        = "declined"
)

const (
	FormInputTypeNone          = "none"
	FormInputTypeMatrix        = "matrix"
	FormInputTypeDate          = "date"
	FormInputTypeTime          = "time"
	FormInputTypeDatetime      = "datetime"
	FormInputTypeVarchar       = "varchar"
	FormInputTypeText          = "text"
	FormInputTypeUser          = "user"
	FormInputTypeInteger       = "integer"
	FormInputTypeDecimal       = "decimal"
	FormInputTypeStatic        = "static"
	FormInputTypeEmail         = "email"
	FormInputTypePhone         = "phone"
	FormInputTypeIdNumber      = "id_number"
	FormInputTypePostalCode    = "postal_code"
	FormInputTypeVehicleNumber = "vehicle_number"
	FormInputTypeFacility      = "facility"
	FormInputTypeRadio         = "radio"
	FormInputTypeCheckbox      = "checkbox"
	FormInputTypeSelect        = "select"
	FormInputTypeAutoComplete  = "autocomplete"
	FormInputTypeFile          = "file"
)

func EnumFormInputType() []string {
	return []string{
		FormInputTypeRadio,
		FormInputTypeCheckbox,
		FormInputTypeSelect,
		FormInputTypeAutoComplete,
	}
}

const (
	FormEntryStatusDraft     = "draft"
	FormEntryStatusSubmitted = "submitted"
	FormEntryStatusApproved  = "approved"
	FormEntryStatusDeclined  = "declined"
)

const (
	FormEntryFieldValueTypeText     = "text"
	FormEntryFieldValueTypeNumber   = "number"
	FormEntryFieldValueTypeDate     = "date"
	FormEntryFieldValueTypeTime     = "time"
	FormEntryFieldValueTypeDatetime = "datetime"
	FormEntryFieldValueTypeBoolean  = "boolean"
	FormEntryFieldValueTypeEnum     = "enum"
	FormEntryFieldValueTypeUser     = "user"
	FormEntryFieldValueTypeFacility = "facility"
)

const (
	FormEntryApprovalStatusPending  = "pending"
	FormEntryApprovalStatusApproved = "approved"
	FormEntryApprovalStatusDeclined = "declined"
)

const (
	FormFileTypePng  = "png"
	FormFileTypeJpg  = "jpg"
	FormFileTypePdf  = "pdf"
	FormFileTypePptx = "pptx"
	FormFileTypeDocx = "docx"
	FormFileTypeTxt  = "txt"
	FormFileTypeXslx = "xslx"
	FormFileTypeCsv  = "csv"
)

func EnumFormFileType() []string {
	return []string{
		FormFileTypePng,
		FormFileTypeJpg,
		FormFileTypePdf,
		FormFileTypePptx,
		FormFileTypeDocx,
		FormFileTypeTxt,
		FormFileTypeXslx,
		FormFileTypeCsv,
	}
}

const (
	FormDateDefaultValueCurrentDate     = "CURRENT_DATE"
	FormDateDefaultValueCurrentDateTime = "CURRENT_DATETIME"
	FormDateDefaultValueCurrentTime     = "CURRENT_TIME"
)

const (
	FormEntryLogActionCreate  = "create"
	FormEntryLogActionUpdate  = "update"
	FormEntryLogActionDelete  = "delete"
	FormEntryLogActionSubmit  = "submit"
	FormEntryLogActionApprove = "approve"
	FormEntryLogActionDecline = "decline"
)
