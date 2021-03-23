package main

import (
	"context"
	"log"

	"github.com/phramos07/learning-go-grpc/chat"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error when dialing grpc server: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from the client",
	}

	res, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling remote method SayHello: %s", err)
	}

	log.Printf("Response from remote method SayHello: %s", res.Body)
}
