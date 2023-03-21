package auth

import (
	"context"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/grpc/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetUser(t *testing.T) {
	app, u := newAppAuthWithGetUserService(t)
	res, err := app.Register(context.Background(), &api.AppRequest{
		Email:    "test@gmail.com",
		Password: "test",
	})
	require.NoError(t, err)
	assert.NotNil(t, res.User)
	assert.NotEqual(t, res.JWTToken, "")

	r, err := u.GetUser(context.Background(), &api.JWTTokenRequest{
		JWTToken: res.JWTToken,
	})

	require.NoError(t, err)
	assert.NotNil(t, r.User)
	assert.NotNil(t, r.Profile)
	assert.NotEqual(t, r.JWTToken, "")
}

func newAppAuthWithGetUserService(t *testing.T) (api.AuthAppServiceClient, api.GetUserServiceClient) {
	conn, err := client.NewConn("localhost:9000")
	require.NoError(t, err)
	app := api.NewAuthAppServiceClient(conn)
	u := api.NewGetUserServiceClient(conn)
	return app, u
}
