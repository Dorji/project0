package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "my_mq_server/api/server/producerMQ"
	"net"
)

type serverMQ struct {
	pb.UnimplementedProducerMQServer
}

func (smq *serverMQ) HandleMessage(context.Context, *pb.SendRequest) (*pb.Reply, error) {
	return nil, nil
}
func (smq *serverMQ) GetQueueList(context.Context, *pb.GetQueuesListRequest) (*pb.ListQueues, error) {
	return nil, nil
}
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 1337))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProducerMQServer(s, &serverMQ{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
