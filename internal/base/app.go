package base

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

//TODO: this look like proxy and in the future will be renamed

type App struct {
	Router *gin.Engine
	Server *http.Server
}

func (a *App) Init() {
	a.Router = gin.New()
}

func (a *App) Run(port string, chErr chan error) {
	a.Server = &http.Server{
		Handler:      a.Router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	defer a.Server.Close()

	logrus.Printf("Running server on port %s", port)

	if err := a.Server.ListenAndServe(); err != nil {
		chErr <- errors.New(fmt.Sprintf("Error server: %s", err.Error()))
	}
}

func (a *App) Close() {
	logrus.Printf("Close server on address %s", a.Server.Addr)
	if err := a.Server.Close(); err != nil {
		logrus.Fatalf("Unable to close server: %v", err)
	}
}
