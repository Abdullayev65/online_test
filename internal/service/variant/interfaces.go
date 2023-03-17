package variant_srvc

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
)

type Repository interface {
	CreateVariant(c context.Context, data *Create, userID int) (*entity.Variant, error)
	VariantByID(c context.Context, id int) (*entity.Variant, error)
	ListVariant(c context.Context, size, page int) ([]entity.Variant, error)
	DeleteVariant(c context.Context, id, userID int) error
	Repository_()
}
