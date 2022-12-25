package main

import (
	"github.com/gin-gonic/gin"
	gbase "github.com/pavel-one/GoStarter/grpc/base"
	"github.com/pavel-one/GoStarter/grpc/client"
	"github.com/pavel-one/GoStarter/grpc/handlers"
	"github.com/pavel-one/GoStarter/grpc/server"
	"github.com/pavel-one/GoStarter/internal/base"
	"github.com/pavel-one/GoStarter/internal/controllers"
	"github.com/pavel-one/GoStarter/internal/middlewares"
	"log"
	"os"
)

func main() {
	fatalChan := make(chan error, 1)
	gRPCFatal := make(chan error, 1)

	app := new(base.App)
	app.Init()

	db := new(gbase.Database)
	db.Init()
	grpcServer := server.NewServer(server.Services{
		App:       handlers.NewAppAuthService(db.Conn),
		GitHub:    handlers.NewGitHubAuthService(db.Conn),
		Google:    handlers.NewGoogleAuthService(db.Conn),
		Telegram:  handlers.NewTelegramAuthService(db.Conn),
		CheckAuth: handlers.NewCheckAuthService(db.Conn),
		Logout:    handlers.NewLogoutAuthService(db.Conn),
	})

	conn, errConn := client.NewConn()
	if errConn != nil {
		log.Printf("[FATAL] %v", errConn)
		os.Exit(1)
	}
	grpcClients := new(client.ServiceClients)
	grpcClients.Init(conn)

	app.Router.Use(gin.Logger())

	githubC := new(controllers.GithubAuthController)
	githubC.Init(grpcClients.GitHub)

	auth := app.Router.Group("/auth")
	authGit := auth.Group("/github")
	authGit.GET("/", githubC.RedirectForAuth)
	authGit.GET("/token", githubC.GetAccessToken)
	authGit.POST("/login", githubC.Login)

	appAuthC := new(controllers.AppAuthController)
	appAuthC.Init(grpcClients.App)

	authApp := auth.Group("/app")
	authApp.POST("/register", appAuthC.Register)
	authApp.POST("/login", appAuthC.Login)

	googleAuthC := new(controllers.GoogleAuthController)
	googleAuthC.Init(grpcClients.Google)
	authGo := auth.Group("/google")
	authGo.GET("/", googleAuthC.RedirectForAuth)
	authGo.GET("/token", googleAuthC.GetAccessToken)
	authGo.POST("/login", googleAuthC.Login)

	tgAuthC := new(controllers.TelegramAuthController)
	tgAuthC.Init(grpcClients.Telegram)
	authTg := auth.Group("/telegram")
	authTg.GET("/login")

	homeC := new(controllers.HomeController)
	homeC.Init()
	home := app.Router.Group("/home")
	home.Use(middlewares.CheckAuthHeader()).Use(middlewares.CheckUserAndToken(grpcClients.CheckAuth))

	home.GET("/test", homeC.Test)
	accounts := home.Group("/account")
	accounts.POST("/github", githubC.AddAccount)
	accounts.POST("/google", googleAuthC.AddAccount)

	logoutC := new(controllers.LogoutController)
	logoutC.Init(grpcClients.Logout)
	home.GET("/logout", logoutC.Logout)

	go app.Run("80", fatalChan)
	go grpcServer.ListenAndServe("9000", gRPCFatal)

	errG := <-gRPCFatal
	if errG != nil {
		grpcServer.Close()
		db.Close()
		log.Printf("[FATAL] %v", errG)
	}

	err := <-fatalChan
	if err != nil {
		app.Close()
		log.Printf("[FATAL] %v", err)
		os.Exit(1)
	}
}
