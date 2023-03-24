package user

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
	"github.com/Abdullayev65/online_test/internal/service/user"
)

type User interface {
	GetByUsername(c context.Context, username string) (*entity.User, error)
	GetByID(c context.Context, userID int) (*entity.User, error)
	Update(c context.Context, userID int, user *user.Update) error
	Create(c context.Context, user *entity.User) error
	Service_()
}
