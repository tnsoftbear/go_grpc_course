package main

import (
	"context"
	"github.com/cshep4/grpc-course/module5-exercise/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
)

const grpcServerAddr = "localhost:50051"

func main() {
	var ctx = context.Background()
	// Decode and verify JWT at https://jwt.io/
	var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQxMTMwMjM1NTAsImlhdCI6MTcxNDczNTk1MCwibmFtZSI6IkNocmlzIiwicm9sZSI6ImFkbWluIiwic3ViIjoidXNlci1pZC0xMjM0In0.2KcYUbgJCGDAtzKnc5z45DsPaadhERyaasuckQ6S5io"

	conn, err := grpc.NewClient(grpcServerAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewTokenServiceClient(conn)

	//tokenT := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"ID": "some-id-1"}, nil)
	//token, err = tokenT.SignedString([]byte("super-secret-key-123"))
	//if err != nil {
	//	panic(fmt.Errorf("failed to sign JWT: %w", err))
	//}

	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(
		"authorization", token,
	))
	res, err := client.Validate(ctx, &proto.ValidateRequest{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response received: %v", res.Claims)
}
