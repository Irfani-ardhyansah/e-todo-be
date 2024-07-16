package service

import (
	"context"
	"database/sql"
	exception "e-todo/excception"
	"e-todo/helper"
	"e-todo/model/domain"
	"e-todo/model/web"
	"e-todo/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanifIfError(err)

	tx, err := service.DB.Begin()
	helper.PanifIfError(err)
	defer helper.CommitOrRollback(tx)

	hashPassword, err := helper.GetHashPassword(request.Password)
	helper.PanifIfError(err)

	user := domain.User{
		Email:    request.Email,
		Password: hashPassword,
	}

	user = service.UserRepository.Save(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Login(ctx context.Context, request web.UserCreateRequest) web.UserLoginResponse {
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
	jwtAccessToken := helper.GenereateJwtToken(expAccessTime, user.Id, user.Email)
	user.AccessToken = jwtAccessToken

	expResfreshTime := time.Now().Add(time.Hour * 2)
	jwtRefreshToken := helper.GenereateJwtToken(expResfreshTime, user.Id, user.Email)
	user.RefreshToken = jwtRefreshToken

	userToken := domain.UserToken{
		UserId:       user.Id,
		RefreshToken: user.RefreshToken,
		IsValid:      1,
	}

	tx, err := service.DB.Begin()
	helper.PanifIfError(err)
	defer helper.CommitOrRollback(tx)

	errorUserToken := service.UserRepository.SaveToken(ctx, tx, userToken)
	if errorUserToken != nil {
		helper.PanifIfError(errorUserToken)
	}

	return helper.ToUserLoginResponse(user)
}
