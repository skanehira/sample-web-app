package server

import (
	"github.com/labstack/echo/v4"
	"github.com/skanehira/sample-web-app/api/app/server/handler"
)

type Server struct {
	echo *echo.Echo
}

func New(e *echo.Echo) *Server {
	return &Server{
		echo: e,
	}
}

func (s *Server) RegistHandler(h handler.Handler) *Server {
	s.echo.GET("/users/info", h.User)
	return s
}

func (s *Server) Start() error {
	return s.echo.Start(":80")
}
