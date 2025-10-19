package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/yunarsuanto/base-go/constants"
)

const NotificationTemplateDataName = "templat notifikasi"

type GetNotificationTemplate struct {
	Id              string     `db:"id"`
	CreatedAt       time.Time  `db:"created_at"`
	UpdatedAt       *time.Time `db:"updated_at"`
	CreatedById     *string    `db:"created_by_id"`
	UpdatedById     *string    `db:"updated_by_id"`
	CreatedByName   *string    `db:"created_by_name"`
	UpdatedByName   *string    `db:"updated_by_name"`
	Name            string     `db:"name"`
	Title           string     `db:"title"`
	Content         string     `db:"content"`
	ShortContent    *string    `db:"short_content"`
	SystemInitiated bool       `db:"system_initiated"`
	Code            *string    `db:"code"`
	SendEmail       bool       `db:"send_email"`
	SendInApp       bool       `db:"send_in_app"`
	IsActive        bool       `db:"is_active"`
}

type CreateNotificationTemplate struct {
	Id              string         `db:"id"`
	CreatedBy       sql.NullString `db:"created_by"`
	Name            string         `db:"name"`
	Title           string         `db:"title"`
	Content         string         `db:"content"`
	ShortContent    sql.NullString `db:"short_content"`
	SystemInitiated bool           `db:"system_initiated"`
	Code            sql.NullString `db:"code"`
	SendEmail       bool           `db:"send_email"`
	SendInApp       bool           `db:"send_in_app"`
}

type GetNotificationTemplateUser struct {
	Id           string  `db:"id"`
	Name         string  `db:"name"`
	Title        string  `db:"title"`
	Content      string  `db:"content"`
	ShortContent *string `db:"short_content"`
	Code         *string `db:"code"`
	SendEmail    bool    `db:"send_email"`
	SendInApp    bool    `db:"send_in_app"`
	IsActive     bool    `db:"is_active"`
	UserId       string  `db:"user_id"`
	UserName     string  `db:"user_name"`
	UserEmail    string  `db:"user_email"`
	RoleId       string  `db:"role_id"`
	RoleName     string  `db:"role_name"`
	HospitalId   string  `db:"hospital_id"`
	HospitalName string  `db:"hospital_name"`
	ModuleId     string  `db:"module_id"`
	ModuleName   string  `db:"module_name"`
	FormId       string  `db:"form_id"`
	FormName     string  `db:"form_name"`
}

func (GetNotificationTemplate) ColumnQuery() string {
	return `
		nt.id,
		nt.created_at,
		nt.updated_at,
		cu.id AS created_by_id,
		uu.id AS updated_by_id,
		cu.name AS created_by_name,
		uu.name AS updated_by_name,
		nt.name,
		nt.title,
		nt.content,
		nt.short_content,
		nt.system_initiated,
		nt.code,
		nt.send_email,
		nt.send_in_app,
		nt.is_active
	`
}

func (GetNotificationTemplate) TableQuery() string {
	return `
		FROM notification_templates nt
		LEFT JOIN users cu ON cu.id = nt.created_by
		LEFT JOIN users uu ON uu.id = nt.updated_by
	`
}

func (GetNotificationTemplateUser) ColumnQuery(withFormColumns bool) string {
	var addtionalColumns string
	if withFormColumns {
		addtionalColumns = `,
			m.id AS module_id,
			m.name AS module_name,
			f.id AS form_id,
			f.name AS form_name
		`
	}

	return fmt.Sprintf(`
		nt.id,
		nt.name,
		nt.title,
		nt.content,
		nt.short_content,
		nt.code,
		nt.send_email,
		nt.send_in_app,
		nt.is_active,
		u.id AS user_id,
		u.name AS user_name,
		u.email AS user_email,
		r.id AS role_id,
		r.name AS role_name,
		h.id AS hospital_id,
		h.name AS hospital_name
		%s
	`, addtionalColumns)
}

func (GetNotificationTemplateUser) EntryReminderTableQuery() string {
	return fmt.Sprintf(`
		FROM forms f
		JOIN notification_templates nt ON nt.code = '%s'
		JOIN modules m ON m.id = f.module_id
		JOIN form_entry_periods fep ON fep.form_id = f.id
		JOIN form_versions fv ON fv.form_id = f.id
		JOIN form_versions_roles fvr ON fvr.form_version_id = fv.id
		JOIN form_versions_hospitals fvh ON fvh.form_version_id = fv.id
		JOIN form_assignments fa ON fa.form_version_id = fv.id AND fa.hospital_id = fvh.hospital_id AND fa.form_approval_id IS NULL
		JOIN users u ON u.id = fa.user_id
		JOIN users_roles ur ON ur.user_id = u.id
		JOIN roles r ON r.id = ur.role_id
		JOIN users_hospitals uh ON uh.user_id = u.id AND uh.hospital_id = fa.hospital_id
		JOIN hospitals h ON h.id = uh.hospital_id
		LEFT JOIN form_entries fe ON 
			fe.form_entry_period_id = fep.id AND
			fe.form_version_id = fv.id AND
			fe.user_id = u.id AND
			fe.hospital_id = h.id AND
			fe.status IN ('%s', '%s')
		WHERE 
			f.is_active IS true AND 
			fep.start_date <= CURRENT_DATE AND CURRENT_DATE < fep.end_date AND
			fv.status = '%s' AND 
			fe.id IS NULL
	`,
		constants.NotificationTemplateCodeFormEntryReminder,
		constants.FormEntryStatusSubmitted,
		constants.FormEntryStatusApproved,
		constants.FormVersionStatusApproved,
	)
}

