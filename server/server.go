package server

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// EchoServer struct
type EchoServer struct {
	e    *echo.Echo
	port int
}

// NewEchoServer echo server constructor
func NewEchoServer(port int) (*EchoServer, error) {
	e := echo.New()
	return &EchoServer{
		e:    e,
		port: port,
	}, nil
}

// Run function
func (s *EchoServer) Run() {
	s.e.Use(middleware.Logger())
	s.e.Logger.Fatal(s.e.Start(fmt.Sprintf(":%d", s.port)))
}
