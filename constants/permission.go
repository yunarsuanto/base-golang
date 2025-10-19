package constants

const (
	PermissionCategoryAccess                  = "Akses"
	PermissionCategoryWebCluster              = "Web App: Klaster"
	PermissionCategoryWebHospital             = "Web App: Rumah Sakit"
	PermissionCategoryWebHospitalFacility     = "Web App: Fasilitas Rumah Sakit"
	PermissionCategoryWebMasterFacility       = "Web App: Master Data Fasilitas Rumah Sakit"
	PermissionCategoryWebDivision             = "Web App: Divisi"
	PermissionCategoryWebPosition             = "Web App: Jabatan"
	PermissionCategoryWebEmployee             = "Web App: Karyawan"
	PermissionCategoryWebEmployeeSchedule     = "Web App: Jadwal Karyawan"
	PermissionCategoryWebWorkShift            = "Web App: Shift"
	PermissionCategoryWebMenu                 = "Web App: Menu"
	PermissionCategoryWebForm                 = "Web App: Formulir"
	PermissionCategoryWebFormEntry            = "Web App: Isian Formulir"
	PermissionCategoryWebFormTemplate         = "Web App: Templat Formulir"
	PermissionCategoryWebFormVersion          = "Web App: Versi Formulir"
	PermissionCategoryWebModule               = "Web App: Modul"
	PermissionCategoryWebNotificationTemplate = "Web App: Templat Notifikasi"
	PermissionCategoryWebPermission           = "Web App: Izin"
	PermissionCategoryWebRole                 = "Web App: Peran"
	PermissionCategoryWebUser                 = "Web App: Pengguna"
)

type permissionCategory struct {
	Name            string
	PermissionCodes []string
}

