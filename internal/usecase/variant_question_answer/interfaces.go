package variant_question_answer_uc

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
	variant_question_answer_srvc "github.com/Abdullayev65/online_test/internal/service/variant_question_answer"
)

type VariantQuestionAnswer interface {
	Create(c context.Context, data *variant_question_answer_srvc.Create, userID int) (*entity.VariantQuestionAnswer, error)
	GetByUserIDAndVariantID(c context.Context, userID, variantID int) ([]entity.VariantQuestionAnswer, error)
	Service_()
}
type Question interface {
	QuestionByID(c context.Context, id int) (*entity.Question, error)
	Service_()
}
type Answer interface {
	AnswersByQuestionID(c context.Context, questionID int) ([]entity.Answer, error)
	CorrectAnswerByQuestionID(c context.Context, questionID int) (*entity.Answer, error)
	Service_()
}
type Variant interface {
	VariantByID(c context.Context, id int) (*entity.Variant, error)
	Service_()
}
