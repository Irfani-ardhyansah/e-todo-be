package helper

import (
	"e-todo/config"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GetHashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), err
}

func CheckPassword(hashPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))

	if err != nil {
		return true, errors.New("Password Is Not Match")
	}

	return false, nil
}

func GenereateJwtToken(expTime time.Time, id int, email string) string {
	claims := &config.JWTClaim{
		Id:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlgo.SignedString(config.ACCESS_KEY)
	PanifIfError(err)

	return token
}

func VerifyToken(tokenString string, typeToken string) (jwt.MapClaims, error) {

	var key string
	if typeToken == "access" {
		key = string(config.ACCESS_KEY)
	} else if typeToken == "refresh" {
		key = string(config.REFRESH_KEY)
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func UserClaims(request *http.Request) jwt.MapClaims {
	claims, ok := request.Context().Value("jwtClaims").(jwt.MapClaims)
	if !ok {
		fmt.Println("Missing claims in context")
	}

	return claims
}
