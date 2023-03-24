package variant_question_answer_repo

import (
	"context"
	"errors"
	"github.com/Abdullayev65/online_test/internal/entity"
	answer_srvc "github.com/Abdullayev65/online_test/internal/service/answer"
	srvc "github.com/Abdullayev65/online_test/internal/service/variant_question_answer"
	"github.com/uptrace/bun"
)

type Repository struct {
	*bun.DB
	AnswerRepo answer_srvc.Repository
}

func NewRepository(DB *bun.DB, answerRepo answer_srvc.Repository) *Repository {
	return &Repository{DB: DB, AnswerRepo: answerRepo}
}

//impls

func (r Repository) GetAll(c context.Context, filter *srvc.Filter) ([]entity.VariantQuestionAnswer, int, error) {
	var list []entity.VariantQuestionAnswer
	q := r.NewSelect().Model(&list)

	if filter.UserID != nil {
		q.WhereGroup(" AND ", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Where("created_by = ?", *filter.UserID)
		})
	}

	if filter.VariantID != nil {
		q.WhereGroup(" AND ", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Where("variant_id = ?", *filter.VariantID)
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

func (r Repository) Create(c context.Context, data *srvc.Create) (*entity.VariantQuestionAnswer, error) {
	answers, _, err := r.AnswerRepo.GetAll(c, &answer_srvc.Filter{QuestionID: data.QuestionID})
	if err != nil {
		return nil, err
	}

	var answer *entity.Answer
	for _, a := range answers {
		if a.ID == *data.AnswerID {
			ans := a
			answer = &ans
		}
	}
	if answer == nil {
		return nil, errors.New("answer in question not found")
	}

	m := &entity.VariantQuestionAnswer{QuestionID: data.QuestionID, VariantID: data.VariantID,
		AnswerID: data.AnswerID, IsCorrect: answer.IsCorrect}
	m.CreatedBy = data.UserID

	_, err = r.DB.NewInsert().Model(m).Exec(c)
	return m, err
}

func (r Repository) Repository_() {
	println("just for inherits")
}
