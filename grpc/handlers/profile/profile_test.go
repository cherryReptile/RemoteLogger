package profile

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

func TestGet(t *testing.T) {
	conn, p := newConnAndService(t)
	app := newAppAuthService(conn)
	defer conn.Close()

	res, err := app.Register(context.Background(), &api.AppRequest{
		Email:    "test@gmail.com",
		Password: "test",
	})
	require.NoError(t, err)
	assert.NotNil(t, res.User)
	assert.NotEqual(t, res.JWTToken, "")

	_, err = p.Get(context.Background(), &api.ProfileUserID{UserID: res.User.ID})
	require.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	time.Sleep(time.Second)
	conn, p := newConnAndService(t)
	app := newAppAuthService(conn)
	defer conn.Close()

	id := login(app, t)

	r, err := p.Update(context.Background(), &api.ProfileRequest{
		UserID:    id,
		FirstName: "testName",
		LastName:  "testLastName",
	})
	require.NoError(t, err)
	assert.NotEqual(t, r.FirstName, "")
	assert.NotEqual(t, r.LastName, "")
}

func TestDelete(t *testing.T) {
	time.Sleep(time.Second)
	conn, p := newConnAndService(t)
	app := newAppAuthService(conn)
	defer conn.Close()

	id := login(app, t)

	_, err := p.Delete(context.Background(), &api.ProfileUserID{UserID: id})
	require.NoError(t, err)
}

func newConnAndService(t *testing.T) (*grpc.ClientConn, api.ProfileServiceClient) {
	conn, err := client.NewConn("localhost:9000")
	require.NoError(t, err)
	p := api.NewProfileServiceClient(conn)
	return conn, p
}

func newAppAuthService(conn *grpc.ClientConn) api.AuthAppServiceClient {
	return api.NewAuthAppServiceClient(conn)
}

func login(app api.AuthAppServiceClient, t *testing.T) string {
	res, err := app.Login(context.Background(), &api.AppRequest{
		Email:    "test@gmail.com",
		Password: "test",
	})
	require.NoError(t, err)
	assert.NotNil(t, res.User)
	assert.NotEqual(t, res.JWTToken, "")
	return res.User.ID
}
