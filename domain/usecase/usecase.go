package usecase

import (
	"context"
	"errors"

	"gitlab.engdb.com.br/apigin/domain/entities"
	"gitlab.engdb.com.br/apigin/interfaces"
	errorUtils "gitlab.engdb.com.br/apigin/utils/error"
)

type userUsecase struct {
	userRepo interfaces.UserRepo
}

func NewUserUsecase(us interfaces.UserRepo) interfaces.UserUseCase {
	return &userUsecase{
		userRepo: us,
	}
}

//GetUser - Returns all users from table users
func (u *userUsecase) GetUser(ctx context.Context) ([]entities.User, int, *entities.Error) {

	users, code, err := u.userRepo.GetUser(ctx)
	if err != nil {
		errResp := errorUtils.CreateError(code, err.Error())
		return nil, code, &errResp
	}

	return users, code, nil
}

//GetUserById - Returns an user by id
func (u *userUsecase) GetUserById(ctx context.Context, userId string) (*entities.User, int, *entities.Error) {

	user, code, err := u.userRepo.GetUserById(ctx, userId)
	if err != nil {
		errResp := errorUtils.CreateError(code, err.Error())
		return nil, code, &errResp
	}
	if user != nil {
		return user, code, nil
	} else {
		code = 404
		err = errors.New("user not found")
		errResp := errorUtils.CreateError(code, err.Error())
		return nil, code, &errResp
	}
}

func (u *userUsecase) CreateUser(ctx context.Context, body entities.User) (int, *entities.Error) {

	code, err := u.userRepo.CreateUser(ctx, body)
	if err != nil {
		errResp := errorUtils.CreateError(code, err.Error())
		return code, &errResp
	}

	return 200, nil
}

func (u *userUsecase) DeleteUser(ctx context.Context, userId string) (int, *entities.Error) {

	//Checking if user exists
	user, code, err := u.userRepo.GetUserById(ctx, userId)
	if err != nil {
		errResp := errorUtils.CreateError(code, err.Error())
		return code, &errResp
	}

	if user != nil {
		code, err := u.userRepo.DeleteUser(ctx, userId)
		if err != nil {
			errResp := errorUtils.CreateError(code, err.Error())
			return code, &errResp
		}

		return code, nil

	} else {
		code = 404
		err = errors.New("user not found")
		errResp := errorUtils.CreateError(code, err.Error())
		return code, &errResp
	}
}

func (u *userUsecase) EditUser(ctx context.Context, body entities.User, userId string) (int, *entities.Error) {

	//Checking if user exists
	user, code, err := u.userRepo.GetUserById(ctx, userId)
	if err != nil {
		errResp := errorUtils.CreateError(code, err.Error())
		return code, &errResp
	}

	if user != nil {
		code, err := u.userRepo.EditUser(ctx, body, userId)
		if err != nil {
			errResp := errorUtils.CreateError(code, err.Error())
			return code, &errResp
		}

		return code, nil

	} else {
		code = 404
		err = errors.New("user not found")
		errResp := errorUtils.CreateError(code, err.Error())
		return code, &errResp
	}
}
