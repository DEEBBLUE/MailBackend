package main

import (
	handlers "authService/internal/Handlers"
	"net"

	"google.golang.org/grpc"
)

func main() {
	serv := grpc.NewServer()
	handlers.InitAuthServer(serv)
	
	listner,err := net.Listen("tcp",":4443")
	if err != nil {
		panic(err)
	}

	if err := serv.Serve(listner); err != nil {
		panic(err)
	}
}

