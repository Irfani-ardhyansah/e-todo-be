package web

type UserLoginResponse struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Token string `json: "token"`
}
