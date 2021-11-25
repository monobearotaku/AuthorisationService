package main

import (
	"context"
	"fmt"
	pb "github.com/ash0tych/gRPC_MusicService/proto/compiled/authorization"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:20100", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reply, err := client.IdentifyUser(ctx, &pb.UserData{Login: "ash0tykit10@gmail.com", Password: "12345678"})
	if err != nil {
		log.Fatal(err)
	}
	if reply.Ok {
		fmt.Println("Access is allowed")
	} else {
		fmt.Printf("Error: %v \n", reply.Err.Err)
		fmt.Printf("Error ID: %v \n", reply.Err.Id)
	}
}
