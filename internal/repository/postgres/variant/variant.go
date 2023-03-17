package variant_repo

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
	variant_srvc "github.com/Abdullayev65/online_test/internal/service/variant"
	"github.com/uptrace/bun"
	"time"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r *Repository) CreateVariant(c context.Context, data *variant_srvc.Create, userID int) (*entity.Variant, error) {
	m := new(entity.Variant)
	if data.Name != nil {
		m.Name = *data.Name
	}
	m.CreatedBy = userID
	_, err := r.DB.NewInsert().Model(m).Exec(c)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *Repository) VariantByID(c context.Context, id int) (*entity.Variant, error) {
	m := new(entity.Variant)
	m.ID = id
	err := r.DB.NewSelect().Model(m).WherePK().Scan(c)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *Repository) ListVariant(c context.Context, size, page int) ([]entity.Variant, error) {
	ms := make([]entity.Variant, 0)
	err := r.DB.NewSelect().Model(&ms).Limit(size).
		Offset((size - 1) * page).Order("id").Scan(c)
	return ms, err
}

func (r *Repository) DeleteVariant(c context.Context, id, userID int) error {
	m, err := r.VariantByID(c, id)
	if err != nil {
		return err
	}
	m.DeletedAt = time.Now()
	m.DeletedBy = userID
	_, err = r.DB.NewUpdate().Model(m).WherePK().Exec(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Repository_() {
	println("just for inherits")
}
