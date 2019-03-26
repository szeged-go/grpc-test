package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/szeged-go/grpc-test/registration"
)

type Server struct{}

func (s *Server) Register(cxt context.Context, req *registration.Request) (*registration.Response, error) {
	log.Print(*req)
	return &registration.Response{
		Status: "OK",
	}, nil
}

func main() {
	port := flag.Int("port", 0, "The port")

	flag.Parse()

	fmt.Println("Port: ", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	registration.RegisterRegistrationServer(grpcServer, &Server{})
	grpcServer.Serve(lis)
}
