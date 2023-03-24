package variant_question_repo

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
	variant_question "github.com/Abdullayev65/online_test/internal/service/variant_question"
	"github.com/uptrace/bun"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r *Repository) GetAll(c context.Context, filter *variant_question.Filter) ([]entity.VariantQuestion, int, error) {
	var list []entity.VariantQuestion
	q := r.NewSelect().Model(&list)

	if filter.VariantID != nil {
		q.WhereGroup(" and ", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Where(" variant_id = ? ", *filter.VariantID)
		})
	}

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

func (r *Repository) GetById(c context.Context, id int) (*entity.VariantQuestion, error) {
	m := new(entity.VariantQuestion)
	m.ID = id
	err := r.DB.NewSelect().Model(m).WherePK().Scan(c)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *Repository) Create(c context.Context, data *variant_question.Create) ([]entity.VariantQuestion, error) {

	qs := make([]entity.Question, 0, data.Count)
	err := r.DB.NewSelect().Model(&qs).Column("id").
		Where("topic_id = ?", data.TopicID).
		OrderExpr("random()").Limit(data.Count).Scan(c)
	if err != nil {
		return nil, err
	}

	ms := make([]entity.VariantQuestion, 0, data.Count)
	for _, q := range qs {
		ms = append(ms, entity.VariantQuestion{QuestionID: &q.ID, VariantID: &data.VariantID})
	}

	_, err = r.DB.NewInsert().Model(&ms).Exec(c)
	if err != nil {
		return nil, err
	}

	return ms, nil
}

func (r *Repository) Delete(c context.Context, id int) error {
	m, err := r.GetById(c, id)
	if err != nil {
		return err
	}
	_, err = r.DB.NewDelete().Model(m).WherePK().Exec(c)
	return err
}

func (r *Repository) Repository_() {
	println("just for inherits")
}
