package variant_srvc

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
)

type Repository interface {
	GetAll(c context.Context, filter *Filter) ([]entity.Variant, int, error)
	GetByID(c context.Context, id int) (*entity.Variant, error)
	Create(c context.Context, data *Create) (*entity.Variant, error)
	Delete(c context.Context, id, userID int) error
	Repository_()
}
