package objects

import "time"

type ListNotificationRequest struct {
	UserId     string
	HospitalId string
}

type ListNotification struct {
	Id                  string
	CreatedAt           time.Time
	Title               string
	Content             string
	ShortContent        *string
	Context             string
	IsRead              bool
	HospitalId          *string
	FormId              *string
	FormVersionId       *string
	FormEntryId         *string
	FormEntryApprovalId *string
	FormEntryPeriodId   *string
}

type DetailNotificationRequest struct {
	Id     string
	UserId string
}

type DetailNotification struct {
	Id                     string
	CreatedAt              time.Time
	UserId                 string
	UserName               string
	UserProfilePictureUrl  string
	Title                  string
	Content                string
	ShortContent           *string
	Context                string
	IsRead                 bool
	HospitalId             *string
	HospitalName           *string
	FormId                 *string
	FormName               *string
	FormVersionId          *string
	FormVersionName        *string
	FormEntryId            *string
	FormEntryCode          *string
	FormEntryApprovalId    *string
	FormEntryApprovalLabel *string
	FormEntryPeriodId      *string
	FormEntryPeriodCode    *string
}

type CountUnreadNotification struct {
	Count int
}

type ParsedNotificationContent struct {
	Title        string
	Content      string
	ShortContent string
}

type SendNotification struct {
	UserIds []string
	Title   string
	Body    string
}
