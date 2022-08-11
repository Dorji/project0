package local

import (
	"fmt"

	pb "employCity/api/grpc_memcached"
	"employCity/models/interfaces"
)

type StorageLocal struct {
	storage map[string]string
}

func (s *StorageLocal) Conn() interfaces.Storage {
	s.storage = map[string]string{}
	return s
}

func (s *StorageLocal) Set(set *pb.SetRequest) (*pb.Reply, error) {
	s.storage[set.Id] = set.Body
	return &pb.Reply{Id: set.Id, Status: "created"}, nil
}

func (s *StorageLocal) Get(get *pb.GetRequest) (*pb.ReplyGet, error) {
	if _, ok := s.storage[get.Id]; !ok {
		return nil, fmt.Errorf("not found in storage local")
	}

	return &pb.ReplyGet{
		Id:   get.Id,
		Body: s.storage[get.Id],
	}, nil
}

func (s *StorageLocal) Delete(del *pb.DeleteRequest) (*pb.Reply, error) {
	if _, ok := s.storage[del.Id]; !ok {
		return nil, fmt.Errorf("not found in storage local")
	}
	delete(s.storage, del.Id)

	return &pb.Reply{Id: del.Id, Status: "deleted"}, nil
}
