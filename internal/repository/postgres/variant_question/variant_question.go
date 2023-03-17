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

func (r *Repository) CreateVariantQuestion(c context.Context,
	data *variant_question.Create, variantID int) ([]entity.VariantQuestion, error) {

	qs := make([]entity.Question, 0, data.Count)
	err := r.DB.NewSelect().Model(&qs).Column("id").
		Where("topic_id = ?", data.TopicID).
		OrderExpr("random()").Limit(data.Count).Scan(c)
	if err != nil {
		return nil, err
	}

	ms := make([]entity.VariantQuestion, 0, data.Count)
	for _, q := range qs {
		ms = append(ms, entity.VariantQuestion{QuestionID: q.ID, VariantID: variantID})
	}

	_, err = r.DB.NewInsert().Model(&ms).Exec(c)
	if err != nil {
		return nil, err
	}

	return ms, nil
}

func (r *Repository) VariantQuestionByID(c context.Context, id int) (*entity.VariantQuestion, error) {
	m := new(entity.VariantQuestion)
	m.ID = id
	err := r.DB.NewSelect().Model(m).WherePK().Scan(c)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *Repository) VariantQuestionByVariantID(c context.Context, variantID int) ([]entity.VariantQuestion, error) {
	ms := make([]entity.VariantQuestion, 0)
	err := r.DB.NewSelect().Model(&ms).Where("variant_id = ?", variantID).
		Order("id").Scan(c)
	return ms, err
}

func (r *Repository) DeleteVariantQuestion(c context.Context, id int) error {
	m, err := r.VariantQuestionByID(c, id)
	if err != nil {
		return err
	}
	_, err = r.DB.NewDelete().Model(m).WherePK().Exec(c)
	return err
}

func (r *Repository) Repository_() {
	println("just for inherits")
}