// TODO::EntryApprovalReminder
func (GetNotificationTemplateUser) EntryApprovalReminderTableQuery() string {
	return fmt.Sprintf(`
		FROM form_entries fe
		JOIN notification_templates nt ON nt.code = '%s'
		JOIN form_entry_periods fep ON fep.id = fe.form_entry_period_id
		JOIN form_versions fv ON fv.id = fe.form_version_id
		JOIN forms f ON f.id = fv.form_id
		JOIN modules m ON m.id = f.module_id
		JOIN (
			SELECT fea.form_entry_id, MIN(fa.sort) AS min_sort
			FROM form_entry_approvals fea
			JOIN form_approvals fa ON fa.id = fea.form_approval_id
			WHERE fea.status = '%s'
			GROUP BY fea.form_entry_id
		) ma ON ma.form_entry_id = fe.id
		JOIN form_entry_approvals fea ON fea.form_entry_id = fe.id
		JOIN form_approvals fa ON fa.id = fea.form_approval_id AND fa.sort = ma.min_sort
		JOIN form_assignments fas ON 
			fas.form_version_id = fe.form_version_id AND
			fas.hospital_id = fe.hospital_id AND
			fas.form_approval_id = fa.id
		JOIN users u ON u.id = fas.user_id
		JOIN users_roles ur ON ur.user_id = u.id
		JOIN roles r ON r.id = ur.role_id
		JOIN users_hospitals uh ON uh.user_id = u.id AND uh.hospital_id = fas.hospital_id
		JOIN hospitals h ON h.id = uh.hospital_id
		WHERE
			fe.status = '%s' AND
			f.is_active IS true AND 
			fep.end_date <= CURRENT_DATE AND
			fv.status = '%s'
	`,
		constants.NotificationTemplateCodeFormEntryApprovalReminder,
		constants.FormEntryApprovalStatusPending,
		constants.FormEntryStatusSubmitted,
		constants.FormVersionStatusApproved,
	)

	// VERSION 1
	// return fmt.Sprintf(`
	// 	FROM form_entry_approvals fea
	// 	JOIN notification_templates nt ON nt.code = '%s'
	// 	JOIN form_entries fe ON fe.id = fea.form_entry_id
	// 	JOIN form_entry_periods fep ON fep.id = fe.form_entry_period_id
	// 	JOIN (
	// 		SELECT fea2.form_entry_id, MIN(fa.sort) AS min_sort
	// 		FROM form_entry_approvals fea2
	// 		JOIN form_approvals fa ON fa.id = fea2.form_approval_id
	// 		WHERE fea2.status = '%s'
	// 		GROUP BY fea2.form_entry_id
	// 	) mf ON mf.form_entry_id = fea.form_entry_id
	// 	JOIN form_approval fa ON fa.id = fea.form_approval_id AND mf.sort = fa.sort
	// 	WHERE
	// 		fea.status = '%s' AND
	// 		fep.end_date <= CURRENT_DATE

	// `,
	// 	constants.NotificationTemplateCodeFormEntryApprovalReminder,
	// 	constants.FormEntryApprovalStatusPending,
	// 	constants.FormEntryApprovalStatusPending,
	// )

	// REFERENCE
	// return `
	// 	FROM notification_templates nt
	// 	CROSS JOIN users u
	// 	JOIN users_roles ur ON ur.user_id = u.id
	// 	JOIN roles r ON r.id = ur.role_id
	// 	JOIN users_hospitals uh ON uh.user_id = u.id
	// 	JOIN hospitals h ON h.id = uh.hospital_id
	// `
}

func (GetNotificationTemplateUser) CustomTableQuery() string {
	return `
		FROM notification_templates nt
		CROSS JOIN users u
		JOIN users_roles ur ON ur.user_id = u.id
		JOIN roles r ON r.id = ur.role_id
		JOIN users_hospitals uh ON uh.user_id = u.id
		JOIN hospitals h ON h.id = uh.hospital_id
	`
}
