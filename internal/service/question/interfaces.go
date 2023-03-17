package question_srvc

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
)

type Repository interface {
	CreateQuestion(c context.Context, data *Create, userID int) (*entity.Question, error)
	QuestionByID(c context.Context, id int) (*entity.Question, error)
	QuestionByIDs(c context.Context, ids []int) ([]entity.Question, error)
	UpdateQuestion(c context.Context, id int, data *Update) error
	ListQuestion(c context.Context, size, page int) ([]entity.Question, error)
	DeleteQuestion(c context.Context, id, userID int) error
	Repository_()
}
