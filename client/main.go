package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/levensspel/simple-grpc-go/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.RegisterUser(ctx, &pb.RegisterUserRequest{
		Name:     "Your Name",
		Username: "your_username",
		Email:    "example@email.com",
		Password: "SuperSecretPassword",
		Settings: map[string]string{
			"Theme":    "dark",
			"Language": "id",
		},
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Returned User ID: %d", r.GetUserId())
	log.Printf("Returned Email Status: %d", r.GetEmailStatus())
}
