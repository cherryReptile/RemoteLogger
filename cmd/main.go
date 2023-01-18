package main

import (
	"github.com/cherryReptile/WS-AUTH/bootstrap"
	"github.com/sirupsen/logrus"
)

func main() {
	gRPCFatal := make(chan error, 1)

	grpcApp := new(bootstrap.RPCApp)
	grpcApp.Init()

	go grpcApp.Server.ListenAndServe("9000", gRPCFatal)

	errG := <-gRPCFatal
	if errG != nil {
		grpcApp.Server.Close()
		grpcApp.DB.Close()
		logrus.Warning(errG)
	}
}
