package main

import (
	"github.com/Abdullayev65/online_test/internal/controller/middleware"
	question_controller "github.com/Abdullayev65/online_test/internal/controller/v1/question"
	topic_controller "github.com/Abdullayev65/online_test/internal/controller/v1/topic"
	user_controller "github.com/Abdullayev65/online_test/internal/controller/v1/user"
	variant_controller "github.com/Abdullayev65/online_test/internal/controller/v1/variant"
	"github.com/Abdullayev65/online_test/internal/pkg/repository/postgres"
	answer_repo "github.com/Abdullayev65/online_test/internal/repository/postgres/answer"
	question_repo "github.com/Abdullayev65/online_test/internal/repository/postgres/question"
	topic_repo "github.com/Abdullayev65/online_test/internal/repository/postgres/topic"
	user_repo "github.com/Abdullayev65/online_test/internal/repository/postgres/user"
	variant_repo "github.com/Abdullayev65/online_test/internal/repository/postgres/variant"
	variant_question_repo "github.com/Abdullayev65/online_test/internal/repository/postgres/variant_question"
	answer_srvc "github.com/Abdullayev65/online_test/internal/service/answer"
	question_srvc "github.com/Abdullayev65/online_test/internal/service/question"
	topic_service "github.com/Abdullayev65/online_test/internal/service/topic"
	user_service "github.com/Abdullayev65/online_test/internal/service/user"
	variant_srvc "github.com/Abdullayev65/online_test/internal/service/variant"
	variant_question_srvc "github.com/Abdullayev65/online_test/internal/service/variant_question"
	question_uc "github.com/Abdullayev65/online_test/internal/usecase/question"
	topic_uc "github.com/Abdullayev65/online_test/internal/usecase/topic"
	user_usecase "github.com/Abdullayev65/online_test/internal/usecase/user"
	variant_uc "github.com/Abdullayev65/online_test/internal/usecase/variant"
	"github.com/Abdullayev65/online_test/internal/utill"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	_ = variant_srvc.Create{}
	db := postgres.New()
	// _ _ _ _ _ _ _ _ _ _
	token := utill.NewToken("salat", 300*time.Hour)

	// repo
	userRepo := user_repo.NewRepository(db)
	topicRepo := topic_repo.NewRepository(db)
	questionRepo := question_repo.NewRepository(db)
	answerRepo := answer_repo.NewRepository(db)
	variantRepo := variant_repo.NewRepository(db)
	variantQuestionRepo := variant_question_repo.NewRepository(db)

	// srvc
	userService := user_service.NewService(userRepo)
	topicService := topic_service.NewService(topicRepo)
	questionService := question_srvc.NewService(questionRepo)
	answerService := answer_srvc.NewService(answerRepo)
	variantService := variant_srvc.NewService(variantRepo)
	variantQuestionService := variant_question_srvc.NewService(variantQuestionRepo)

	// use case
	userUC := user_usecase.NewUseCase(userService, token)
	topicUC := topic_uc.NewUseCase(topicService)
	questionUC := question_uc.NewUseCase(questionService, answerService)
	variantUC := variant_uc.NewUseCase(variantService, variantQuestionService, questionService)

	// middleware
	MW := middleware.New(token, userUC)

	// controller
	userController := user_controller.NewController(userUC)
	topicController := topic_controller.NewController(topicUC)
	questionController := question_controller.NewController(questionUC)
	variantController := variant_controller.NewController(variantUC)

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
		question.POST("/", MW.UserIDFromToken, questionController.CreateQuestion)
		question.GET("/list", MW.SetIntFromQuery("page", "size"), questionController.List)
		question.GET("/:id", MW.SetIntFromParam("id"), questionController.GetQuestionDetailByID)
		question.PUT("/:id", MW.SetIntFromParam("id"), questionController.UpdateQuestion)
		question.DELETE("/:id", MW.SetIntFromParam("id"), questionController.DeleteQuestion)
	}

	variant := router.Group("variant")
	{
		variant.POST("/generate", MW.UserIDFromToken, variantController.GenerateVariant)
		variant.GET("/list", MW.SetIntFromQuery("page", "size"), variantController.List)
		variant.GET("/:id", MW.SetIntFromParam("id"), variantController.GetVariantDetailByID)
		variant.DELETE("/:id", MW.SetIntFromParam("id"), MW.UserIDFromToken,
			variantController.DeleteVariant)
	}

	// * * * * * * * * * * * * *
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
