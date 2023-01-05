package interfaces

import (
	"context"

	"gitlab.engdb.com.br/apigin/domain/entities"
)

type UserRepo interface {
	GetUser(ctx context.Context) ([]entities.User, int, error)
	GetUserById(ctx context.Context, userId string) (*entities.User, int, error)
	CreateUser(ctx context.Context, body entities.User) (int, error)
	EditUser(ctx context.Context, body entities.User, userId string) (int, error)
	DeleteUser(ctx context.Context, userId string) (int, error)
}
