package auth

import (
	"context"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/grpc/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	conn, app := newConnAndAuthAppService(t)
	defer conn.Close()
	res, err := app.Register(context.Background(), &api.AppRequest{
		Email:    "test@gmail.com",
		Password: "test",
	})
	require.NoError(t, err)
	assert.NotNil(t, res.User)
	assert.NotEqual(t, res.JWTToken, "")
}

func TestBadRegister(t *testing.T) {
	conn, app := newConnAndAuthAppService(t)
	defer conn.Close()
	_, err := app.Register(context.Background(), &api.AppRequest{
		Email:    "test@gmail.com",
		Password: "test",
	})
	require.Error(t, err)
}

func TestLogin(t *testing.T) {
	time.Sleep(time.Second)
	conn, app := newConnAndAuthAppService(t)
	defer conn.Close()
	res, err := app.Login(context.Background(), &api.AppRequest{
		Email:    "test@gmail.com",
		Password: "test",
	})
	require.NoError(t, err)
	assert.NotNil(t, res.User)
	assert.NotEqual(t, res.JWTToken, "")
}

func TestBadLogin(t *testing.T) {
	conn, app := newConnAndAuthAppService(t)
	defer conn.Close()
	_, err := app.Login(context.Background(), &api.AppRequest{
		Email:    "testTest@gmail.com",
		Password: "testTest",
	})
	require.Error(t, err)
}

func newConnAndAuthAppService(t *testing.T) (*grpc.ClientConn, api.AuthAppServiceClient) {
	conn, err := client.NewConn("localhost:9000")
	require.NoError(t, err)
	app := api.NewAuthAppServiceClient(conn)
	return conn, app
}
