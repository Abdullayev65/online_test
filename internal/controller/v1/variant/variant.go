package variant_controller

import (
	"github.com/Abdullayev65/online_test/internal/controller/response"
	variant_srvc "github.com/Abdullayev65/online_test/internal/service/variant"
	variant_uc "github.com/Abdullayev65/online_test/internal/usecase/variant"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	useCase *variant_uc.UseCase
}

func NewController(useCase *variant_uc.UseCase) *Controller {
	return &Controller{useCase: useCase}
}

func (cc *Controller) GenerateVariant(ctx *gin.Context) {
	userID := ctx.GetInt("userID")
	var create variant_srvc.Create
	ctx.Bind(&create)
	create.UserID = &userID

	topicDTO, err := cc.useCase.GenerateVariant(ctx, &create)
	if err != nil {
		response.FailErr(ctx, err)
		return
	}
	response.Success(ctx, topicDTO)
}

func (cc *Controller) GetVariantDetailByID(ctx *gin.Context) {
	id := ctx.GetInt("id")

	detail, err := cc.useCase.GetVariantDetailByID(ctx, id)
	if err != nil {
		response.FailErr(ctx, err)
		return
	}

	response.Success(ctx, detail)
}

func (cc *Controller) DeleteVariant(ctx *gin.Context) {
	id := ctx.GetInt("id")
	userID := ctx.GetInt("userID")

	err := cc.useCase.DeleteVariant(ctx, id, userID)
	if err != nil {
		response.FailErr(ctx, err)
		return
	}
	response.Success(ctx, "deleted!")
}

func (cc *Controller) List(ctx *gin.Context) {
	page := ctx.GetInt("page")
	size := ctx.GetInt("size")

	lists, err := cc.useCase.ListVariant(ctx, size, page)
	if err != nil {
		response.FailErr(ctx, err)
		return
	}

	response.Success(ctx, &lists)
}
