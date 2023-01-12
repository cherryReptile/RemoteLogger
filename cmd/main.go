package main

import (
	gbase "github.com/cherryReptile/WS-AUTH/grpc/base"
	"github.com/cherryReptile/WS-AUTH/grpc/client"
	"github.com/cherryReptile/WS-AUTH/grpc/handlers/auth"
	"github.com/cherryReptile/WS-AUTH/grpc/handlers/profile"
	"github.com/cherryReptile/WS-AUTH/grpc/server"
	"github.com/cherryReptile/WS-AUTH/internal/base"
	"github.com/cherryReptile/WS-AUTH/internal/controllers"
	"github.com/cherryReptile/WS-AUTH/internal/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	fatalChan := make(chan error, 1)
	gRPCFatal := make(chan error, 1)

	app := new(base.App)
	app.Init()

	db := new(gbase.Database)
	db.Init()
	grpcServer := server.NewServer(server.Services{
		App:       auth.NewAppAuthService(db.Conn),
		GitHub:    auth.NewGitHubAuthService(db.Conn),
		Google:    auth.NewGoogleAuthService(db.Conn),
		Telegram:  auth.NewTelegramAuthService(db.Conn),
		CheckAuth: auth.NewCheckAuthService(db.Conn),
		Logout:    auth.NewLogoutAuthService(db.Conn),
		Profile:   profile.NewUserProfileService(db.Conn),
	})

	conn, errConn := client.NewConn(":9000")
	if errConn != nil {
		logrus.Fatalf("failed to create connection to gRPC with error:%v", errConn)
	}
	grpcClients := new(client.ServiceClients)
	grpcClients.Init(conn)

	app.Router.Use(gin.Logger())

	githubController := new(controllers.GithubAuthController)
	githubController.Init(grpcClients.GitHub)

	authGit := app.Router.Group("/github")
	authGit.GET("/", githubController.RedirectToGoogle)
	authGit.GET("/token", githubController.GetAccessToken)
	authGit.POST("/login", githubController.Login)

	appController := new(controllers.AppAuthController)
	appController.Init(grpcClients.App)

	authApp := app.Router.Group("/app")
	authApp.POST("/register", appController.Register)
	authApp.POST("/login", appController.Login)

	googleController := new(controllers.GoogleAuthController)
	googleController.Init(grpcClients.Google)
	authGo := app.Router.Group("/google")
	authGo.GET("/", googleController.RedirectToGoogle)
	authGo.GET("/token", googleController.GetAccessToken)
	authGo.POST("/login", googleController.Login)

	tgAuthC := new(controllers.TelegramAuthController)
	tgAuthC.Init(grpcClients.Telegram)
	authTg := app.Router.Group("/telegram")
	authTg.GET("/login")

	homeC := new(controllers.HomeController)
	home := app.Router.Group("/home")
	home.Use(middlewares.CheckAuthHeader()).Use(middlewares.CheckUserAndToken(grpcClients.CheckAuth))

	//example route for demonstrate how it will look in main app
	home.GET("/test", homeC.Test)
	accounts := home.Group("/account")
	accounts.POST("/github", githubController.AddAccount)
	accounts.POST("/google", googleController.AddAccount)
	accounts.POST("/app", appController.AddAccount)

	logoutC := new(controllers.LogoutController)
	logoutC.Init(grpcClients.Logout)
	home.GET("/logout", logoutC.Logout)

	profileC := new(controllers.ProfileController)
	profileC.Init(grpcClients.Profile)
	profile := home.Group("/profile")
	profile.POST("/create", profileC.Create)
	profile.GET("/get", profileC.Get)
	profile.PATCH("/update", profileC.Update)
	profile.DELETE("/delete", profileC.Delete)

	go app.Run("2000", fatalChan)
	go grpcServer.ListenAndServe("9000", gRPCFatal)

	errG := <-gRPCFatal
	if errG != nil {
		grpcServer.Close()
		db.Close()
		logrus.Warning(errG)
	}

	err := <-fatalChan
	if err != nil {
		app.Close()
	}
}
