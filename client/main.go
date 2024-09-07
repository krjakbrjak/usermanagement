package main

import (
	"context"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/krjakbrjak/usermanagement/generated"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a new client
	c := generated.NewPasswordPolicyServiceClient(conn)

	// Call the GetPasswordPolicy method
	resp, err := c.GetPasswordPolicy(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatalf("could not get password policy: %v", err)
	}

	fmt.Println("Password Policy:")
	fmt.Printf("Min Length: %d\n", resp.GetMinLength())
	fmt.Printf("Max Age (Days): %d\n", resp.GetMaxDays())
}
