package main

import (
	pb "github.com/ash0tych/gRPC_MusicService/proto/compiled/authorization"
	"github.com/ash0tych/gRPC_MusicService/service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":20100")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := server.Server{}
	err = grpcServer.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer grpcServer.Close()
	log.Printf("start listening for emails at port %s", ":20100")

	rpcserv := grpc.NewServer()
	pb.RegisterUserServiceServer(rpcserv, &grpcServer)
	reflection.Register(rpcserv)

	err = rpcserv.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}
