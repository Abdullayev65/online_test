package question_srvc

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
)

type Repository interface {
	GetAll(c context.Context, filter *Filter) ([]entity.Question, int, error)
	GetByID(c context.Context, id int) (*entity.Question, error)
	Create(c context.Context, data *Create) (*entity.Question, error)
	Update(c context.Context, data *Update) error
	Delete(c context.Context, id, userID int) error
	Repository_()
}
