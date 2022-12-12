package client

import (
	"github.com/pavel-one/GoStarter/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClients struct {
	Conn      *grpc.ClientConn
	App       api.AuthAppServiceClient
	GitHub    api.AuthGithubServiceClient
	Google    api.AuthGoogleServiceClient
	Telegram  api.AuthTelegramServiceClient
	CheckAuth api.CheckAuthServiceClient
	Logout    api.LogoutServiceClient
}

func (s *ServiceClients) Init(conn *grpc.ClientConn) {
	s.Conn = conn
	s.App = api.NewAuthAppServiceClient(s.Conn)
	s.GitHub = api.NewAuthGithubServiceClient(s.Conn)
	s.Google = api.NewAuthGoogleServiceClient(s.Conn)
	s.Telegram = api.NewAuthTelegramServiceClient(s.Conn)
	s.CheckAuth = api.NewCheckAuthServiceClient(s.Conn)
	s.Logout = api.NewLogoutServiceClient(s.Conn)
}

func NewConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return conn, nil
}
