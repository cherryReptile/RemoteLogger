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
		App:       auth.NewAppAuthService(db.Conn),
		GitHub:    auth.NewGitHubAuthService(db.Conn),
		Google:    auth.NewGoogleAuthService(db.Conn),
		Telegram:  auth.NewTelegramAuthService(db.Conn),
		CheckAuth: auth.NewCheckAuthService(db.Conn),
		Logout:    auth.NewLogoutAuthService(db.Conn),
		Profile:   profile.NewUserProfileService(db.Conn),
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
	authGit.GET("/", githubC.RedirectToGoogle)
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
	authGo.GET("/", googleAuthC.RedirectToGoogle)
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
	accounts.POST("/app", appAuthC.AddAccount)

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
