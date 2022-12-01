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

	githubC := new(controllers.GithubAuthController)
	githubC.Init()
	app.Router.Use(gin.Logger())

	auth := app.Router.Group("/auth")
	authGit := auth.Group("/github")
	authGit.GET("/", githubC.RedirectForAuth)
	authGit.GET("/login", githubC.Login)

	appAuthC := new(controllers.AppAuthController)
	authApp := auth.Group("/app")
	authApp.POST("/register", appAuthC.Register)
	authApp.POST("/login", appAuthC.Login)

	googleAuthC := new(controllers.GoogleAuthController)
	googleAuthC.Init()
	authGo := auth.Group("/google")
	authGo.GET("/", googleAuthC.RedirectForAuth)
	authGo.GET("/login", googleAuthC.Login)

	testC := new(controllers.TestController)
	testC.Init()
	home := app.Router.Group("/home")
	home.Use(middlewars.CheckAuthHeader()).Use(middlewars.CheckUserAndToken())
	home.GET("/test", testC.Test)
	home.GET("/app/logout", appAuthC.Logout)
	home.GET("/github/logout", githubC.Logout)
	home.GET("/google/logout", googleAuthC.Logout)

	go app.Run("80", fatalChan)

	err := <-fatalChan
	if err != nil {
		app.Close()
		log.Printf("[FATAL] %v", err)
		os.Exit(1)
	}
}
