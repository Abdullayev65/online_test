package topic_controller

import (
	"github.com/Abdullayev65/online_test/internal/controller/response"
	topic_service "github.com/Abdullayev65/online_test/internal/service/topic"
	topic_uc "github.com/Abdullayev65/online_test/internal/usecase/topic"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	useCase *topic_uc.UseCase
}

func NewController(useCase *topic_uc.UseCase) *Controller {
	return &Controller{useCase: useCase}
}

func (cc *Controller) CreateTopic(ctx *gin.Context) {
	userID := ctx.GetInt("userID")
	var create topic_service.Create
	ctx.Bind(&create)
	topicDTO, err := cc.useCase.AdminCreateTopic(ctx, &create, userID)
	if err != nil {
		response.FailErr(ctx, err)
		return
	}
	response.Success(ctx, topicDTO)
}
func (cc *Controller) GetTopic(ctx *gin.Context) {
	id := ctx.GetInt("id")
	topicDTO, err := cc.useCase.AdminGetTopicDetail(ctx, id)
	if err != nil {
		response.FailErr(ctx, err)
		return
	}
	response.Success(ctx, topicDTO)
}
func (cc *Controller) UpdateTopic(ctx *gin.Context) {
	id := ctx.GetInt("id")
	var update topic_service.Update
	ctx.Bind(&update)
	err := cc.useCase.AdminUpdateTopic(ctx, id, &update)
	if err != nil {
		response.FailErr(ctx, err)
		return
	}
	response.Success(ctx, id)
}
func (cc *Controller) ListOfTopic(ctx *gin.Context) {
	page := ctx.GetInt("page")
	size := ctx.GetInt("size")
	topicList, err := cc.useCase.GetListTopic(ctx, size, page)
	if err != nil {
		response.FailErr(ctx, err)
		return
	}
	response.Success(ctx, topicList)
}
func (cc *Controller) DeleteTopic(ctx *gin.Context) {
	userID := ctx.GetInt("userID")
	id := ctx.GetInt("id")
	err := cc.useCase.AdminDeleteTopic(ctx, id, userID)
	if err != nil {
		response.FailErr(ctx, err)
		return
	}
	response.Success(ctx, id)
}
