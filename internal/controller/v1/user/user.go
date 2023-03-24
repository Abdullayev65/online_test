package user

import (
	"github.com/Abdullayev65/online_test/internal/controller/response"
	user_service "github.com/Abdullayev65/online_test/internal/service/user"
	user_usecase "github.com/Abdullayev65/online_test/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	useCase *user_usecase.UseCase
}

func NewController(useCase *user_usecase.UseCase) *Controller {
	return &Controller{useCase: useCase}
}

func (cc *Controller) CreateUser(ctx *gin.Context) {
	var createUser user_service.Create
	ctx.Bind(&createUser)
	user, err := cc.useCase.CreateUser(ctx, &createUser)
	if err != nil {
		response.FailErr(ctx, err)
		return
	}
	response.Success(ctx, user)
}
func (cc *Controller) SignIn(ctx *gin.Context) {
	var sign user_service.SignIn
	ctx.Bind(&sign)

	token, err := cc.useCase.UserGenerateToken(ctx, &sign)
	if err != nil {
		response.FailErr(ctx, err)
		return
	}
	response.Success(ctx, token)
}
func (cc *Controller) GetUserMe(ctx *gin.Context) {
	userID := ctx.GetInt("userID")
	userDTO, err := cc.useCase.GetUserDetail(ctx, userID)
	if err != nil {
		response.FailErr(ctx, err)
		return
	}

	response.Success(ctx, userDTO)
}
func (cc *Controller) UpdateUser(ctx *gin.Context) {
	userID := ctx.GetInt("userID")
	var update user_service.Update
	ctx.Bind(&update)
	err := cc.useCase.UpdateUser(ctx, userID, &update)
	if err != nil {
		response.FailErr(ctx, err)
		return
	}
	response.Success(ctx, userID)
}
