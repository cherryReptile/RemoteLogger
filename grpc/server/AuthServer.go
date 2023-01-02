package server

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/pavel-one/GoStarter/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"net"
)

type Services struct {
	App       api.AuthAppServiceServer
	GitHub    api.AuthGithubServiceServer
	Google    api.AuthGoogleServiceServer
	Telegram  api.AuthTelegramServiceServer
	CheckAuth api.CheckAuthServiceServer
	Logout    api.LogoutServiceServer
	Profile   api.ProfileServiceServer
}

type Logger struct {
	Logrus     *logrus.Logger
	CustomFunc grpc_logrus.CodeToLevel
}

type Server struct {
	Services
	srv *grpc.Server
}

func NewServer(services Services) *Server {
	logger := new(Logger)
	logger.Logrus = logrus.New()
	logrusEntry := logrus.NewEntry(logger.Logrus)
	logger.CustomFunc = func(code codes.Code) logrus.Level {
		if code == codes.OK {
			return logrus.InfoLevel
		}
		return logrus.ErrorLevel
	}

	logrusOpts := []grpc_logrus.Option{
		grpc_logrus.WithLevels(logger.CustomFunc),
	}

	grpc_logrus.ReplaceGrpcLogger(logrusEntry)

	opt := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			grpc_logrus.UnaryServerInterceptor(logrusEntry, logrusOpts...),
		),
	}

	return &Server{
		Services: services,
		srv:      grpc.NewServer(opt...),
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
	api.RegisterCheckAuthServiceServer(s.srv, s.Services.CheckAuth)
	api.RegisterLogoutServiceServer(s.srv, s.Services.Logout)
	api.RegisterProfileServiceServer(s.srv, s.Services.Profile)
	logrus.Info("[DEBUG] Running gRPC server on port " + port)
	if err = s.srv.Serve(l); err != nil {
		errCh <- err
	}
}

func (s *Server) Close() {
	logrus.Info("[DEBUG] grpc auth server close")
	s.srv.Stop()
}
