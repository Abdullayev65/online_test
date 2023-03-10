package main

import (
	"github.com/Abdullayev65/online_test/internal/controller/middleware"
	topic_controller "github.com/Abdullayev65/online_test/internal/controller/v1/topic"
	user_controller "github.com/Abdullayev65/online_test/internal/controller/v1/user"
	"github.com/Abdullayev65/online_test/internal/pkg/repository/postgres"
	topic_repo "github.com/Abdullayev65/online_test/internal/repository/postgres/topic"
	user_repo "github.com/Abdullayev65/online_test/internal/repository/postgres/user"
	topic_service "github.com/Abdullayev65/online_test/internal/service/topic"
	user_service "github.com/Abdullayev65/online_test/internal/service/user"
	topic_uc "github.com/Abdullayev65/online_test/internal/usecase/topic"
	user_usecase "github.com/Abdullayev65/online_test/internal/usecase/user"
	"github.com/Abdullayev65/online_test/internal/utill"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	db := postgres.New()
	// _ _ _ _ _ _ _ _ _ _
	token := utill.NewToken("salat", 26*time.Hour)
	userRepo := user_repo.NewRepository(db)
	userService := user_service.NewService(userRepo)
	userUC := user_usecase.NewUseCase(userService, token)
	userController := user_controller.NewController(userUC)
	MW := middleware.New(token, userUC)

	topicRepo := topic_repo.NewRepository(db)
	topicService := topic_service.NewService(topicRepo)
	topicUC := topic_uc.NewUseCase(topicService)
	topicController := topic_controller.NewController(topicUC)

	// API routers
	router := gin.Default()

	user := router.Group("/user")
	{
		user.POST("/sign-up", userController.CreateUser)
		user.POST("/sign-in", userController.SignIn)
		user.GET("/me", MW.UserIDFromToken, userController.UserMe)
		user.PUT("/", MW.UserIDFromToken, userController.UpdateUser)
	}

	topic := router.Group("topic")
	{
		topic.POST("/", MW.UserIDFromToken, topicController.CreateTopic)
		topic.GET("/list", MW.SetIntFromQuery("page", "size"),
			topicController.ListOfTopic)
		topic.GET("/:id", MW.SetIntFromParam("id"),
			topicController.GetTopic)
		topic.PUT("/:id", MW.SetIntFromParam("id"),
			topicController.UpdateTopic)
		topic.DELETE("/:id", MW.SetIntFromParam("id"),
			MW.UserIDFromToken, topicController.DeleteTopic)
	}

	question := router.Group("question")
	{
		question.POST("/", userController.CreateUser)
		question.GET("/list", userController.SignIn)
		question.GET("/:id", MW.UserIDFromToken, userController.UserMe)
		question.PUT("/", MW.UserIDFromToken, userController.UpdateUser)
		question.DELETE("/", MW.UserIDFromToken, userController.UpdateUser)
	}

	// * * * * * * * * * * * * *
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
