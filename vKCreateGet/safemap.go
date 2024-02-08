package main

import "sync"

// Package join contains an interview task.
//
// Есть интерфейс SafeMap, предоставляющий конкурентную безопасную работу с map[int]int.
// Нужно реализовать его имплементацию, не используя sync.Map.
//
// Замечание: количество чтений во много раз больше числа записей.
//
// Как проверять:
// go test -race -v

// SafeMap is concurrent safe map interface.
type SafeMap interface {
	GetOrCreate(key, value int) int
}

// safeMap is an implementation of SafeMap interface.
type safeMap struct {
	mu   sync.RWMutex
	data map[int]int
}

//func NewSafeMap(mu *sync.RWMutex) SafeMap {
//	myMap := make(map[int]int)
//	return &safeMap{
//		mu:   mu,
//		data: myMap,
//	}
//}

// GetOrCreate implements SafeMap interface.
func (s *safeMap) GetOrCreate(key, value int) int {
	var res int
	s.mu.RLock()
	res, ok := s.data[key]
	s.mu.RUnlock()
	if ok {
		return res
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	res, ok = s.data[key]
	if ok {
		return res
	}
	s.data[key] = value

	return value
}
