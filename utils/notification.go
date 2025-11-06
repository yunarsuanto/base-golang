package utils

import (
	"strings"
	"time"

	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/objects"
)

func NotificationTemplateReferences(systemInitiated bool) []constants.NotificationTemplateData {
	var result []constants.NotificationTemplateData

	if systemInitiated {
		result = []constants.NotificationTemplateData{
			constants.NotificationTemplateRole(),
			constants.NotificationTemplateModule(),
			constants.NotificationTemplateForm(),
			constants.NotificationTemplateUser(),
			constants.NotificationTemplateDate(),
			constants.NotificationTemplateTime(),
			constants.NotificationTemplateHospital(),
		}
	} else {
		if constants.NotificationTemplateRole().AllowOnCustom {
			result = append(result, constants.NotificationTemplateRole())
		}
		if constants.NotificationTemplateModule().AllowOnCustom {
			result = append(result, constants.NotificationTemplateModule())
		}
		if constants.NotificationTemplateForm().AllowOnCustom {
			result = append(result, constants.NotificationTemplateForm())
		}
		if constants.NotificationTemplateUser().AllowOnCustom {
			result = append(result, constants.NotificationTemplateUser())
		}
		if constants.NotificationTemplateDate().AllowOnCustom {
			result = append(result, constants.NotificationTemplateDate())
		}
		if constants.NotificationTemplateTime().AllowOnCustom {
			result = append(result, constants.NotificationTemplateTime())
		}
		if constants.NotificationTemplateHospital().AllowOnCustom {
			result = append(result, constants.NotificationTemplateHospital())
		}
	}

	return result
}

func ParseNotification(data models.GetNotificationTemplateUser) objects.ParsedNotificationContent {
	now := time.Now()
	var result objects.ParsedNotificationContent
	title := data.Title
	content := data.Content
	shortContent := NullScan(data.ShortContent)

	roleRule := constants.NotificationTemplateRole()
	userRule := constants.NotificationTemplateUser()
	dateRule := constants.NotificationTemplateDate()
	timeRule := constants.NotificationTemplateTime()
	// hospitalRule := constants.NotificationTemplateHospital()
	// moduleRule := constants.NotificationTemplateModule()
	// formRule := constants.NotificationTemplateForm()

	title = strings.ReplaceAll(title, roleRule.Value, data.RoleName)
	content = strings.ReplaceAll(content, roleRule.Value, data.RoleName)
	shortContent = strings.ReplaceAll(shortContent, roleRule.Value, data.RoleName)

	title = strings.ReplaceAll(title, userRule.Value, data.UserName)
	content = strings.ReplaceAll(content, userRule.Value, data.UserName)
	shortContent = strings.ReplaceAll(shortContent, userRule.Value, data.UserName)

	title = strings.ReplaceAll(title, dateRule.Value, now.Format("01-02-2006"))
	content = strings.ReplaceAll(content, dateRule.Value, now.Format("01-02-2006"))
	shortContent = strings.ReplaceAll(shortContent, dateRule.Value, now.Format("01-02-2006"))

	title = strings.ReplaceAll(title, timeRule.Value, now.Format(constants.HourMinuteOnly))
	content = strings.ReplaceAll(content, timeRule.Value, now.Format(constants.HourMinuteOnly))
	shortContent = strings.ReplaceAll(shortContent, timeRule.Value, now.Format(constants.HourMinuteOnly))

	// title = strings.ReplaceAll(title, hospitalRule.Value, data.HospitalName)
	// content = strings.ReplaceAll(content, hospitalRule.Value, data.HospitalName)
	// shortContent = strings.ReplaceAll(shortContent, hospitalRule.Value, data.HospitalName)

	// title = strings.ReplaceAll(title, moduleRule.Value, data.ModuleName)
	// content = strings.ReplaceAll(content, moduleRule.Value, data.ModuleName)
	// shortContent = strings.ReplaceAll(shortContent, moduleRule.Value, data.ModuleName)

	// title = strings.ReplaceAll(title, formRule.Value, data.FormName)
	// content = strings.ReplaceAll(content, formRule.Value, data.FormName)
	// shortContent = strings.ReplaceAll(shortContent, formRule.Value, data.FormName)

	result = objects.ParsedNotificationContent{
		Title:        title,
		Content:      content,
		ShortContent: shortContent,
	}

	return result
}
