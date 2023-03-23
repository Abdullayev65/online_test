package variant_question_answer_repo

import (
	"context"
	"errors"
	"github.com/Abdullayev65/online_test/internal/entity"
	answer_srvc "github.com/Abdullayev65/online_test/internal/service/answer"
	variant_question_answer_srvc "github.com/Abdullayev65/online_test/internal/service/variant_question_answer"
	"github.com/uptrace/bun"
)

type Repository struct {
	*bun.DB
	AnswerRepo answer_srvc.Repository
}

func NewRepository(DB *bun.DB, answerRepo answer_srvc.Repository) *Repository {
	return &Repository{DB: DB, AnswerRepo: answerRepo}
}

func (r Repository) Create(c context.Context, data *variant_question_answer_srvc.Create,
	userID int) (*entity.VariantQuestionAnswer, error) {

	answers, err := r.AnswerRepo.AnswersByQuestionID(c, *data.QuestionID)
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

	m := &entity.VariantQuestionAnswer{QuestionID: *data.QuestionID, VariantID: *data.VariantID,
		AnswerID: *data.AnswerID, IsCorrect: answer.IsCorrect}
	m.CreatedBy = userID

	_, err = r.DB.NewInsert().Model(m).Exec(c)
	return m, err
}

func (r Repository) GetByUserIDAndVariantID(c context.Context, userID,
	variantID int) ([]entity.VariantQuestionAnswer, error) {

	ms := make([]entity.VariantQuestionAnswer, 0)

	err := r.DB.NewSelect().Model(&ms).
		Where("created_by = ? AND variant_id = ?", userID, variantID).
		Scan(c)

	return ms, err
}

func (r Repository) Repository_() {
	println("just for inherits")
}
