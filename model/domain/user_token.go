package domain

type UserToken struct {
	Id      int
	UserId  int
	RefreshToken   string
	IsValid int
}
