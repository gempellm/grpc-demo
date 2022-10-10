package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc/credentials/insecure"

	demo_service "github.com/gempellm/my-grpc/pkg/demo_service"

	"google.golang.org/grpc"
)

const (
	defaultName = "World"
)

var (
	addr = flag.String("addr", "localhost:8081", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := demo_service.NewMyGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GreetIncomer(ctx, &demo_service.GreetRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetGreet())
}
