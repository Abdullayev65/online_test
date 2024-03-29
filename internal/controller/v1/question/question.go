package question_controller

import (
	"github.com/Abdullayev65/online_test/internal/controller/response"
	question_srvc "github.com/Abdullayev65/online_test/internal/service/question"
	question_uc "github.com/Abdullayev65/online_test/internal/usecase/question"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	useCase *question_uc.UseCase
}

func NewController(useCase *question_uc.UseCase) *Controller {
	return &Controller{useCase: useCase}
}

func (cc *Controller) CreateQuestion(ctx *gin.Context) {
	userID := ctx.GetInt("userID")
	var create question_srvc.Create
	create.UserId = &userID
	ctx.Bind(&create)
	topicDTO, err := cc.useCase.CreateQuestion(ctx, &create)
	if err != nil {
		response.FailErr(ctx, err)
		return
	}
	response.Success(ctx, topicDTO)
}
func (cc *Controller) GetQuestionDetail(ctx *gin.Context) {
	id := ctx.GetInt("id")
	questionDetail, err := cc.useCase.GetQuestionDetail(ctx, id)
	if err != nil {
		response.FailErr(ctx, err)
		return
	}
	response.Success(ctx, questionDetail)
}
func (cc *Controller) UpdateQuestion(ctx *gin.Context) {
	userID := ctx.GetInt("userID")
	id := ctx.GetInt("id")
	var update question_srvc.Update
	ctx.Bind(&update)

	update.ID = id
	update.UserID = &userID

	err := cc.useCase.UpdateQuestion(ctx, &update)
	if err != nil {
		response.FailErr(ctx, err)
		return
	}
	response.Success(ctx, id)
}
func (cc *Controller) GetQuestionList(ctx *gin.Context) {
	page := ctx.GetInt("page")
	size := ctx.GetInt("size")
	topicList, err := cc.useCase.GetListQuestion(ctx, size, page)
	if err != nil {
		response.FailErr(ctx, err)
		return
	}
	response.Success(ctx, topicList)
}
func (cc *Controller) DeleteQuestion(ctx *gin.Context) {
	userID := ctx.GetInt("userID")
	id := ctx.GetInt("id")
	err := cc.useCase.DeleteQuestion(ctx, id, userID)
	if err != nil {
		response.FailErr(ctx, err)
		return
	}
	response.Success(ctx, id)
}
