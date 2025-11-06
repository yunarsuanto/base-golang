
		package objects

		type CreatePermissionRequest struct {
			Name string
		}

		type UpdatePermissionRequest struct {
			Id       string
			Name string
		}

		type DeletePermissionRequest struct {
			Id string
		}

		type ListPermissionResponse struct {
			Id       string
			Name string
		}

	