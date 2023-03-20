package profile

import (
	"context"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/grpc/client"
	"github.com/stretchr/testify/assert"
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
	assert.NoError(t, err)

	_, err = p.Get(context.Background(), &api.ProfileUserID{UserID: res.User.ID})
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	time.Sleep(time.Second)
	conn, p := newConnAndService(t)
	app := newAppAuthService(conn)
	defer conn.Close()

	res, err := app.Login(context.Background(), &api.AppRequest{
		Email:    "test@gmail.com",
		Password: "test",
	})
	assert.NoError(t, err)

	_, err = p.Update(context.Background(), &api.ProfileRequest{
		UserID:    res.User.ID,
		FirstName: "testName",
		LastName:  "testLastName",
	})
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	time.Sleep(time.Second)
	conn, p := newConnAndService(t)
	app := newAppAuthService(conn)
	defer conn.Close()

	res, err := app.Login(context.Background(), &api.AppRequest{
		Email:    "test@gmail.com",
		Password: "test",
	})
	assert.NoError(t, err)

	_, err = p.Delete(context.Background(), &api.ProfileUserID{UserID: res.User.ID})
	assert.NoError(t, err)
}

func newConnAndService(t *testing.T) (*grpc.ClientConn, api.ProfileServiceClient) {
	conn, err := client.NewConn("localhost:9000")
	assert.NoError(t, err)
	p := api.NewProfileServiceClient(conn)
	return conn, p
}

func newAppAuthService(conn *grpc.ClientConn) api.AuthAppServiceClient {
	return api.NewAuthAppServiceClient(conn)
}
