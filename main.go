package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/musobarlab/gorengan/server"
)

func main() {
	portStr, ok := os.LookupEnv("PORT")
	if !ok {
		fmt.Println("port not specified")
		os.Exit(1)
	}

	port, _ := strconv.Atoi(portStr)

	s, err := server.NewEchoServer(port)
	if err != nil {
		fmt.Println("port not specified")
		os.Exit(1)
	}

	s.Run()
}
