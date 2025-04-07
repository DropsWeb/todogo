package services

import (
	"crypto/tls"
	"net/http"

	"github.com/DropsWeb/todogo/internal/database/postgres"
	"github.com/DropsWeb/todogo/pkg/swagger/ops"
)

type Server struct {
	ConnDb *postgres.PgxConnect
}

func New() *Server {
	return &Server{}
}

func (*Server) ConfigureFlags(api *ops.TodoAPIAPI)                  {}
func (*Server) ConfigureTLS(tlsConfig *tls.Config)                  {}
func (*Server) ConfigureServer(s *http.Server, scheme, addr string) {}
func (*Server) CustomConfigure(api *ops.TodoAPIAPI)                 {}
func (*Server) SetupMiddlewares(handler http.Handler) http.Handler {
	return handler
}
func (*Server) SetupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
