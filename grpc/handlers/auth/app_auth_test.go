package auth

import (
	"context"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/grpc/client"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	conn, err := client.NewConn("localhost:9000")
	assert.NoError(t, err)
	defer conn.Close()

	app := api.NewAuthAppServiceClient(conn)
	_, err = app.Register(context.Background(), &api.AppRequest{
		Email:    "test@gmail.com",
		Password: "test",
	})
	assert.NoError(t, err)
}

func TestBadRegister(t *testing.T) {
	conn, err := client.NewConn("localhost:9000")
	assert.NoError(t, err)
	defer conn.Close()

	app := api.NewAuthAppServiceClient(conn)
	_, err = app.Register(context.Background(), &api.AppRequest{
		Email:    "test@gmail.com",
		Password: "test",
	})
	assert.Error(t, err)
}

func TestLogin(t *testing.T) {
	time.Sleep(time.Second)
	conn, err := client.NewConn("localhost:9000")
	assert.NoError(t, err)
	defer conn.Close()

	app := api.NewAuthAppServiceClient(conn)
	_, err = app.Login(context.Background(), &api.AppRequest{
		Email:    "test@gmail.com",
		Password: "test",
	})
	assert.NoError(t, err)
}

func TestBadLogin(t *testing.T) {
	conn, err := client.NewConn("localhost:9000")
	assert.NoError(t, err)
	defer conn.Close()

	app := api.NewAuthAppServiceClient(conn)
	_, err = app.Login(context.Background(), &api.AppRequest{
		Email:    "testTest@gmail.com",
		Password: "testTest",
	})
	assert.Error(t, err)
}
