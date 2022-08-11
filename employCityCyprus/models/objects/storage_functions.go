package objects

import (
	"fmt"

	pb "employCity/api/grpc_memcached"
	"employCity/models/interfaces"
)

func SaveToStorage(set *pb.SetRequest, storage interfaces.Storage) (*pb.Reply, error) {
	reply, err := storage.Set(set)
	if err != nil {
		return reply, fmt.Errorf("save eror: %w", err)
	}
	return reply, nil
}

func GetFromStorage(get *pb.GetRequest, storage interfaces.Storage) (*pb.ReplyGet, error) {
	reply, err := storage.Get(get)
	if err != nil {
		return nil, fmt.Errorf("save eror: %w", err)
	}
	return reply, nil
}

func DeleteFromStorage(del *pb.DeleteRequest, storage interfaces.Storage) (*pb.Reply, error) {
	reply, err := storage.Delete(del)
	if err != nil {
		return nil, fmt.Errorf("save eror: %w", err)
	}
	return reply, nil
}
