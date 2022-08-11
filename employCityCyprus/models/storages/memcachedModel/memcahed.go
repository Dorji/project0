package memcachedModel

import (
	"fmt"
	"net"
	"time"

	pb "employCity/api/grpc_memcached"
	"employCity/models/interfaces"

	"github.com/bradfitz/gomemcache/memcache"
)

type StorageMemcached struct {
	conn *memcache.Client
}

func (sM *StorageMemcached) Conn() interfaces.Storage {
	sM.conn = memcache.New("127.0.0.1:11211")
	cp, _ := NewConnPool(func() (ConnRes, error) {
		return net.Dial("tcp", ":8080")
	}, 10, time.Second*10)

	return sM
}

func (sM *StorageMemcached) Set(set *pb.SetRequest) (*pb.Reply, error) {
	err := sM.conn.Set(&memcache.Item{Key: set.GetId(), Value: []byte(set.GetBody())})
	if err != nil {
		return nil, err
	}
	return &pb.Reply{Id: set.Id, Status: "created"}, nil
}

func (sM *StorageMemcached) Get(get *pb.GetRequest) (*pb.ReplyGet, error) {
	val, err := sM.conn.Get(get.GetId())
	if err != nil {
		return nil, fmt.Errorf("get error in storage memcached: %w", err)
	}

	return &pb.ReplyGet{
		Id:   get.Id,
		Body: string(val.Value),
	}, nil
}

func (sM *StorageMemcached) Delete(del *pb.DeleteRequest) (*pb.Reply, error) {
	err := sM.conn.Delete(del.Id)
	if err != nil {
		return nil, fmt.Errorf("delete error in storage memcached: %w", err)
	}
	return &pb.Reply{Id: del.Id, Status: "deleted"}, nil
}
