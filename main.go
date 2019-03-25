package main

import (
	"fmt"
	"os"

	"github.com/musobarlab/gorengan/config"
	"github.com/musobarlab/gorengan/server"
)

func main() {
	// call config.Load() before start up
	err := config.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s, err := server.NewEchoServer(config.Port)
	if err != nil {
		fmt.Println("port not specified")
		os.Exit(1)
	}

	s.Run()
}
