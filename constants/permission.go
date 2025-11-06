package constants

const (
	PermissionCategoryAccess  = "Akses"
	PermissionCategoryWebUser = "Web App: Pengguna"
)

// type permissionCategory struct {
// 	Name            string
// 	PermissionCodes []string
// }

// func PermissionCategories() []permissionCategory {
// 	return []permissionCategory{
// 		{
// 			Name: PermissionCategoryAccess,
// 			PermissionCodes: []string{
// 				PermissionAccessMobile,
// 				PermissionAccessWeb,
// 			},
// 		},
// 		{
// 			Name: PermissionCategoryWebUser,
// 			PermissionCodes: []string{
// 				PermissionWebUserCreate,
// 				PermissionWebUserList,
// 			},
// 		},
// 	}
// }

const (
	PermissionAccessMobile = "access_mobile"
	PermissionAccessWeb    = "access_web"

	PermissionUserList   = "user_list"
	PermissionUserDetail = "user_detail"
	PermissionUserCreate = "user_create"
	PermissionUserUpdate = "user_update"
	PermissionUserDelete = "user_delete"

	PermissionRoleList   = "role_list"
	PermissionRoleDetail = "role_detail"
	PermissionRoleCreate = "role_create"
	PermissionRoleUpdate = "role_update"
	PermissionRoleDelete = "role_delete"

	PermissionPermissionList   = "permission_list"
	PermissionPermissionCreate = "permission_create"
	PermissionPermissionUpdate = "permission_update"
	PermissionPermissionDelete = "permission_delete"

	PermissionRolePermissionUpsert = "role_permission_upsert"
	PermissionRolePermissionDelete = "role_permission_delete"
	PermissionUserRoleUpsert       = "user_role_upsert"
	PermissionUserRoleDelete       = "user_role_delete"

	PermissionCategoryLessonList   = "category_lesson_list"
	PermissionCategoryLessonDetail = "category_lesson_detail"
	PermissionCategoryLessonCreate = "category_lesson_create"
	PermissionCategoryLessonUpdate = "category_lesson_update"
	PermissionCategoryLessonDelete = "category_lesson_delete"

	PermissionLessonList   = "lesson_list"
	PermissionLessonDetail = "lesson_detail"
	PermissionLessonCreate = "lesson_create"
	PermissionLessonUpdate = "lesson_update"
	PermissionLessonDelete = "lesson_delete"

	PermissionLessonItemList   = "lesson_item_list"
	PermissionLessonItemDetail = "lesson_item_detail"
	PermissionLessonItemCreate = "lesson_item_create"
	PermissionLessonItemUpdate = "lesson_item_update"
	PermissionLessonItemDelete = "lesson_item_delete"
)
