package answer_repo

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
	answer_srvc "github.com/Abdullayev65/online_test/internal/service/answer"
	"github.com/uptrace/bun"
	"time"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r *Repository) CreateAnswer(c context.Context, data *answer_srvc.Create,
	userID, questionID int) (*entity.Answer, error) {

	m := new(entity.Answer)
	if data.Text != nil {
		m.Text = *data.Text
	}
	if data.Description != nil {
		m.Description = *data.Description
	}
	if data.IsCorrect != nil {
		m.IsCorrect = *data.IsCorrect
	}
	m.CreatedBy = userID
	m.QuestionID = questionID
	_, err := r.DB.NewInsert().Model(m).Exec(c)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *Repository) AnswerByID(c context.Context, id int) (*entity.Answer, error) {

	m := new(entity.Answer)
	m.ID = id
	err := r.DB.NewSelect().Model(m).WherePK().Scan(c)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *Repository) UpdateAnswer(c context.Context, data *answer_srvc.Update) error {

	m := new(entity.Answer)
	m.ID = data.ID

	if data.Text != nil {
		m.Text = *data.Text
	}
	if data.Description != nil {
		m.Description = *data.Description
	}
	if data.IsCorrect != nil {
		m.IsCorrect = *data.IsCorrect
	}

	_, err := r.DB.NewUpdate().Model(m).WherePK().Exec(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) AnswersByQuestionID(c context.Context, questionID int) ([]entity.Answer, error) {

	ms := make([]entity.Answer, 0)
	err := r.DB.NewSelect().Model(&ms).Where("question_id = ?", questionID).Scan(c)

	return ms, err
}

func (r *Repository) DeleteAnswer(c context.Context, id, userID int) error {
	m, err := r.AnswerByID(c, id)
	if err != nil {
		return err
	}
	m.DeletedAt = time.Now()
	m.DeletedBy = userID
	_, err = r.DB.NewUpdate().Model(m).WherePK().Exec(c)
	return err
}

func (r *Repository) CorrectAnswerByQuestionID(c context.Context,
	questionID int) (*entity.Answer, error) {

	m := new(entity.Answer)
	err := r.DB.NewSelect().Model(m).
		Where("question_id = ? AND is_correct = true", questionID).Scan(c)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *Repository) Repository_() {
	println("just for inherits")
}
