package question_repo

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
	"github.com/Abdullayev65/online_test/internal/service/question"
	"github.com/uptrace/bun"
	"time"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r *Repository) GetAll(c context.Context, filter *question_srvc.Filter) ([]entity.Question, int, error) {
	var list []entity.Question
	q := r.NewSelect().Model(&list)

	if filter.Limit != nil {
		q.Limit(*filter.Limit)
	}

	if filter.Offset != nil {
		q.Offset(*filter.Offset)
	}

	if filter.TopicID != nil {
		q.WhereGroup(" and ", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Where("topic_id = ?", *filter.TopicID)
		})
	}

	if filter.IDs != nil {
		q.WhereGroup(" and ", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Where("id in (?)", bun.In(&filter.IDs))
		})
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

func (r *Repository) GetByID(c context.Context, id int) (*entity.Question, error) {
	m := new(entity.Question)
	m.ID = id
	err := r.DB.NewSelect().Model(m).WherePK().Scan(c)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *Repository) Create(c context.Context,
	data *question_srvc.Create) (*entity.Question, error) {

	m := new(entity.Question)
	m.ImagePath = &data.ImagePath
	m.Text = data.Text
	m.Description = data.Description
	m.TopicID = data.TopicID
	m.CreatedBy = data.UserId

	_, err := r.DB.NewInsert().Model(m).Exec(c)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (r *Repository) Update(c context.Context, data *question_srvc.Update) error {

	m := new(entity.Question)
	m.ID = data.ID

	m.Text = data.Text
	m.Description = data.Description
	m.TopicID = data.TopicID

	_, err := r.DB.NewUpdate().Model(m).WherePK().Exec(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, id, userID int) error {
	_, err := r.NewUpdate().Table("questions").Set(
		"deleted_at = ?, deleted_by = ?",
		time.Now(),
		userID,
	).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Repository_() {
	println("just for inherits")
}
