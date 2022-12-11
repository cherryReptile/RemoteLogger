package client

import (
	"context"
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

func AppRegister(request *api.AppRequest) (*api.AppResponse, error) {
	conn, err := NewConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := api.NewAuthAppServiceClient(conn)
	res, err := c.Register(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func AppLogin(request *api.AppRequest) (*api.AppResponse, error) {
	conn, err := NewConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := api.NewAuthAppServiceClient(conn)
	res, err := c.Login(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GithubLogin(request *api.GitHubRequest) (*api.AppResponse, error) {
	conn, err := NewConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := api.NewAuthGithubServiceClient(conn)
	res, err := c.Login(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GoogleLogin(request *api.GoogleRequest) (*api.AppResponse, error) {
	conn, err := NewConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := api.NewAuthGoogleServiceClient(conn)
	res, err := c.Login(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func TelegramLogin(request *api.TelegramRequest) (*api.AppResponse, error) {
	conn, err := NewConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := api.NewAuthTelegramServiceClient(conn)
	res, err := c.Login(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func Logout(request *api.TokenRequest) (*api.LogoutResponse, error) {
	conn, err := NewConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := api.NewLogoutServiceClient(conn)
	res, err := c.Logout(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func CheckAuth(request *api.TokenRequest) (*api.CheckAuthResponse, error) {
	conn, err := NewConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := api.NewCheckAuthServiceClient(conn)
	res, err := c.CheckAuth(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return conn, nil
}
