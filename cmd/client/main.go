package main

import (
	"context"
	"flag"
	"log"

	"github.com/youssefhmidi/gRPC-example/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	msg := flag.String("msg", "hey", "sending the message to gRPC server")

	flag.Parse()

	conn, err := grpc.NewClient("localhost:54004", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("couldn't start client , %s", err)
		return
	}

	client := services.NewMsgServerClient(conn)

	other := make(map[string]string)

	other["date"] = "today"
	other["sent"] = "today"

	sendMsg := services.ObjMessage{
		ID: 1,
		Msg: *msg,
		Data: other,
	}
	resp, err := client.Send(context.Background(), &sendMsg)
	if err != nil {
		log.Fatalf("Server error, %s", err)
		return
	}

	log.Printf("Response : %s\n", resp.Res);
}
