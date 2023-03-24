package answer_srvc

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
)

type Repository interface {
	GetAll(c context.Context, filter *Filter) ([]entity.Answer, int, error)
	GetByID(c context.Context, id int) (*entity.Answer, error)
	Create(c context.Context, data *Create) (*entity.Answer, error)
	Update(c context.Context, data *Update) error
	Delete(ctx context.Context, id, userID int) error
	CorrectAnswerByQuestionID(c context.Context, questionID int) (*entity.Answer, error)
	Repository_()
}
