package variant_repo

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
	"github.com/Abdullayev65/online_test/internal/service/variant"
	"github.com/uptrace/bun"
	"time"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

// impl

func (r *Repository) GetAll(c context.Context, filter *variant_srvc.Filter) ([]entity.Variant, int, error) {
	var list []entity.Variant
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

func (r *Repository) GetByID(c context.Context, id int) (*entity.Variant, error) {
	m := new(entity.Variant)
	m.ID = id
	err := r.DB.NewSelect().Model(m).WherePK().Scan(c)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *Repository) Create(c context.Context, data *variant_srvc.Create) (*entity.Variant, error) {

	m := new(entity.Variant)
	m.Name = data.Name

	m.CreatedBy = data.UserID
	_, err := r.DB.NewInsert().Model(m).Exec(c)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *Repository) Delete(c context.Context, id, userID int) error {
	m, err := r.GetByID(c, id)
	if err != nil {
		return err
	}
	m.DeletedAt = time.Now()
	m.DeletedBy = &userID
	_, err = r.DB.NewUpdate().Model(m).WherePK().Exec(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Repository_() {
	println("just for inherits")
}
