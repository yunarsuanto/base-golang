package objects

type ListUserRequest struct {
	IsReporter    *bool
	IsVerificator *bool
	IsActive      *bool
}

type DetailUserRequest struct {
	Id string
}

type ListUserHospital struct {
	Id   string
	Name string
}

type ListUserRole struct {
	Id            string
	Name          string
	IsReporter    bool
	IsVerificator bool
}

type GetUser struct {
	Id                 string
	Name               string
	Email              string
	Username           string
	Password           string
	IdNumber           *string
	PhoneNumber        *string
	ProfilePicturePath *string
	ProfilePictureUrl  string
	IsActive           bool
	EmployeeId         *string
	EmployeeNumber     *string
	Hospitals          []ListUserHospital
	Roles              []ListUserRole
}

type CreateUser struct {
	Name               string
	Email              string
	Username           string
	Password           string
	IdNumber           string
	PhoneNumber        string
	ProfilePicturePath string
	IsActive           bool
	HospitalIds        []string
	RoleIds            []string
}

type UpdateUser struct {
	Id                 string
	Name               string
	Email              string
	Username           string
	IdNumber           string
	PhoneNumber        string
	ProfilePicturePath string
	IsActive           bool
	HospitalIds        []string
	RoleIds            []string
}

type UpdateProfileUser struct {
	Username           string
	PhoneNumber        string
	ProfilePicturePath string
}

type DeleteUser struct {
	Id string
}

type UpdatePasswordUser struct {
	Id       string
	Password string
}

type UpdateActivationUser struct {
	Id       string
	IsActive bool
}
