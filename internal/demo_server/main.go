package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	demo_service "github.com/gempellm/my-grpc/pkg/demo_service"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8081, "The server port")
)

type server struct {
	demo_service.UnimplementedMyGreeterServer
}

func (s *server) GreetIncomer(ctx context.Context, in *demo_service.GreetRequest) (*demo_service.GreetReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &demo_service.GreetReply{Greet: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	demo_service.RegisterMyGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}
