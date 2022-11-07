package main

import (
	"github.com/pavel-one/GoStarter/internal/base"
	"github.com/pavel-one/GoStarter/internal/controllers"
	"log"
	"os"
)

func main() {
	fatalChan := make(chan error, 1)

	app := new(base.App)
	app.Init()

	testController := new(controllers.TestController)

	app.Router.GET("/", testController.Test)
	go app.Run("80", fatalChan)

	err := <-fatalChan
	if err != nil {
		app.Close()
		log.Printf("[FATAL] %v", err)
		os.Exit(1)
	}
}
