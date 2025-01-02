package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"

	pb "github.com/levensspel/simple-grpc-go/proto/user"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) RegisterUser(_ context.Context, in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	log.Printf("Received Name: %v", in.GetName())
	log.Printf("Received Username: %v", in.GetUsername())
	log.Printf("Received Email: %v", in.GetEmail())
	for k, v := range in.GetSettings() {
		log.Printf("Received Settings->%s: %s", k, v)
	}

	return &pb.RegisterUserResponse{
		UserId:      rand.Uint32(),
		EmailStatus: pb.EmailStatus_EMAIL_UNVERIFIED,
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
