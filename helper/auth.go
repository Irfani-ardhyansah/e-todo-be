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

func GenereateJwtToken(expTime time.Time, id int, email string, typeToken string) string {
	claims := &config.JWTClaim{
		Id:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	key, err := whetherTypeToken(typeToken)
	PanifIfError(err)

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlgo.SignedString(key)
	PanifIfError(err)

	return token
}

func VerifyToken(tokenString string, typeToken string) (jwt.MapClaims, error) {
	key, err := whetherTypeToken(typeToken)
	PanifIfError(err)

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Token Is Not Valid")
	}

	return claims, nil
}

func whetherTypeToken(typeToken string) ([]byte, error) {
	var key []byte
	if typeToken == "access" {
		key = config.ACCESS_KEY
	} else if typeToken == "refresh" {
		key = config.REFRESH_KEY
	}

	return key, nil
}

func UserClaims(request *http.Request) jwt.MapClaims {
	claims, ok := request.Context().Value("jwtClaims").(jwt.MapClaims)
	if !ok {
		errors.New("Missing claims in context")
	}

	return claims
}

func ExtractID(claims map[string]interface{}) (int, error) {
	if idValue, ok := claims["id"]; ok {
		switch v := idValue.(type) {
		case float64:
			return int(v), nil
		case int:
			return v, nil
		case string:
			return 0, fmt.Errorf("id is a string, conversion not implemented")
		default:
			return 0, fmt.Errorf("unexpected type for id: %T", idValue)
		}
	}

	return 0, fmt.Errorf("id not found in claims")
}