func PermissionCategories() []permissionCategory {
	return []permissionCategory{
		{
			Name: PermissionCategoryAccess,
			PermissionCodes: []string{
				PermissionAccessMobile,
				PermissionAccessWeb,
			},
		},
		{
			Name: PermissionCategoryWebCluster,
			PermissionCodes: []string{
				PermissionWebClusterCreate,
				PermissionWebClusterDelete,
				PermissionWebClusterGet,
				PermissionWebClusterUpdate,
				PermissionWebClusterUpdateActivation,
			},
		},
		{
			Name: PermissionCategoryWebForm,
			PermissionCodes: []string{
				PermissionWebFormAssign,
				PermissionWebFormCreate,
				PermissionWebFormDelete,
				PermissionWebFormGet,
				PermissionWebFormUpdate,
			},
		},
		{
			Name: PermissionCategoryWebFormEntry,
			PermissionCodes: []string{
				PermissionWebFormEntryGet,
			},
		},
		{
			Name: PermissionCategoryWebFormTemplate,
			PermissionCodes: []string{
				PermissionWebFormTemplateCreate,
				PermissionWebFormTemplateDelete,
				PermissionWebFormTemplateGet,
				PermissionWebFormTemplateUpdate,
			},
		},
		{
			Name: PermissionCategoryWebFormVersion,
			PermissionCodes: []string{
				PermissionWebFormVersionApprove,
				PermissionWebFormVersionCreate,
				PermissionWebFormVersionDelete,
				PermissionWebFormVersionGet,
				PermissionWebFormVersionUpdate,
			},
		},
		{
			Name: PermissionCategoryWebHospital,
			PermissionCodes: []string{
				PermissionWebHospitalCreate,
				PermissionWebHospitalDelete,
				PermissionWebHospitalGet,
				PermissionWebHospitalUpdate,
				PermissionWebHospitalUpdateActivation,
				PermissionWebHospitalUpdateFacility,
			},
		},
		{
			Name: PermissionCategoryWebHospitalFacility,
			PermissionCodes: []string{
				PermissionWebHospitalBuildingCreate,
				PermissionWebHospitalBuildingDelete,
				PermissionWebHospitalBuildingGet,
				PermissionWebHospitalBuildingUpdate,
				PermissionWebHospitalBuildingUpdateActivation,
				PermissionWebHospitalLevelCreate,
				PermissionWebHospitalLevelDelete,
				PermissionWebHospitalLevelGet,
				PermissionWebHospitalLevelUpdate,
				PermissionWebHospitalLevelUpdateActivation,
				PermissionWebHospitalAreaCreate,
				PermissionWebHospitalAreaDelete,
				PermissionWebHospitalAreaGet,
				PermissionWebHospitalAreaUpdate,
				PermissionWebHospitalAreaUpdateActivation,
				PermissionWebHospitalRoomCreate,
				PermissionWebHospitalRoomDelete,
				PermissionWebHospitalRoomGet,
				PermissionWebHospitalRoomUpdate,
				PermissionWebHospitalRoomUpdateActivation,
			},
		},
		{
			Name: PermissionCategoryWebMasterFacility,
			PermissionCodes: []string{
				PermissionWebMasterBuildingCreate,
				PermissionWebMasterBuildingDelete,
				PermissionWebMasterBuildingGet,
				PermissionWebMasterBuildingUpdate,
				PermissionWebMasterBuildingUpdateActivation,
				PermissionWebMasterLevelCreate,
				PermissionWebMasterLevelDelete,
				PermissionWebMasterLevelGet,
				PermissionWebMasterLevelUpdate,
				PermissionWebMasterLevelUpdateActivation,
				PermissionWebMasterAreaCreate,
				PermissionWebMasterAreaDelete,
				PermissionWebMasterAreaGet,
				PermissionWebMasterAreaUpdate,
				PermissionWebMasterAreaUpdateActivation,
				PermissionWebMasterRoomCreate,
				PermissionWebMasterRoomDelete,
				PermissionWebMasterRoomGet,
				PermissionWebMasterRoomUpdate,
				PermissionWebMasterRoomUpdateActivation,
			},
		},
		{
			Name: PermissionCategoryWebDivision,
			PermissionCodes: []string{
				PermissionWebDivisionCreate,
				PermissionWebDivisionDelete,
				PermissionWebDivisionGet,
				PermissionWebDivisionUpdate,
				PermissionWebDivisionUpdateActivation,
			},
		},
		{
			Name: PermissionCategoryWebPosition,
			PermissionCodes: []string{
				PermissionWebPositionCreate,
				PermissionWebPositionDelete,
				PermissionWebPositionGet,
				PermissionWebPositionUpdate,
				PermissionWebPositionUpdateActivation,
			},
		},
		{
			Name: PermissionCategoryWebEmployee,
			PermissionCodes: []string{
				PermissionWebEmployeeCreate,
				PermissionWebEmployeeDelete,
				PermissionWebEmployeeGet,
				PermissionWebEmployeeImport,
				PermissionWebEmployeeUpdate,
			},
		},
		{
			Name: PermissionCategoryWebEmployeeSchedule,
			PermissionCodes: []string{
				PermissionWebEmployeeScheduleCreate,
				PermissionWebEmployeeScheduleDelete,
				PermissionWebEmployeeScheduleGet,
				PermissionWebEmployeeScheduleImport,
				PermissionWebEmployeeScheduleUpdate,
			},
		},
		{
			Name: PermissionCategoryWebWorkShift,
			PermissionCodes: []string{
				PermissionWebWorkShiftCreate,
				PermissionWebWorkShiftDelete,
				PermissionWebWorkShiftGet,
				PermissionWebWorkShiftUpdate,
				PermissionWebWorkShiftUpdateActivation,
			},
		},
		{
			Name: PermissionCategoryWebMenu,
			PermissionCodes: []string{
				PermissionWebMenuGizi,
				PermissionWebMenuKeamanan,
				PermissionWebMenuKebersihan,
				PermissionWebMenuLaundry,
				PermissionWebMenuMaintenanceMedis,
				PermissionWebMenuMaintenanceUmum,
				PermissionWebMenuOperatorTelepon,
				PermissionWebMenuTransport,
				PermissionWebMenuUmum,
				PermissionWebMenuMasterRumahSakit,
				PermissionWebMenuMasterUnitFasilitas,
				PermissionWebMenuMasterJenisLaporan,
				PermissionWebMenuMasterKaryawan,
				PermissionWebMenuMasterShift,
				PermissionWebMenuReport,
				PermissionWebMenuFormManagement,
				PermissionWebMenuUserManagement,
				PermissionWebMenuAssignmentManagement,
				PermissionWebMenuNotificationManagement,
				PermissionWebMenuSetting,
			},
		},
		{
			Name: PermissionCategoryWebModule,
			PermissionCodes: []string{
				PermissionWebModuleGet,
			},
		},
		{
			Name: PermissionCategoryWebNotificationTemplate,
			PermissionCodes: []string{
				PermissionWebNotificationTemplateCreate,
				PermissionWebNotificationTemplateDelete,
				PermissionWebNotificationTemplateGet,
				PermissionWebNotificationTemplateUpdate,
			},
		},
		{
			Name: PermissionCategoryWebPermission,
			PermissionCodes: []string{
				PermissionWebPermissionGet,
			},
		},
		{
			Name: PermissionCategoryWebRole,
			PermissionCodes: []string{
				PermissionWebRoleCreate,
				PermissionWebRoleUpdate,
				PermissionWebRoleGet,
				PermissionWebRoleDelete,
			},
		},
		{
			Name: PermissionCategoryWebUser,
			PermissionCodes: []string{
				PermissionWebUserCreate,
				PermissionWebUserUpdate,
				PermissionWebUserUpdateActivation,
				PermissionWebUserUpdatePassword,
				PermissionWebUserGet,
				PermissionWebUserDelete,
			},
		},
	}
}

