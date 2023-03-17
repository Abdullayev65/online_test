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

func (r *Repository) CreateQuestion(c context.Context,
	data *question_srvc.Create, userID int) (*entity.Question, error) {

	m := new(entity.Question)
	if data.Text != nil {
		m.Text = *data.Text
	}
	if data.Description != nil {
		m.Description = *data.Description
	}
	if data.TopicID != nil {
		m.TopicID = *data.TopicID
	}
	m.CreatedBy = userID
	_, err := r.DB.NewInsert().Model(m).Exec(c)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *Repository) QuestionByID(c context.Context, id int) (*entity.Question, error) {
	m := new(entity.Question)
	m.ID = id
	err := r.DB.NewSelect().Model(m).WherePK().Scan(c)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *Repository) QuestionByIDs(c context.Context, ids []int) ([]entity.Question, error) {
	ms := make([]entity.Question, 0, len(ids))
	err := r.DB.NewSelect().Model(&ms).Where("id in (?)", bun.In(ids)).Scan(c)
	return ms, err
}

func (r *Repository) UpdateQuestion(c context.Context, id int,
	data *question_srvc.Update) error {

	m := new(entity.Question)
	m.ID = id

	if data.Text != nil {
		m.Text = *data.Text
	}
	if data.Description != nil {
		m.Description = *data.Description
	}
	if data.TopicID != nil {
		m.TopicID = *data.TopicID
	}

	_, err := r.DB.NewUpdate().Model(m).WherePK().Exec(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) ListQuestion(c context.Context,
	size, page int) ([]entity.Question, error) {

	ms := make([]entity.Question, 0)
	err := r.DB.NewSelect().Model(&ms).Limit(size).
		Offset((size - 1) * page).Order("id").Scan(c)
	if err != nil {
		return nil, err
	}
	return ms, nil
}

func (r *Repository) DeleteQuestion(c context.Context, id, userID int) error {
	m, err := r.QuestionByID(c, id)
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
