package config

import "github.com/golang-jwt/jwt/v5"

var REFRESH_KEY = []byte("klkasdjqpwkacnsi1293724029")

var ACCESS_KEY = []byte("ashdjqy9283409bsdk1kg8hda01")

type JWTClaim struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
