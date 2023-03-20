package client

import (
	"github.com/cherryReptile/WS-AUTH/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClients struct {
	Conn     *grpc.ClientConn
	App      api.AuthAppServiceClient
	GitHub   api.AuthGithubServiceClient
	Google   api.AuthGoogleServiceClient
	Profile  api.ProfileServiceClient
	GetUser  api.GetUserServiceClient
	JwtToken api.JWTTokenServiceClient
	UserInfo api.UserInfoServiceClient
}

func (s *ServiceClients) Init(conn *grpc.ClientConn) {
	s.Conn = conn
	s.App = api.NewAuthAppServiceClient(s.Conn)
	s.GitHub = api.NewAuthGithubServiceClient(s.Conn)
	s.Google = api.NewAuthGoogleServiceClient(s.Conn)
	s.Profile = api.NewProfileServiceClient(s.Conn)
	s.GetUser = api.NewGetUserServiceClient(s.Conn)
	s.JwtToken = api.NewJWTTokenServiceClient(s.Conn)
	s.UserInfo = api.NewUserInfoServiceClient(s.Conn)
}

func NewConn(target string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return conn, nil
}