const (
	PermissionAccessMobile = "access_mobile"
	PermissionAccessWeb    = "access_web"

	PermissionWebClusterCreate           = "web_cluster_create"
	PermissionWebClusterDelete           = "web_cluster_delete"
	PermissionWebClusterGet              = "web_cluster_get"
	PermissionWebClusterUpdate           = "web_cluster_update"
	PermissionWebClusterUpdateActivation = "web_cluster_update_activation"

	PermissionWebDivisionCreate           = "web_division_create"
	PermissionWebDivisionDelete           = "web_division_delete"
	PermissionWebDivisionGet              = "web_division_get"
	PermissionWebDivisionUpdate           = "web_division_update"
	PermissionWebDivisionUpdateActivation = "web_division_update_activation"

	PermissionWebFormAssign = "web_form_assign"
	PermissionWebFormCreate = "web_form_create"
	PermissionWebFormDelete = "web_form_delete"
	PermissionWebFormGet    = "web_form_get"
	PermissionWebFormUpdate = "web_form_update"

	PermissionWebFormEntryGet = "web_form_entry_get"

	PermissionWebFormTemplateCreate = "web_form_template_create"
	PermissionWebFormTemplateDelete = "web_form_template_delete"
	PermissionWebFormTemplateGet    = "web_form_template_get"
	PermissionWebFormTemplateUpdate = "web_form_template_update"

	PermissionWebFormVersionApprove = "web_form_version_approve"
	PermissionWebFormVersionCreate  = "web_form_version_create"
	PermissionWebFormVersionDelete  = "web_form_version_delete"
	PermissionWebFormVersionGet     = "web_form_version_get"
	PermissionWebFormVersionUpdate  = "web_form_version_update"

	PermissionWebHospitalCreate           = "web_hospital_create"
	PermissionWebHospitalDelete           = "web_hospital_delete"
	PermissionWebHospitalGet              = "web_hospital_get"
	PermissionWebHospitalUpdate           = "web_hospital_update"
	PermissionWebHospitalUpdateActivation = "web_hospital_update_activation"
	PermissionWebHospitalUpdateFacility   = "web_hospital_update_facility"

	PermissionWebHospitalBuildingCreate           = "web_hospital_building_create"
	PermissionWebHospitalBuildingDelete           = "web_hospital_building_delete"
	PermissionWebHospitalBuildingGet              = "web_hospital_building_get"
	PermissionWebHospitalBuildingUpdate           = "web_hospital_building_update"
	PermissionWebHospitalBuildingUpdateActivation = "web_hospital_building_update_activation"

	PermissionWebHospitalLevelCreate           = "web_hospital_level_create"
	PermissionWebHospitalLevelDelete           = "web_hospital_level_delete"
	PermissionWebHospitalLevelGet              = "web_hospital_level_get"
	PermissionWebHospitalLevelUpdate           = "web_hospital_level_update"
	PermissionWebHospitalLevelUpdateActivation = "web_hospital_level_update_activation"

	PermissionWebHospitalAreaCreate           = "web_hospital_area_create"
	PermissionWebHospitalAreaDelete           = "web_hospital_area_delete"
	PermissionWebHospitalAreaGet              = "web_hospital_area_get"
	PermissionWebHospitalAreaUpdate           = "web_hospital_area_update"
	PermissionWebHospitalAreaUpdateActivation = "web_hospital_area_update_activation"

	PermissionWebHospitalRoomCreate           = "web_hospital_room_create"
	PermissionWebHospitalRoomDelete           = "web_hospital_room_delete"
	PermissionWebHospitalRoomGet              = "web_hospital_room_get"
	PermissionWebHospitalRoomUpdate           = "web_hospital_room_update"
	PermissionWebHospitalRoomUpdateActivation = "web_hospital_room_update_activation"

	PermissionWebMasterBuildingCreate           = "web_master_building_create"
	PermissionWebMasterBuildingDelete           = "web_master_building_delete"
	PermissionWebMasterBuildingGet              = "web_master_building_get"
	PermissionWebMasterBuildingUpdate           = "web_master_building_update"
	PermissionWebMasterBuildingUpdateActivation = "web_master_building_update_activation"

	PermissionWebMasterLevelCreate           = "web_master_level_create"
	PermissionWebMasterLevelDelete           = "web_master_level_delete"
	PermissionWebMasterLevelGet              = "web_master_level_get"
	PermissionWebMasterLevelUpdate           = "web_master_level_update"
	PermissionWebMasterLevelUpdateActivation = "web_master_level_update_activation"

	PermissionWebMasterAreaCreate           = "web_master_area_create"
	PermissionWebMasterAreaDelete           = "web_master_area_delete"
	PermissionWebMasterAreaGet              = "web_master_area_get"
	PermissionWebMasterAreaUpdate           = "web_master_area_update"
	PermissionWebMasterAreaUpdateActivation = "web_master_area_update_activation"

	PermissionWebMasterRoomCreate           = "web_master_room_create"
	PermissionWebMasterRoomDelete           = "web_master_room_delete"
	PermissionWebMasterRoomGet              = "web_master_room_get"
	PermissionWebMasterRoomUpdate           = "web_master_room_update"
	PermissionWebMasterRoomUpdateActivation = "web_master_room_update_activation"

	PermissionWebMenuGizi                   = "web_menu_gizi"
	PermissionWebMenuKeamanan               = "web_menu_keamanan"
	PermissionWebMenuKebersihan             = "web_menu_kebersihan"
	PermissionWebMenuLaundry                = "web_menu_laundry"
	PermissionWebMenuMaintenanceMedis       = "web_menu_maintenance_medis"
	PermissionWebMenuMaintenanceUmum        = "web_menu_maintenance_umum"
	PermissionWebMenuOperatorTelepon        = "web_menu_operator_telepon"
	PermissionWebMenuTransport              = "web_menu_transport"
	PermissionWebMenuUmum                   = "web_menu_umum"
	PermissionWebMenuMasterRumahSakit       = "web_menu_master_rumah_sakit"
	PermissionWebMenuMasterUnitFasilitas    = "web_menu_master_unit_fasilitas"
	PermissionWebMenuMasterJenisLaporan     = "web_menu_master_jenis_laporan"
	PermissionWebMenuMasterKaryawan         = "web_menu_master_karyawan"
	PermissionWebMenuMasterShift            = "web_menu_master_shift"
	PermissionWebMenuReport                 = "web_menu_report"
	PermissionWebMenuFormManagement         = "web_menu_form_management"
	PermissionWebMenuUserManagement         = "web_menu_user_management"
	PermissionWebMenuAssignmentManagement   = "web_menu_assignment_management"
	PermissionWebMenuNotificationManagement = "web_menu_notification_management"
	PermissionWebMenuSetting                = "web_menu_setting"

	PermissionWebModuleGet = "web_module_get"

	PermissionWebNotificationTemplateCreate = "web_notification_template_create"
	PermissionWebNotificationTemplateDelete = "web_notification_template_delete"
	PermissionWebNotificationTemplateGet    = "web_notification_template_get"
	PermissionWebNotificationTemplateUpdate = "web_notification_template_update"

	PermissionWebPermissionGet = "web_permission_get"

	PermissionWebPositionCreate           = "web_position_create"
	PermissionWebPositionDelete           = "web_position_delete"
	PermissionWebPositionGet              = "web_position_get"
	PermissionWebPositionUpdate           = "web_position_update"
	PermissionWebPositionUpdateActivation = "web_position_update_activation"

	PermissionWebEmployeeCreate = "web_employee_create"
	PermissionWebEmployeeDelete = "web_employee_delete"
	PermissionWebEmployeeGet    = "web_employee_get"
	PermissionWebEmployeeImport = "web_employee_import"
	PermissionWebEmployeeUpdate = "web_employee_update"

	PermissionWebEmployeeScheduleCreate = "web_employee_schedule_create"
	PermissionWebEmployeeScheduleDelete = "web_employee_schedule_delete"
	PermissionWebEmployeeScheduleGet    = "web_employee_schedule_get"
	PermissionWebEmployeeScheduleImport = "web_employee_schedule_import"
	PermissionWebEmployeeScheduleUpdate = "web_employee_schedule_update"

	PermissionWebRoleCreate = "web_role_create"
	PermissionWebRoleUpdate = "web_role_update"
	PermissionWebRoleGet    = "web_role_get"
	PermissionWebRoleDelete = "web_role_delete"

	PermissionWebUserCreate           = "web_user_create"
	PermissionWebUserUpdate           = "web_user_update"
	PermissionWebUserUpdateActivation = "web_user_update_activation"
	PermissionWebUserUpdatePassword   = "web_user_update_password"
	PermissionWebUserGet              = "web_user_get"
	PermissionWebUserDelete           = "web_user_delete"

	PermissionWebWorkShiftCreate           = "web_work_shift_create"
	PermissionWebWorkShiftDelete           = "web_work_shift_delete"
	PermissionWebWorkShiftGet              = "web_work_shift_get"
	PermissionWebWorkShiftUpdate           = "web_work_shift_update"
	PermissionWebWorkShiftUpdateActivation = "web_work_shift_update_activation"
)
