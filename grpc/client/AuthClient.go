package client

import (
	"context"
	"github.com/pavel-one/GoStarter/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func Register(request *api.RegisterRequest) (*api.RegisteredResponse, error) {
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	c := api.NewAuthServiceClient(conn)
	res, err := c.Register(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
