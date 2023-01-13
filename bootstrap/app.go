package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

//TODO: this look like proxy and in the future will be renamed

type GinApp struct {
	Router *gin.Engine
	Server *http.Server
}

func (a *GinApp) Init() {
	a.Router = gin.New()
}

func (a *GinApp) Run(port string, chErr chan error) {
	a.Server = &http.Server{
		Handler:      a.Router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	defer a.Server.Close()

	logrus.Printf("Running server on port %s", port)

	if err := a.Server.ListenAndServe(); err != nil {
		chErr <- fmt.Errorf("error server: %v", err)
	}
}

func (a *GinApp) Close() {
	logrus.Printf("Close server on address %s", a.Server.Addr)
	if err := a.Server.Close(); err != nil {
		logrus.Fatalf("Unable to close server: %v", err)
	}
}
