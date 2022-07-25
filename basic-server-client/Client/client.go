package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	chat "github.com/jays2812/gRPC/basic-server-client/chat"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could Not Connect: %s", err)
	}

	defer conn.Close()
	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello says the Client!",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from Server: %s", response.Body)
}
