package client

import (
	"context"
	"github.com/pavel-one/GoStarter/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Register(request *api.AppRequest) (*api.AppResponse, error) {
	conn, err := newConn()
	if err != nil {
		return nil, err
	}
	c := api.NewAuthServiceClient(conn)
	res, err := c.Register(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func Login(request *api.AppRequest) (*api.AppResponse, error) {
	conn, err := newConn()
	if err != nil {
		return nil, err
	}
	c := api.NewAuthServiceClient(conn)
	res, err := c.Login(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func newConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return conn, nil
}
