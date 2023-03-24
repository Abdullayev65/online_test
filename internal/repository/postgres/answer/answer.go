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

func (r *Repository) GetAll(c context.Context, filter *answer_srvc.Filter) ([]entity.Answer, int, error) {
	var list []entity.Answer
	q := r.NewSelect().Model(&list)

	if filter.Limit != nil {
		q.Limit(*filter.Limit)
	}

	if filter.Offset != nil {
		q.Offset(*filter.Offset)
	}

	if filter.QuestionID != nil {
		q.WhereGroup(" and ", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Where("question_id = ?", *filter.QuestionID)
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
func (r *Repository) GetByID(c context.Context, id int) (*entity.Answer, error) {

	m := new(entity.Answer)
	m.ID = id
	err := r.DB.NewSelect().Model(m).WherePK().Scan(c)
	if err != nil {
		return nil, err
	}
	return m, nil
}
func (r *Repository) Create(c context.Context, data *answer_srvc.Create) (*entity.Answer, error) {

	m := new(entity.Answer)
	m.Text = data.Text
	m.Description = data.Description
	m.IsCorrect = data.IsCorrect
	m.CreatedBy = data.UserID
	m.QuestionID = data.QuestionID
	_, err := r.DB.NewInsert().Model(m).Exec(c)
	if err != nil {
		return nil, err
	}
	return m, nil
}
func (r *Repository) Update(c context.Context, data *answer_srvc.Update) error {

	m := new(entity.Answer)
	m.ID = data.ID

	m.Text = data.Text
	m.Description = data.Description
	m.IsCorrect = data.IsCorrect

	_, err := r.DB.NewUpdate().Model(m).WherePK().Exec(c)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) Delete(c context.Context, id, userID int) error {
	m, err := r.GetByID(c, id)
	if err != nil {
		return err
	}
	m.DeletedAt = time.Now()
	m.DeletedBy = &userID
	_, err = r.DB.NewUpdate().Model(m).WherePK().Exec(c)
	return err
}

// spesfic functions
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
