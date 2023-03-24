package variant_uc

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
	answer_srvc "github.com/Abdullayev65/online_test/internal/service/answer"
	question_srvc "github.com/Abdullayev65/online_test/internal/service/question"
	variant_srvc "github.com/Abdullayev65/online_test/internal/service/variant"
	variant_question_srvc "github.com/Abdullayev65/online_test/internal/service/variant_question"
)

type Variant interface {
	GetAll(c context.Context, filter *variant_srvc.Filter) ([]entity.Variant, int, error)
	GetByID(c context.Context, id int) (*entity.Variant, error)
	Create(c context.Context, data *variant_srvc.Create) (*entity.Variant, error)
	Delete(c context.Context, id, userID int) error
	Service_()
}
type VariantQuestion interface {
	GetAll(c context.Context, filter *variant_question_srvc.Filter) ([]entity.VariantQuestion, int, error)
	GetByID(c context.Context, id int) (*entity.VariantQuestion, error)
	Create(c context.Context, data *variant_question_srvc.Create) ([]entity.VariantQuestion, error)
	Delete(c context.Context, id int) error
	Service_()
}
type Question interface {
	GetAll(c context.Context, filter *question_srvc.Filter) ([]entity.Question, int, error)
	Service_()
}
type Answer interface {
	GetAll(c context.Context, filter *answer_srvc.Filter) ([]entity.Answer, int, error)
	Service_()
}
