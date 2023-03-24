package variant_question_answer_uc

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
	answer_srvc "github.com/Abdullayev65/online_test/internal/service/answer"
	variant_question_answer_srvc "github.com/Abdullayev65/online_test/internal/service/variant_question_answer"
)

type VariantQuestionAnswer interface {
	Create(c context.Context, data *variant_question_answer_srvc.Create) (*entity.VariantQuestionAnswer, error)
	GetAll(c context.Context, filter *variant_question_answer_srvc.Filter) ([]entity.VariantQuestionAnswer, int, error)
	Service_()
}
type Question interface {
	GetByID(c context.Context, id int) (*entity.Question, error)
	Service_()
}
type Answer interface {
	GetAll(c context.Context, filter *answer_srvc.Filter) ([]entity.Answer, int, error)
	CorrectAnswerByQuestionID(c context.Context, questionID int) (*entity.Answer, error)
	Service_()
}
type Variant interface {
	GetByID(c context.Context, id int) (*entity.Variant, error)
	Service_()
}
