package user

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
)

type Repository interface {
	GetByUsername(c context.Context, username string) (*entity.User, error)
	GetByID(c context.Context, userID int) (*entity.User, error)
	Update(c context.Context, userID int, user *Update) error
	Create(c context.Context, user *entity.User) error
	Repository_()
}
