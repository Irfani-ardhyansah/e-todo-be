package service

import (
	"context"
	"database/sql"
	exception "e-todo/excception"
	"e-todo/helper"
	"e-todo/model/domain"
	"e-todo/model/web"
	"e-todo/repository"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewAuthService(authRepository repository.AuthRepository, userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *AuthServiceImpl) Login(ctx context.Context, request web.UserCreateRequest) web.UserLoginResponse {
	err := service.Validate.Struct(request)
	helper.PanifIfError(err)

	user, err := service.UserRepository.FindByEmail(ctx, service.DB, request.Email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	errCheckPass, err := helper.CheckPassword(user.Password, request.Password)
	if errCheckPass {
		panic(exception.NewBadRequestError(err.Error()))
	}

	expAccessTime := time.Now().Add(time.Minute * 30)
	jwtAccessToken := helper.GenereateJwtToken(expAccessTime, user.Id, user.Email, "access")
	user.AccessToken = jwtAccessToken

	expRefreshTime := time.Now().Add(time.Hour * 2)
	jwtRefreshToken := helper.GenereateJwtToken(expRefreshTime, user.Id, user.Email, "refresh")
	user.RefreshToken = jwtRefreshToken

	userToken := domain.UserToken{
		UserId:       user.Id,
		RefreshToken: user.RefreshToken,
		IsValid:      1,
	}

	tx, err := service.DB.Begin()
	helper.PanifIfError(err)
	defer helper.CommitOrRollback(tx)

	errorUserToken := service.AuthRepository.SaveToken(ctx, tx, userToken)
	if errorUserToken != nil {
		helper.PanifIfError(err)
	}

	return helper.ToUserLoginResponse(user)
}

func (service *AuthServiceImpl) RefreshToken(ctx context.Context, userId int, refreshToken string) web.UserLoginResponse {
	validRefreshToken := service.AuthRepository.CheckValidToken(ctx, service.DB, userId, refreshToken)
	fmt.Println(validRefreshToken)
	if !validRefreshToken {
		panic(errors.New("Token Data Is NOt Valid From DB"))
	}

	_, err := helper.VerifyToken(refreshToken, "refresh")
	if err != nil {
		helper.PanifIfError(err)
	}

	user, err := service.UserRepository.FindById(ctx, service.DB, userId)

	expAccessTime := time.Now().Add(time.Minute * 30)
	jwtAccessToken := helper.GenereateJwtToken(expAccessTime, user.Id, user.Email, "access")
	user.AccessToken = jwtAccessToken

	user.RefreshToken = refreshToken

	return helper.ToUserLoginResponse(user)
}
