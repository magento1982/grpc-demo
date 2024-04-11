package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/magento1982/grpc-demo/api/vs"
)

func callSayHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional Streaming started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	wait := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while receiving %v", err)
			}
			log.Print(message)
		}

		close(wait)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("could not send request with name %v", name)
		}
		log.Printf("Sent the request with name: %v", name)
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-wait
	log.Printf("Bidirectional Streaming finished")
}
