package interfaces

import (
	pb "employCity/api/grpc_memcached"
)

type Storage interface {
	Get(*pb.GetRequest) (*pb.ReplyGet, error)
	Set(*pb.SetRequest) (*pb.Reply, error)
	Delete(*pb.DeleteRequest) (*pb.Reply, error)
}
