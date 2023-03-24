package user

import (
	"context"
	"errors"
	"github.com/Abdullayev65/online_test/internal/entity"
	"github.com/Abdullayev65/online_test/internal/service/user"
	"github.com/uptrace/bun"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r *Repository) GetAll(c context.Context, filter user.Filter) ([]entity.User, int, error) {
	var list []entity.User
	q := r.NewSelect().Model(&list)

	if filter.Limit != nil {
		q.Limit(*filter.Limit)
	}

	if filter.Offset != nil {
		q.Offset(*filter.Offset)
	}

	if filter.Order != nil {
		q.Order(*filter.Order)
	} else {
		q.Order("id desc")
	}

	if filter.AllWithDeleted {
		q.WhereAllWithDeleted()
	} else if filter.OnlyDeleted {
		q.WhereDeleted()
	}

	count, err := q.ScanAndCount(c)

	return list, count, err
}

func (r *Repository) GetByID(c context.Context, userID int) (*entity.User, error) {
	user := &entity.User{ID: userID}
	err := r.DB.NewSelect().Model(user).WherePK().Scan(c)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) Create(c context.Context, user *entity.User) error {
	_, err := r.DB.NewInsert().Model(user).Exec(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Update(c context.Context, userID int, update *user.Update) error {
	user, err := r.GetByID(c, userID)
	if err != nil {
		return errors.New("user not found")
	}
	if update.ID != nil && *update.ID != userID {
		if *user.Type != 1 {
			return errors.New("no permission")
		}
		user, err = r.GetByID(c, *update.ID)
		if err != nil {
			return errors.New("user not found")
		}
	}
	user.Username = update.Username
	user.Password = update.Password
	_, err = r.DB.NewUpdate().Model(user).WherePK().Exec(c)
	return err
}

// specific functions
func (r *Repository) GetByUsername(c context.Context, username string) (*entity.User, error) {
	user := new(entity.User)
	err := r.DB.NewSelect().Model(user).
		Where("username = ?", username).Scan(c)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) Repository_() {

}
