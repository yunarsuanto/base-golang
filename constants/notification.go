package constants

type NotificationTemplateData struct {
	Name              string
	Value             string
	AllowOnCustom     bool
	DatabaseReference string
}

const (
	NotificationTemplateCodeFormEntryReminder         = "form_entry_reminder"
	NotificationTemplateCodeFormEntryOverdue          = "form_entry_overdue"
	NotificationTemplateCodeFormEntryApprovalReminder = "form_entry_approval_reminder"
)

func NotificationTemplateRole() NotificationTemplateData {
	return NotificationTemplateData{Name: "Role", Value: "{{ ROLE }}", AllowOnCustom: true, DatabaseReference: "r.name"}
}

func NotificationTemplateModule() NotificationTemplateData {
	return NotificationTemplateData{Name: "Modul", Value: "{{ MODULE }}", AllowOnCustom: false, DatabaseReference: "m.name"}
}

func NotificationTemplateForm() NotificationTemplateData {
	return NotificationTemplateData{Name: "Nama Formulir", Value: "{{ FORM }}", AllowOnCustom: false, DatabaseReference: "f.name"}
}

func NotificationTemplateUser() NotificationTemplateData {
	return NotificationTemplateData{Name: "Nama User", Value: "{{ USER }}", AllowOnCustom: true, DatabaseReference: "u.name"}
}

func NotificationTemplateDate() NotificationTemplateData {
	return NotificationTemplateData{Name: "Tanggal", Value: "{{ DATE }}", AllowOnCustom: true, DatabaseReference: "DATE_FORMAT(CURRENT_DATE, '%d %M %Y')"}
}
func NotificationTemplateTime() NotificationTemplateData {
	return NotificationTemplateData{Name: "Jam", Value: "{{ TIME }}", AllowOnCustom: true, DatabaseReference: "DATE_FORMAT(CURRENT_TIME, '%H:%i')"}
}

func NotificationTemplateHospital() NotificationTemplateData {
	return NotificationTemplateData{Name: "Rumah Sakit", Value: "{{ HOSPITAL }}", AllowOnCustom: true, DatabaseReference: "h.name"}
}
