package variant_question_srvc

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
)

type Repository interface {
	GetAll(c context.Context, filter *Filter) ([]entity.VariantQuestion, int, error)
	GetById(c context.Context, id int) (*entity.VariantQuestion, error)
	Create(c context.Context, data *Create) ([]entity.VariantQuestion, error)
	Delete(c context.Context, id int) error
	Repository_()
}
