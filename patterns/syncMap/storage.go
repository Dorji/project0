package main

import (
	"sync"
)

type ShardedStorage struct {
	mp  map[string]*shard
}

type shard struct {
	mx    sync.RWMutex
	value []item
}
type item struct {
	value any
	ttl   uint64
}

func NewShardedStorage(ttlSeconds uint64) *Storage {
	return &ShardedStorage{
		mp:  make(map[string]item),
		ttl: 123,
	}
}
func Get() {
	
}
func Set() {

}

func Delete() {

}
