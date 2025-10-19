package objects

type ListUserTokenRequest struct {
	UserIds     []string
	Platform    string
	GetFcmToken *bool
}
