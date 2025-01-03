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
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Register user with role "User"
	user, err := c.RegisterUser(ctx, &pb.RegisterUserRequest{
		Name:     "User Name",
		Username: "user_username",
		Email:    "user@email.com",
		Password: "$2y$10$MslQAgEjHya7lnOSpQgtfeUGaQyUDCIGAO3lmSuFRiH4XnyZBHHfi",
		Settings: map[string]string{
			"Theme":    "dark",
			"Language": "id",
		},
		Role: &pb.RegisterUserRequest_UserRole{
			UserRole: &pb.UserRole{
				AccountLevel: 1,
			},
		},
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Returned User ID: %d", user.GetUserId())
	log.Printf("Returned Email Status: %d", user.GetEmailStatus())
	log.Printf("")

	// Register user with role "Membership"
	member, err := c.RegisterUser(ctx, &pb.RegisterUserRequest{
		Name:     "Member Name",
		Username: "member_username",
		Email:    "member@email.com",
		Password: "$2y$10$g.1vwDij8rB10cS4RTUOEOmF7Nart2G.hUImUmVuED2dEk1Gn1jHe",
		Settings: map[string]string{
			"Theme":    "light",
			"Language": "en",
		},
		Role: &pb.RegisterUserRequest_MembershipRole{
			MembershipRole: &pb.MembershipRole{
				AccountLevel: 3,
				Perks: []pb.MembershipPerks{
					pb.MembershipPerks_PERKS_FREE_ADS,
					pb.MembershipPerks_PERKS_BACKGROUND_PLAY,
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Returned Member ID: %d", member.GetUserId())
	log.Printf("Returned Member's Email Status: %d", member.GetEmailStatus())

	// Register user with role "User"
	creator, err := c.RegisterUser(ctx, &pb.RegisterUserRequest{
		Name:     "Creator Name",
		Username: "creator_username",
		Email:    "creator@email.com",
		Password: "$2y$10$g.1vwDij8rB10cS4RTUOEOmF7Nart2G.hUImUmVuED2dEk1Gn1jHe",
		Settings: map[string]string{
			"Theme":    "light",
			"Language": "id",
		},
		Role: &pb.RegisterUserRequest_CreatorRole{
			CreatorRole: &pb.CreatorRole{
				Permissions: []pb.RolePermission{
					pb.RolePermission_PERMISSION_UPLOAD_HIGHER_VIDEO_QUALITY,
					pb.RolePermission_PERMISSION_MONETIZE_VIDEO,
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Returned Member ID: %d", creator.GetUserId())
	log.Printf("Returned Member's Email Status: %d", creator.GetEmailStatus())
}
