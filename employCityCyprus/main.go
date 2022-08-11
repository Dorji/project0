package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "employCity/api/grpc_memcached"
	"employCity/models/interfaces"
	"employCity/models/objects"
	"employCity/models/storages/local"
	"employCity/models/storages/memcachedModel"

	"google.golang.org/grpc"
)

type server struct {
	StorageCurrent interfaces.Storage
	pb.UnimplementedMyGRPCServer
}

func (s *server) Set(ctx context.Context, set *pb.SetRequest) (*pb.Reply, error) {
	reply, err := objects.SaveToStorage(set, s.StorageCurrent)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *server) Get(ctx context.Context, get *pb.GetRequest) (*pb.ReplyGet, error) {

	reply, err := objects.GetFromStorage(get, s.StorageCurrent)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *server) Delete(ctx context.Context, delete *pb.DeleteRequest) (*pb.Reply, error) {

	reply, err := objects.DeleteFromStorage(delete, s.StorageCurrent)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("not enough arguments")
	}

	var sl interfaces.Storage
	if flag.Arg(0) == "LOCAL" {

		sl = &local.StorageLocal{}
		sl = sl.Conn()

	} else if flag.Arg(0) == "MEMCACHED" {

		sl = &memcachedModel.StorageMemcached{}
		sl = sl.Conn()
	} else {
		log.Fatalf("Not found storage by name: %v", flag.Arg(0))
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMyGRPCServer(s, &server{StorageCurrent: sl})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
