package topic_uc

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
)

type User interface {
	UserByUsername(c context.Context, username string) (*entity.User, error)
	UserByID(c context.Context, userID int) (*entity.User, error)
	AddUser(c context.Context, user *entity.User) error
	UserUpdate(c context.Context, user *entity.User) error
}
