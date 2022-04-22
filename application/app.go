package application

import (
	"github.com/Saucon/errcntrct"
	"github.com/gin-gonic/gin"
	"log"
	"number1/config/db"
	"number1/config/env"
	"number1/controllers"
	"number1/middlewares"
	"number1/repo/login"
	"number1/repo/transaction"
	error2 "number1/usecase/error"
	uc "number1/usecase/login"
	tr "number1/usecase/transaction"
)

func StartApp() {
	router := gin.New()
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"responseCode": "404", "responseMessage": "Page not found"})
	})
	if err := errcntrct.InitContract(env.Config.JSONPathFile); err != nil {
		log.Fatal(err, "main : init contract", nil)
	}

	DB := db.DB
	router.Use(gin.Recovery())

	loginRepo := login.NewLoginRepoImpl(DB)
	transRepo := transaction.NewTransReportRepo(DB)

	loginUsecase := uc.NewLoginUsecaseImpl(loginRepo)
	transUsecase := tr.NewTransReportUsecase(transRepo)
	errorHandlerUseCase := error2.NewErrorHandlerUsecase()

	newRoute := router.Group("")
	middlewares.NewErrorHandler(newRoute, errorHandlerUseCase)

	controllers.NewLoginControllerImpl(newRoute, loginUsecase, errorHandlerUseCase)

	newRoute.Use(middlewares.TokenAuthMiddleware())
	controllers.NewReportControllerImpl(newRoute, transUsecase)

	if err := router.Run(env.Config.ServiceHost + ":" + env.Config.AppPort); err != nil {
		log.Fatal("error starting server", err)
	}

}
