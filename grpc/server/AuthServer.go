package server

import (
	"fmt"
	"github.com/pavel-one/GoStarter/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Services struct {
	App       api.AuthAppServiceServer
	GitHub    api.AuthGithubServiceServer
	Google    api.AuthGoogleServiceServer
	Telegram  api.AuthTelegramServiceServer
	CheckAuth api.CheckAuthServiceServer
	Logout    api.LogoutServiceServer
}

type Server struct {
	Services
	srv *grpc.Server
}

func NewServer(services Services) *Server {
	return &Server{
		Services: services,
		srv:      grpc.NewServer(),
	}

}

func (s *Server) ListenAndServe(port string, errCh chan error) {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		errCh <- err
	}

	api.RegisterAuthAppServiceServer(s.srv, s.Services.App)
	api.RegisterAuthGithubServiceServer(s.srv, s.Services.GitHub)
	api.RegisterAuthGoogleServiceServer(s.srv, s.Services.Google)
	api.RegisterAuthTelegramServiceServer(s.srv, s.Services.Telegram)
	api.RegisterLogoutServiceServer(s.srv, s.Services.Logout)
	api.RegisterCheckAuthServiceServer(s.srv, s.Services.CheckAuth)
	log.Println("[DEBUG] Running gRCP server on port " + port)
	if err = s.srv.Serve(l); err != nil {
		errCh <- err
	}
}

func (s *Server) Close() {
	fmt.Println("[DEBUG] grpc auth server close")
	s.srv.Stop()
}
