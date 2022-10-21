package usecase

import (
	"context"
	"errors"
	"fmt"

	"gitlab.engdb.com.br/apigin/domain/entities"
	"gitlab.engdb.com.br/apigin/infrastructure/repository"
	errorUtils "gitlab.engdb.com.br/apigin/utils/error"
)

//GetUser - Returns all users from table users
func GetUser(ctx context.Context) ([]entities.User, int, *entities.Error) {

	//Just a simulation to get values from Header, it's not necessary to have a specific token.
	token := ctx.Value("token")
	if token == "" {
		err := errors.New("Bad Request")
		fmt.Println("Token error")
		errResp := errorUtils.CreateError(400, err.Error())
		return nil, 400, &errResp
	}

	users, code, err := repository.GetUser(ctx)
	if err != nil {
		errResp := errorUtils.CreateError(code, err.Error())
		return nil, code, &errResp
	}

	return users, code, nil
}

//GetUserById - Returns an user by id
func GetUserById(ctx context.Context, userId string) (*entities.User, int, *entities.Error) {

	//Just a simulation to get values from Header, it's not necessary to have a specific token.
	token := ctx.Value("token")
	if token == "" {
		err := errors.New("Bad Request")
		fmt.Println("Token error")
		errResp := errorUtils.CreateError(400, err.Error())
		return nil, 400, &errResp
	}

	user, code, err := repository.GetUserById(ctx, userId)
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

func CreateUser(ctx context.Context, body entities.User) (int, *entities.Error) {

	//Just a simulation to get values from Header, it's not necessary to have a specific token.
	token := ctx.Value("token")
	if token == "" {
		err := errors.New("Bad Request")
		fmt.Println("Token error")
		errResp := errorUtils.CreateError(400, err.Error())
		return 400, &errResp
	}

	code, err := repository.CreateUser(ctx, body)
	if err != nil {
		errResp := errorUtils.CreateError(code, err.Error())
		return code, &errResp
	}

	return 200, nil
}

func DeleteUser(ctx context.Context, userId string) (int, *entities.Error) {

	//Just a simulation to get values from Header, it's not necessary to have a specific token.
	token := ctx.Value("token")
	if token == "" {
		err := errors.New("Bad Request")
		fmt.Println("Token error")
		errResp := errorUtils.CreateError(400, err.Error())
		return 400, &errResp
	}

	//Checking if user exists
	user, code, err := repository.GetUserById(ctx, userId)
	if err != nil {
		errResp := errorUtils.CreateError(code, err.Error())
		return code, &errResp
	}

	if user != nil {
		code, err := repository.DeleteUser(ctx, userId)
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

func EditUser(ctx context.Context, body entities.User, userId string) (int, *entities.Error) {

	//Just a simulation to get values from Header, it's not necessary to have a specific token.
	token := ctx.Value("token")
	if token == "" {
		err := errors.New("Bad Request")
		fmt.Println("Token error")
		errResp := errorUtils.CreateError(400, err.Error())
		return 400, &errResp
	}

	//Checking if user exists
	user, code, err := repository.GetUserById(ctx, userId)
	if err != nil {
		errResp := errorUtils.CreateError(code, err.Error())
		return code, &errResp
	}

	if user != nil {
		code, err := repository.EditUser(ctx, body, userId)
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
