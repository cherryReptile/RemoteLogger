package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pavel-one/GoStarter/internal/base"
	"github.com/pavel-one/GoStarter/internal/controllers"
	"github.com/pavel-one/GoStarter/internal/middlewars"
	"log"
	"os"
)

func main() {
	fatalChan := make(chan error, 1)

	app := new(base.App)
	app.Init()

	githubController := new(controllers.GithubAuthController)
	githubController.Init()
	app.Router.Use(gin.Logger())

	auth := app.Router.Group("/auth")
	authGit := auth.Group("/github")
	authGit.GET("/", githubController.RedirectForAuth)
	authGit.GET("/login", githubController.Login)

	testController := new(controllers.TestController)
	testController.Init()
	authorized := app.Router.Group("/home")
	authorized.Use(middlewars.CheckAuthHeader()).Use(middlewars.CheckUserAndToken())
	authorized.GET("/test", testController.Test)

	appAuthController := new(controllers.AppAuthController)
	authApp := auth.Group("/app")
	authApp.POST("/register", appAuthController.Register)

	go app.Run("80", fatalChan)

	err := <-fatalChan
	if err != nil {
		app.Close()
		log.Printf("[FATAL] %v", err)
		os.Exit(1)
	}
}
