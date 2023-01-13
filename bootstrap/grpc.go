package bootstrap

import (
	"github.com/cherryReptile/WS-AUTH/grpc/handlers/auth"
	"github.com/cherryReptile/WS-AUTH/grpc/handlers/profile"
	"github.com/cherryReptile/WS-AUTH/grpc/server"
)

type RPCApp struct {
	Server *server.Server
	DB     *Database
}

func (a *RPCApp) Init() {
	a.DB = new(Database)
	a.DB.Init()
	a.Server = server.NewServer(server.Services{
		App:       auth.NewAppAuthService(a.DB.Conn),
		GitHub:    auth.NewGitHubAuthService(a.DB.Conn),
		Google:    auth.NewGoogleAuthService(a.DB.Conn),
		Telegram:  auth.NewTelegramAuthService(a.DB.Conn),
		CheckAuth: auth.NewCheckAuthService(a.DB.Conn),
		Logout:    auth.NewLogoutAuthService(a.DB.Conn),
		Profile:   profile.NewUserProfileService(a.DB.Conn),
	})
}
