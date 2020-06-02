package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vbansal/login_service/config"
	"github.com/vbansal/login_service/controller"

	_ "github.com/vbansal/login_service/docs"
)

// @title Login API
// @version 1.0
// @description Swagger API for Login Service.
// @termsOfService http://swagger.io/terms/

// @contact.name Login API Support
// @contact.email vbansal13@gmail.com

// @BasePath /v1
func main() {
	router := gin.Default()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userCtrl := controller.NewUserController()
	v1 := router.Group("/v1")
	{
		v1.POST("/me", userCtrl.SignupUserHandler)
		v1.POST("/me/login", userCtrl.LoginUserHandler)
		v1.GET("/me", userCtrl.ProfileUserHandler)
	}
	router.Run(fmt.Sprintf(":%v", config.GetInstance().ServerPort))

}
