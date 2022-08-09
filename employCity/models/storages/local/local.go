package local

import (
	"fmt"

	pb "employCity/api/grpc_memcached"
)

type StorageLocal struct {
	storage map[int32]string
}

func (s *StorageLocal) Get(dgr *pb.DeleteGetRequest) (*pb.Reply, error) {
	if _, ok := s.storage[dgr.Id]; !ok {
		return &pb.Reply{}, fmt.Errorf("")
	}

	return &pb.Reply{
		Id:     dgr.Id,
		Status: "",
	}, nil
}

func (s *StorageLocal) Set(sr *pb.SetRequest) (*pb.Reply, error) {
	s.storage[sr.Id] = sr.Body
	return &pb.Reply{Id: sr.Id, Status: "created"}, nil
}

func (s *StorageLocal) Delete(dgr *pb.DeleteGetRequest) (*pb.Reply, error) {

	delete(s.storage, dgr.Id)
	return &pb.Reply{Id: dgr.Id, Status: "deleted"}, nil
}
