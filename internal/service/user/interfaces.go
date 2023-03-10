package user

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
)

type Repository interface {
	UserByUsername(c context.Context, username string) (*entity.User, error)
	UserByID(c context.Context, userID int) (*entity.User, error)
	UserUpdate(c context.Context, userID int, user *Update) error
	AddUser(c context.Context, user *entity.User) error
	Repository_()
}
