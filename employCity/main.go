package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "employCity/api/grpc_memcached"

	"google.golang.org/grpc"
)

type Storage interface {
	Get(context.Context, *pb.DeleteGetRequest) (*pb.Reply, error)
	Set(context.Context, *pb.SetRequest) (*pb.Reply, error)
	Delete(context.Context, *pb.DeleteGetRequest) (*pb.Reply, error)
}
type server struct {
	pb.UnimplementedMyGRPCServer
}

func (s *server) Get(context.Context, *pb.DeleteGetRequest) (*pb.Reply, error) {
	return &pb.Reply{Id: 1, Status: "created"}, nil
}

func (s *server) Set(context.Context, *pb.SetRequest) (*pb.Reply, error) {
	return &pb.Reply{Id: 1, Status: "updated"}, nil
}

func (s *server) Delete(context.Context, *pb.DeleteGetRequest) (*pb.Reply, error) {
	return &pb.Reply{Id: 1, Status: "deleted"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMyGRPCServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
