package storages

import (
	"context"

	pb "employCity/api/grpc_memcached"
)

type Storage interface {
	Get(context.Context, *pb.DeleteGetRequest) (*pb.Reply, error)
	Set(context.Context, *pb.SetRequest) (*pb.Reply, error)
	Delete(context.Context, *pb.DeleteGetRequest) (*pb.Reply, error)
}
