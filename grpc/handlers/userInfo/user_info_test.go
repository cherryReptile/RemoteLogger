package userInfo

import (
	"context"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/grpc/client"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
)

func TestGetAllUsers(t *testing.T) {
	_, u := newConnAndService(t)
	_, err := u.GetAllUsersWithSortAndFilter(context.Background(), &api.GetUsersRequest{
		OrderBy: "desc",
		Field:   "created_at",
		Filter:  map[string]string{"login": "tEsT"},
	})
	assert.NoError(t, err)
}

func newConnAndService(t *testing.T) (*grpc.ClientConn, api.UserInfoServiceClient) {
	conn, err := client.NewConn("localhost:9000")
	assert.NoError(t, err)
	u := api.NewUserInfoServiceClient(conn)
	return conn, u
}
