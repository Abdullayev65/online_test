package variant_uc

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
	variant_srvc "github.com/Abdullayev65/online_test/internal/service/variant"
	variant_question_srvc "github.com/Abdullayev65/online_test/internal/service/variant_question"
)

type Variant interface {
	CreateVariant(c context.Context, data *variant_srvc.Create, userID int) (*entity.Variant, error)
	VariantByID(c context.Context, id int) (*entity.Variant, error)
	ListVariant(c context.Context, size, page int) ([]entity.Variant, error)
	DeleteVariant(c context.Context, id, userID int) error
	Service_()
}
type VariantQuestion interface {
	CreateVariantQuestion(c context.Context, data *variant_question_srvc.Create, variantID int) ([]entity.VariantQuestion, error)
	VariantQuestionByID(c context.Context, id int) (*entity.VariantQuestion, error)
	VariantQuestionByVariantID(c context.Context, variantID int) ([]entity.VariantQuestion, error)
	DeleteVariantQuestion(c context.Context, id int) error
	Service_()
}
type Question interface {
	QuestionByIDs(c context.Context, ids []int) ([]entity.Question, error)
	Service_()
}
type Answer interface {
	AnswersByQuestionID(c context.Context, questionID int) ([]entity.Answer, error)
	Service_()
}
