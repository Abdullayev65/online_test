package topic_repo

import (
	"context"
	"errors"
	"github.com/Abdullayev65/online_test/internal/entity"
	topic "github.com/Abdullayev65/online_test/internal/service/topic"
	"github.com/uptrace/bun"
	"time"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r Repository) GetAll(c context.Context, filter *topic.Filter) ([]entity.Topic, int, error) {
	var list []entity.Topic

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

	if filter.CreatedBy != nil {
		q.WhereGroup(" and ", func(query *bun.SelectQuery) *bun.SelectQuery {
			query.Where("created_by = ?", *filter.CreatedBy)
			return query
		})
	}

	if filter.AllWithDeleted {
		q.WhereAllWithDeleted()
	} else if filter.OnlyDeleted {
		q.WhereDeleted()
	}

	count, err := q.ScanAndCount(c)

	return list, count, err
}
func (r *Repository) GetByID(c context.Context, id int) (*entity.Topic, error) {
	ent := new(entity.Topic)
	ent.ID = id
	err := r.DB.NewSelect().Model(ent).WherePK().Scan(c)
	if err != nil {
		return nil, err
	}
	return ent, nil
}
func (r *Repository) Create(c context.Context, ent *entity.Topic) error {
	_, err := r.DB.NewInsert().Model(ent).Exec(c)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) Update(c context.Context, id int, update *topic.Update) error {
	ent, err := r.GetByID(c, id)
	if err != nil {
		return errors.New("topic not found")
	}

	ent.Name = update.Name
	_, err = r.DB.NewUpdate().Model(ent).WherePK().Exec(c)
	return err
}
func (r *Repository) Delete(c context.Context, id int, userID int) error {
	topic, err := r.GetByID(c, id)
	if err != nil {
		return err
	}
	topic.DeletedBy = &userID
	topic.DeletedAt = time.Now()
	_, err = r.DB.NewUpdate().Model(topic).WherePK().Exec(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Repository_() {
}
