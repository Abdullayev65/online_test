package question_uc

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
	answer_srvc "github.com/Abdullayev65/online_test/internal/service/answer"
	question_srvc "github.com/Abdullayev65/online_test/internal/service/question"
	"mime/multipart"
)

type Question interface {
	GetAll(c context.Context, filter *question_srvc.Filter) ([]entity.Question, int, error)
	GetByID(c context.Context, id int) (*entity.Question, error)
	Create(c context.Context, data *question_srvc.Create) (*entity.Question, error)
	Update(c context.Context, data *question_srvc.Update) error
	Delete(c context.Context, id, userID int) error
	Service_()
}

type Answer interface {
	GetAll(c context.Context, filter *answer_srvc.Filter) ([]entity.Answer, int, error)
	GetByID(c context.Context, id int) (*entity.Answer, error)
	Create(c context.Context, data *answer_srvc.Create) (*entity.Answer, error)
	Update(c context.Context, data *answer_srvc.Update) error
	Delete(ctx context.Context, id, userID int) error
	CorrectAnswerByQuestionID(c context.Context, questionID int) (*entity.Answer, error)
	Service_()
}

type File interface {
	Upload(file *multipart.FileHeader, folder string) (string, error)
	Service_()
}
