package auth

import (
	"context"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/grpc/client"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	conn, app := newConnAndService(t)
	defer conn.Close()
	_, err := app.Register(context.Background(), &api.AppRequest{
		Email:    "test@gmail.com",
		Password: "test",
	})
	assert.NoError(t, err)
}

func TestBadRegister(t *testing.T) {
	conn, app := newConnAndService(t)
	defer conn.Close()
	_, err := app.Register(context.Background(), &api.AppRequest{
		Email:    "test@gmail.com",
		Password: "test",
	})
	assert.Error(t, err)
}

func TestLogin(t *testing.T) {
	time.Sleep(time.Second)
	conn, app := newConnAndService(t)
	defer conn.Close()
	_, err := app.Login(context.Background(), &api.AppRequest{
		Email:    "test@gmail.com",
		Password: "test",
	})
	assert.NoError(t, err)
}

func TestBadLogin(t *testing.T) {
	conn, app := newConnAndService(t)
	defer conn.Close()
	_, err := app.Login(context.Background(), &api.AppRequest{
		Email:    "testTest@gmail.com",
		Password: "testTest",
	})
	assert.Error(t, err)
}

func newConnAndService(t *testing.T) (*grpc.ClientConn, api.AuthAppServiceClient) {
	conn, err := client.NewConn("localhost:9000")
	assert.NoError(t, err)
	app := api.NewAuthAppServiceClient(conn)
	return conn, app
}
