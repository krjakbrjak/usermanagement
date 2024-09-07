package main

import (
	"fmt"
	"log"
	"net"

	"github.com/krjakbrjak/usermanagement/agent"
	"github.com/krjakbrjak/usermanagement/generated"
	"google.golang.org/grpc"
)

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    generated.RegisterPasswordPolicyServiceServer(s, &agent.Agent{})
    fmt.Println("Server is running on port 50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
