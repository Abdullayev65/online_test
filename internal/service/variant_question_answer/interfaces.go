package variant_question_answer_srvc

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
)

type Repository interface {
	GetAll(c context.Context, filter *Filter) ([]entity.VariantQuestionAnswer, int, error)
	Create(c context.Context, data *Create) (*entity.VariantQuestionAnswer, error)
	Repository_()
}
