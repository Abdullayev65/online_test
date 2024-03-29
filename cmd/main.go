package main

import (
	"github.com/Abdullayev65/online_test/internal/controller/middleware"
	question_controller "github.com/Abdullayev65/online_test/internal/controller/v1/question"
	topic_controller "github.com/Abdullayev65/online_test/internal/controller/v1/topic"
	user_controller "github.com/Abdullayev65/online_test/internal/controller/v1/user"
	variant_controller "github.com/Abdullayev65/online_test/internal/controller/v1/variant"
	variant_question_answer_ctrl "github.com/Abdullayev65/online_test/internal/controller/v1/variant_question_answer"
	"github.com/Abdullayev65/online_test/internal/pkg/repository/postgres"
	answer_repo "github.com/Abdullayev65/online_test/internal/repository/postgres/answer"
	question_repo "github.com/Abdullayev65/online_test/internal/repository/postgres/question"
	topic_repo "github.com/Abdullayev65/online_test/internal/repository/postgres/topic"
	user_repo "github.com/Abdullayev65/online_test/internal/repository/postgres/user"
	variant_repo "github.com/Abdullayev65/online_test/internal/repository/postgres/variant"
	variant_question_repo "github.com/Abdullayev65/online_test/internal/repository/postgres/variant_question"
	variant_question_answer_repo "github.com/Abdullayev65/online_test/internal/repository/postgres/variant_question_answer"
	answer_srvc "github.com/Abdullayev65/online_test/internal/service/answer"
	file_srvc "github.com/Abdullayev65/online_test/internal/service/file"
	question_srvc "github.com/Abdullayev65/online_test/internal/service/question"
	topic_service "github.com/Abdullayev65/online_test/internal/service/topic"
	user_service "github.com/Abdullayev65/online_test/internal/service/user"
	variant_srvc "github.com/Abdullayev65/online_test/internal/service/variant"
	variant_question_srvc "github.com/Abdullayev65/online_test/internal/service/variant_question"
	variant_question_answer_srvc "github.com/Abdullayev65/online_test/internal/service/variant_question_answer"
	question_uc "github.com/Abdullayev65/online_test/internal/usecase/question"
	topic_uc "github.com/Abdullayev65/online_test/internal/usecase/topic"
	user_usecase "github.com/Abdullayev65/online_test/internal/usecase/user"
	variant_uc "github.com/Abdullayev65/online_test/internal/usecase/variant"
	variant_question_answer_uc "github.com/Abdullayev65/online_test/internal/usecase/variant_question_answer"
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
	variantQuestionAnswerRepo := variant_question_answer_repo.NewRepository(db, answerRepo)

	// service
	fileService := file_srvc.New()
	userService := user_service.NewService(userRepo)
	topicService := topic_service.NewService(topicRepo)
	questionService := question_srvc.NewService(questionRepo)
	answerService := answer_srvc.NewService(answerRepo)
	variantService := variant_srvc.NewService(variantRepo)
	variantQuestionService := variant_question_srvc.NewService(variantQuestionRepo)
	variantQuestionAnswerService := variant_question_answer_srvc.NewService(variantQuestionAnswerRepo)

	// use case
	userUC := user_usecase.NewUseCase(userService, token)
	topicUC := topic_uc.NewUseCase(topicService)
	questionUC := question_uc.NewUseCase(questionService, answerService, fileService)
	variantUC := variant_uc.NewUseCase(variantService, variantQuestionService,
		questionService, answerService)
	variantQuestionAnswerUC := variant_question_answer_uc.NewUseCase(variantQuestionAnswerService,
		questionService, answerService, variantService)

	// middleware
	MW := middleware.New(token, userUC)

	// controller
	userController := user_controller.NewController(userUC)
	topicController := topic_controller.NewController(topicUC)
	questionController := question_controller.NewController(questionUC)
	variantController := variant_controller.NewController(variantUC)
	variantQuestionAnswerController := variant_question_answer_ctrl.NewController(variantQuestionAnswerUC)

	// API routers
	router := gin.Default()

	router.Static("/media", "./media")

	user := router.Group("user")
	{
		user.POST("/sign-up", userController.CreateUser)
		user.POST("/sign-in", userController.SignIn)
		user.GET("/me", MW.UserIDFromToken, userController.GetUserMe)
		user.PUT("/", MW.UserIDFromToken, userController.UpdateUser)
	}

	topic := router.Group("topic")
	{
		topic.POST("/", MW.UserIDFromToken, topicController.CreateTopic)
		topic.GET("/list", MW.SetIntFromQuery("page", "size"),
			topicController.GetTopicList)
		topic.GET("/:id", MW.SetIntFromParam("id"),
			topicController.GetTopicDetail)
		topic.PUT("/:id", MW.SetIntFromParam("id"),
			topicController.UpdateTopic)
		topic.DELETE("/:id", MW.SetIntFromParam("id"),
			MW.UserIDFromToken, topicController.DeleteTopic)
	}

	question := router.Group("question")
	{
		question.POST("/", MW.UserIDFromToken, questionController.CreateQuestion)
		question.GET("/list", MW.SetIntFromQuery("page", "size"), questionController.GetQuestionList)
		question.GET("/:id", MW.SetIntFromParam("id"), questionController.GetQuestionDetail)
		question.PUT("/:id", MW.SetIntFromParam("id"), questionController.UpdateQuestion)
		question.DELETE("/:id", MW.SetIntFromParam("id"), questionController.DeleteQuestion)
	}

	variant := router.Group("variant")
	{
		variant.POST("/generate", MW.UserIDFromToken, variantController.GenerateVariant)
		variant.GET("/list", MW.SetIntFromQuery("page", "size"), variantController.GetVariantList)
		variant.GET("/:id", MW.SetIntFromParam("id"), variantController.GetVariantDetail)
		variant.DELETE("/:id", MW.SetIntFromParam("id"), MW.UserIDFromToken,
			variantController.DeleteVariant)
	}

	questionAnswer := router.Group("question-answer")
	{
		questionAnswer.POST("/test", MW.UserIDFromToken, variantQuestionAnswerController.CreateVariantAnswer)
		questionAnswer.GET("/my-variant-answer/:variantID", MW.UserIDFromToken,
			MW.SetIntFromParam("variantID"), variantQuestionAnswerController.GetMyVariantAnswerDetail)
		questionAnswer.GET("/variant-answer", MW.SetIntFromQuery("variant_id", "user_id"),
			variantQuestionAnswerController.GetVariantAnswerDteail)
	}

	// * * * * * * * * * * * * *
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
