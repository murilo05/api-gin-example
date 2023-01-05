package interfaces

import (
	"context"

	"gitlab.engdb.com.br/apigin/domain/entities"
)

type UserUseCase interface {
	GetUser(ctx context.Context) ([]entities.User, int, *entities.Error)
	GetUserById(ctx context.Context, userId string) (*entities.User, int, *entities.Error)
	CreateUser(ctx context.Context, body entities.User) (int, *entities.Error)
	EditUser(ctx context.Context, body entities.User, userId string) (int, *entities.Error)
	DeleteUser(ctx context.Context, userId string) (int, *entities.Error)
}
