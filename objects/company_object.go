package objects

type CreateCompanyRequest struct {
	Code            string
	Name            string
	Image           string
	Email           string
	Address         string
	PostalCode      string
	NationalId      string
	NationalIdImage string
	Npwp            string
	NpwpImage       string
	DirectorName    string
	Longitude       float32
	Latitude        float32
	IsActive        bool
}

type UpdateCompanyRequest struct {
	Id              string
	Code            string
	Name            string
	Image           string
	Email           string
	Address         string
	PostalCode      string
	NationalId      string
	NationalIdImage string
	Npwp            string
	NpwpImage       string
	DirectorName    string
	Longitude       float32
	Latitude        float32
	IsActive        bool
}

type DeleteCompanyRequest struct {
	Id string
}

type ListCompanyResponse struct {
	Id       string
	Code     string
	Name     string
	Image    string
	Email    string
	IsActive bool
}

type DetailCompanyRequest struct {
	Id string
}

type DetailCompanyResponse struct {
	Id              string
	Code            string
	Name            string
	Image           string
	Email           string
	Address         string
	PostalCode      string
	NationalId      string
	NationalIdImage string
	Npwp            string
	NpwpImage       string
	DirectorName    string
	Longitude       float32
	Latitude        float32
	IsActive        bool
}
