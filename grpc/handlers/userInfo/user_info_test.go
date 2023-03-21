package userInfo

import (
	"context"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/grpc/client"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestGetAllUsers(t *testing.T) {
	_, u := newConnAndService(t)
	_, err := u.GetAllUsersWithSortAndFilter(context.Background(), &api.GetUsersRequest{
		OrderBy: "desc",
		Field:   "created_at",
		Filter:  map[string]string{"login": "tEsT"},
	})
	time.Sleep(time.Second)
	require.NoError(t, err)
}

func TestGetAllUsersWithUndefinedColumn(t *testing.T) {
	_, u := newConnAndService(t)
	_, err := u.GetAllUsersWithSortAndFilter(context.Background(), &api.GetUsersRequest{
		Filter: map[string]string{"test": "test"},
	})
	time.Sleep(time.Second)
	require.Error(t, err)
}

func newConnAndService(t *testing.T) (*grpc.ClientConn, api.UserInfoServiceClient) {
	conn, err := client.NewConn("localhost:9000")
	require.NoError(t, err)
	u := api.NewUserInfoServiceClient(conn)
	return conn, u
}
