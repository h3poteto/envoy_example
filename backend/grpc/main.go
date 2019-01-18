package main

import (
	"log"
	"net"
	"os"

	"golang.org/x/net/context"

	pb "github.com/h3poteto/envoy_example/lib"
	"google.golang.org/grpc"
)

type hello struct{}

func (h *hello) Say(c context.Context, p *pb.RequestType) (*pb.ResponseType, error) {
	log.Println(p.Msg)
	res := &pb.ResponseType{
		Msg: "World",
	}
	return res, nil
}

func main() {
	log.Println("Server starting...")
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("faild to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterHelloServer(server, new(hello))
	log.Println("Listen :", port)
	server.Serve(lis)
}
