package web

type RefreshTokenRequest struct {
	UserId       string `validate: "required" json:"userId"`
	RefreshToken string `validate: "required" json:"refreshToken"`
}
