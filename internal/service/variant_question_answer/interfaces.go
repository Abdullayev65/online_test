package variant_question_answer_srvc

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
)

type Repository interface {
	Create(c context.Context, data *Create, userID int) (*entity.VariantQuestionAnswer, error)
	GetByUserIDAndVariantID(c context.Context, userID, variantID int) ([]entity.VariantQuestionAnswer, error)
	Repository_()
}
