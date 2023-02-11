package handler

import (
	"context"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"gitlab.engdb.com.br/apigin/domain/entities"
	"gitlab.engdb.com.br/apigin/interfaces"
)

func NewUserUseCase(user []entities.User, codex int, err *entities.Error) interfaces.UserUseCase {
	return &userUseCaseMock{
		user:  user,
		codex: codex,
		err:   err,
	}
}

type userUseCaseMock struct {
	user  []entities.User
	codex int
	err   *entities.Error
}

func (us *userUseCaseMock) GetUser(ctx context.Context) (user []entities.User, codex int, err *entities.Error) {
	user = us.user
	codex = us.codex
	err = us.err
	return
}

func (us *userUseCaseMock) GetUserById(ctx context.Context, userId string) (user *entities.User, codex int, err *entities.Error) {
	return
}

func (us *userUseCaseMock) CreateUser(ctx context.Context, body entities.User) (codex int, err *entities.Error) {
	return
}

func (us *userUseCaseMock) EditUser(ctx context.Context, body entities.User, userId string) (codex int, err *entities.Error) {
	return
}

func (us *userUseCaseMock) DeleteUser(ctx context.Context, userId string) (codex int, err *entities.Error) {
	return
}

func TestNewProjectHandler(t *testing.T) {
	type args struct {
		r  *gin.Engine
		us interfaces.UserUseCase
	}
	tests := []struct {
		name string
		args args
		want *gin.Context
	}{
		{
			name: "sucesso",
			args: args{r: gin.Default(), us: NewUserUseCase([]entities.User{}, 0, &entities.Error{})},
			want: &gin.Context{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProjectHandler(tt.args.r, tt.args.us); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProjectHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserHandler_getUsers(t *testing.T) {
	type fields struct {
		userUseCase interfaces.UserUseCase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uh := &UserHandler{
				userUseCase: tt.fields.userUseCase,
			}
			uh.getUsers(tt.args.c)
		})
	}
}

func TestUserHandler_getUserById(t *testing.T) {
	type fields struct {
		userUseCase interfaces.UserUseCase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uh := &UserHandler{
				userUseCase: tt.fields.userUseCase,
			}
			uh.getUserById(tt.args.c)
		})
	}
}

func TestUserHandler_createUser(t *testing.T) {
	type fields struct {
		userUseCase interfaces.UserUseCase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uh := &UserHandler{
				userUseCase: tt.fields.userUseCase,
			}
			uh.createUser(tt.args.c)
		})
	}
}

func TestUserHandler_deleteUser(t *testing.T) {
	type fields struct {
		userUseCase interfaces.UserUseCase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uh := &UserHandler{
				userUseCase: tt.fields.userUseCase,
			}
			uh.deleteUser(tt.args.c)
		})
	}
}

func TestUserHandler_editUser(t *testing.T) {
	type fields struct {
		userUseCase interfaces.UserUseCase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uh := &UserHandler{
				userUseCase: tt.fields.userUseCase,
			}
			uh.editUser(tt.args.c)
		})
	}
}
