package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pavel-one/GoStarter/internal/base"
	"github.com/pavel-one/GoStarter/internal/controllers"
	"log"
	"os"
)

func main() {
	fatalChan := make(chan error, 1)

	app := new(base.App)
	app.Init()

	githubController := new(controllers.GithubAuthController)
	githubController.Init(app.DB)
	app.Router.Use(gin.Logger())

	app.Router.GET("/auth", githubController.RedirectForAuth)
	app.Router.GET("/auth/github/login", githubController.Login)
	go app.Run("80", fatalChan)

	err := <-fatalChan
	if err != nil {
		app.Close()
		log.Printf("[FATAL] %v", err)
		os.Exit(1)
	}
}
