package main

import (
	"UserService/internal/user"
	pb "UserService/proto/api/v1/user"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	address := ":9090"

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterUserServiceServer(server, &user.UserServiceImpl{})

	log.Printf("gRPC server running on %s\n", address)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("serve error: %v", err)
	}
}
