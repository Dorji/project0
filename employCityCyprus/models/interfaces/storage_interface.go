package interfaces

import (
	pb "employCity/api/grpc_memcached"
)

type Storage interface {
	Set(set *pb.SetRequest) (*pb.Reply, error)
	Get(get *pb.GetRequest) (*pb.ReplyGet, error)
	Delete(del *pb.DeleteRequest) (*pb.Reply, error)
	Conn() Storage
}
type Requests interface {
	GetStorage() string
}

type ConnRes interface {
	Close() error
}
