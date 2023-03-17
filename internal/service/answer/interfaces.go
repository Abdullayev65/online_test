package answer_srvc

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
)

type Repository interface {
	CreateAnswer(c context.Context, data *Create, userID, questionID int) (*entity.Answer, error)
	AnswerByID(c context.Context, id int) (*entity.Answer, error)
	UpdateAnswer(c context.Context, data *Update) error
	AnswersByQuestionID(c context.Context, questionID int) ([]entity.Answer, error)
	DeleteAnswer(c context.Context, id, userID int) error
	Repository_()
}
