package variant_question_srvc

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
)

type Repository interface {
	CreateVariantQuestion(c context.Context, data *Create, variantID int) ([]entity.VariantQuestion, error)
	VariantQuestionByID(c context.Context, id int) (*entity.VariantQuestion, error)
	VariantQuestionByVariantID(c context.Context, variantID int) ([]entity.VariantQuestion, error)
	DeleteVariantQuestion(c context.Context, id int) error
	Repository_()
}
