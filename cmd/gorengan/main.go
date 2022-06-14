package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/musobarlab/gorengan/config"
	"github.com/musobarlab/gorengan/internal/server"
)

func main() {
	// call config.Load() before start up
	err := config.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s, err := server.NewHTTPServer(config.Port)
	if err != nil {
		fmt.Println("port not specified")
		os.Exit(1)
	}

	go s.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	s.Exit()
}
