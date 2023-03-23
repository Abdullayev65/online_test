package question_uc

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
	answer_srvc "github.com/Abdullayev65/online_test/internal/service/answer"
	question_srvc "github.com/Abdullayev65/online_test/internal/service/question"
)

type Question interface {
	CreateQuestion(c context.Context, data *question_srvc.Create) (*entity.Question, error)
	QuestionByID(c context.Context, id int) (*entity.Question, error)
	UpdateQuestion(c context.Context, id int, data *question_srvc.Update) error
	ListQuestion(c context.Context, size, page int) ([]entity.Question, error)
	DeleteQuestion(c context.Context, id, userID int) error
	Service_()
}

type Answer interface {
	CreateAnswer(c context.Context, data *answer_srvc.Create, userID, questionID int) (*entity.Answer, error)
	AnswerByID(c context.Context, id int) (*entity.Answer, error)
	UpdateAnswer(c context.Context, data *answer_srvc.Update) error
	AnswersByQuestionID(c context.Context, questionID int) ([]entity.Answer, error)
	DeleteAnswer(c context.Context, id, userID int) error
	Service_()
}
