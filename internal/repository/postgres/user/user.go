package user

import (
	"context"
	"errors"
	"github.com/Abdullayev65/online_test/internal/entity"
	user_service "github.com/Abdullayev65/online_test/internal/service/user"
	"github.com/uptrace/bun"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r *Repository) Repository_() {

}
func (r *Repository) UserByUsername(c context.Context, username string) (*entity.User, error) {
	user := new(entity.User)
	err := r.DB.NewSelect().Model(user).
		Where("username = ?", username).Scan(c)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *Repository) UserByID(c context.Context, userID int) (*entity.User, error) {
	user := &entity.User{ID: userID}
	err := r.DB.NewSelect().Model(user).WherePK().Scan(c)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *Repository) UserUpdate(c context.Context, userID int, update *user_service.Update) error {
	user, err := r.UserByID(c, userID)
	if err != nil {
		return errors.New("user not found")
	}
	if update.ID != nil && *update.ID != userID {
		if user.Type != 1 {
			return errors.New("no permission")
		}
		user, err = r.UserByID(c, *update.ID)
		if err != nil {
			return errors.New("user not found")
		}
	}
	if update.Username != nil {
		user.Username = *update.Username
	}
	if update.Password != nil {
		user.Password = *update.Password
	}
	_, err = r.DB.NewUpdate().Model(user).WherePK().Exec(c)
	return err
}
func (r *Repository) AddUser(c context.Context, user *entity.User) error {
	_, err := r.DB.NewInsert().Model(user).Exec(c)
	if err != nil {
		return err
	}
	return nil
}
