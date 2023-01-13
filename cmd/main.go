package main

import (
	"github.com/cherryReptile/WS-AUTH/bootstrap"
	"github.com/cherryReptile/WS-AUTH/grpc/client"
	"github.com/cherryReptile/WS-AUTH/rest/controllers"
	"github.com/cherryReptile/WS-AUTH/rest/middlewares"
	"github.com/cherryReptile/WS-AUTH/rest/router"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	fatalChan := make(chan error, 1)
	gRPCFatal := make(chan error, 1)

	app := new(bootstrap.GinApp)
	app.Init()

	grpcApp := new(bootstrap.RPCApp)
	grpcApp.Init()

	conn, errConn := client.NewConn(":9000")
	if errConn != nil {
		logrus.Fatalf("failed to create connection to gRPC with error:%v", errConn)
	}
	grpcClients := new(client.ServiceClients)
	grpcClients.Init(conn)

	app.Router.Use(gin.Logger())
	df := app.Router.Group("/")

	ghc := new(controllers.GithubController)
	ghc.Init(grpcClients.GitHub)
	router.NewGitHubRouter(df, ghc)

	ac := new(controllers.AppController)
	ac.Init(grpcClients.App)
	router.NewAppRouter(df, ac)

	gc := new(controllers.GoogleController)
	gc.Init(grpcClients.Google)
	router.NewGoogleRouter(df, gc)

	hc := new(controllers.HomeController)
	h := app.Router.Group("/home")
	h.Use(middlewares.CheckAuthHeader()).Use(middlewares.CheckUserAndToken(grpcClients.CheckAuth))

	//example route for demonstrate how it will look in main app
	h.GET("/test", hc.Test)

	acc := h.Group("/account")
	acc.POST("/github", ghc.AddAccount)
	acc.POST("/google", gc.AddAccount)
	acc.POST("/app", ac.AddAccount)

	lc := new(controllers.LogoutController)
	lc.Init(grpcClients.Logout)
	h.GET("/logout", lc.Logout)

	pc := new(controllers.ProfileController)
	pc.Init(grpcClients.Profile)
	router.NewProfileRouter(h, pc)

	go app.Run("2000", fatalChan)
	go grpcApp.Server.ListenAndServe("9000", gRPCFatal)

	errG := <-gRPCFatal
	if errG != nil {
		grpcApp.Server.Close()
		grpcApp.DB.Close()
		logrus.Warning(errG)
	}

	err := <-fatalChan
	if err != nil {
		app.Close()
	}
}
