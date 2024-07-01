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

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
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

func getHashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), err
}

func checkPassword(hashPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))

	if err != nil {
		return true, errors.New("Password Is Not Match")
	}

	return false, nil
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanifIfError(err)

	tx, err := service.DB.Begin()
	helper.PanifIfError(err)
	defer helper.CommitOrRollback(tx)

	hashPassword, err := getHashPassword(request.Password)
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

	errCheckPass, err := checkPassword(user.Password, request.Password)
	if errCheckPass {
		panic(exception.NewBadRequestError(err.Error()))
	}

	return helper.ToUserLoginResponse(user)
}
