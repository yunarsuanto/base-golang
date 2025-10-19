package objects

import "html/template"

type SendChangePasswordOtpMailData struct {
	To             string
	Name           string
	Otp            string
	OtpLifeMinutes int
}

type SendFormVersionApprovalMailData struct {
	To              string
	RoleName        string
	ModuleName      string
	FormName        string
	FormVersionName string
	ReportType      string
	CreatedBy       string
	DetailUrl       string
}

type SendCustomMailData struct {
	To      []string
	Title   string
	Content template.HTML
}
