package server

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// EchoServer struct
type EchoServer struct {
	port int
}

// NewEchoServer echo server constructor
func NewEchoServer(port int) (*EchoServer, error) {
	return &EchoServer{
		port: port,
	}, nil
}

// Run function
func (s *EchoServer) Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.port)))
}
