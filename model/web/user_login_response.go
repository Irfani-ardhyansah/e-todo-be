package web

type UserLoginResponse struct {
	Id           int    `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
