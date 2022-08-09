package atomyzeServer

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "atomyzeServer/api/transactions"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnsafeTransactionsGRPCServer
}

func (s server) CreateTransaction(ctx context.Context, request *pb.CreateRequest) (*pb.CreateReply, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) VerifyTransaction(ctx context.Context, request *pb.VerifyRequest) (*pb.VerifyReplyGet, error) {
	//TODO implement me
	panic("implement me")
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 1337))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTransactionsGRPCServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
