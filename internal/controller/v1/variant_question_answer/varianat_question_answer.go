package variant_question_answer_ctrl

import (
	"github.com/Abdullayev65/online_test/internal/controller/response"
	srvc "github.com/Abdullayev65/online_test/internal/service/variant_question_answer"
	uc "github.com/Abdullayev65/online_test/internal/usecase/variant_question_answer"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	useCase *uc.UseCase
}

func NewController(useCase *uc.UseCase) *Controller {
	return &Controller{useCase: useCase}
}

func (cc *Controller) Create(ctx *gin.Context) {
	userID := ctx.GetInt("userID")
	var create srvc.Create
	ctx.Bind(&create)
	create.UserID = &userID
	dto, err := cc.useCase.Create(ctx, &create)
	if err != nil {
		response.FailErr(ctx, err)
		return
	}
	response.Success(ctx, dto)
}

func (cc *Controller) GetMyVariantAnswer(ctx *gin.Context) {
	userID := ctx.GetInt("userID")
	variantID := ctx.GetInt("variantID")

	dto, err := cc.useCase.VariantAnswerByUserIDAndVariantID(ctx, userID, variantID)
	if err != nil {
		response.FailErr(ctx, err)
		return
	}
	response.Success(ctx, dto)
}
func (cc *Controller) GetUserVariantAnswer(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	variantID := ctx.GetInt("variant_id")

	dto, err := cc.useCase.VariantAnswerByUserIDAndVariantID(ctx, userID, variantID)
	if err != nil {
		response.FailErr(ctx, err)
		return
	}
	response.Success(ctx, dto)
}
