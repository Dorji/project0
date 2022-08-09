package local

import (
	"context"
	"fmt"

	pb "employCity/api/grpc_memcached"
)

type StorageLocal struct {
	storage map[int32]string
}

func (s *StorageLocal) Get(ctx context.Context, dgr *pb.DeleteGetRequest) (*pb.Reply, error) {
	if _, ok := s.storage[dgr.Id]; !ok {
		return &pb.Reply{}, fmt.Errorf("")
	}

	return &pb.Reply{
		Id:     dgr.Id,
		Status: "",
	}, nil
}

func (s *StorageLocal) Set(ctx context.Context, sr *pb.SetRequest) (*pb.Reply, error) {
	s.storage[sr.Id] = sr.Body
	return &pb.Reply{Id: sr.Id, Status: "created"}, nil
}

func (s *StorageLocal) Delete(ctx context.Context, dgr *pb.DeleteGetRequest) (*pb.Reply, error) {

	delete(s.storage, dgr.Id)
	return &pb.Reply{Id: dgr.Id, Status: "deleted"}, nil
}
