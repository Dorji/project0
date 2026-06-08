package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================================
// СТРУКТУРЫ ДАННЫХ
// ============================================================

// item — значение с опциональным TTL
type item struct {
	value  any   // само значение (interface{} — аллокация на куче!)
	expiry int64 // unix nano, 0 = бессрочно
}

// shard — один сегмент данных со своим мьютексом
type shard struct {
	mu   sync.RWMutex
	data map[string]item
	// Можно добавить метрики:
	// hits   uint64
	// misses uint64
}

// ShardedStorage — основная структура
type ShardedStorage struct {
	shards    []*shard
	shardMask uint32 // битовая маска для быстрого взятия остатка
}

// ============================================================
// КОНСТРУКТОР
// ============================================================

func NewSharded(shardCount int) *ShardedStorage {
	// Проверяем, что shardCount — степень двойки
	// Это критично для операции & вместо %
	if shardCount <= 0 || (shardCount&(shardCount-1)) != 0 {
		panic("shardCount must be a power of two")
	}

	shards := make([]*shard, shardCount)
	for i := range shards {
		shards[i] = &shard{
			data: make(map[string]item),
		}
	}

	return &ShardedStorage{
		shards:    shards,
		shardMask: uint32(shardCount - 1), // например, для 16 шардов: 0b1111
	}
}
func (s *ShardedStorage) getShard(key string) *shard {
	// Вариант 1: FNV-1a (хорош для коротких строк)
	hash := uint32(2166136261) // offset basis
	for i := 0; i < len(key); i++ {
		hash ^= uint32(key[i])
		hash *= 16777619 // prime
	}
	return s.shards[hash&s.shardMask]

	// Вариант 2: через hash/fnv (медленнее из-за аллокаций интерфейса)
	// h := fnv.New32a()
	// h.Write([]byte(key))
	// return s.shards[h.Sum32() & s.shardMask]

	// Вариант 3: unsafe string to bytes (самый быстрый, но небезопасный)
	// См. ниже в разделе "Продвинутые оптимизации"
}

// ============================================================
// ОСНОВНЫЕ ОПЕРАЦИИ
// ============================================================

// Set — запись с TTL
func (s *ShardedStorage) Set(key string, value any, ttl time.Duration) {
	sh := s.getShard(key) // выбор шарда ДО блокировки

	sh.mu.Lock()
	defer sh.mu.Unlock()

	var expiry int64
	if ttl > 0 {
		expiry = time.Now().Add(ttl).UnixNano()
		// Важно: time.Now() — syscall, дорогая операция
		// В высоконагруженных системах используют кэшированное время
	}

	sh.data[key] = item{
		value:  value,
		expiry: expiry,
	}
}

// Get — чтение с ленивым удалением истёкших ключей
func (s *ShardedStorage) Get(key string) (any, bool) {
	sh := s.getShard(key)

	// Быстрый путь: читаем под RLock
	sh.mu.RLock()
	it, ok := sh.data[key]
	sh.mu.RUnlock()

	if !ok {
		return nil, false
	}

	// Проверка TTL без блокировки (it.expiry не изменится)
	if it.expiry > 0 && time.Now().UnixNano() > it.expiry {
		// Медленный путь: нужно удалить истёкший ключ
		sh.mu.Lock()

		// DOUBLE-CHECK: за время между RUnlock и Lock
		// другой писатель мог удалить ключ или перезаписать с новым TTL
		current, exists := sh.data[key]
		if exists && current.expiry == it.expiry {
			delete(sh.data, key)
		}

		sh.mu.Unlock()
		return nil, false
	}

	return it.value, true
}

// Delete — явное удаление
func (s *ShardedStorage) Delete(key string) {
	sh := s.getShard(key)
	sh.mu.Lock()
	delete(sh.data, key)
	sh.mu.Unlock()
}

// ============================================================
// ВСПОМОГАТЕЛЬНЫЕ МЕТОДЫ
// ============================================================

// Len — возвращает общее количество ключей (приблизительно, без глобальной блокировки)
func (s *ShardedStorage) Len() int {
	total := 0
	for _, sh := range s.shards {
		sh.mu.RLock()
		total += len(sh.data)
		sh.mu.RUnlock()
	}
	return total
}

// CleanExpired — активная очистка (можно вызывать периодически)
func (s *ShardedStorage) CleanExpired() int {
	now := time.Now().UnixNano()
	cleaned := 0

	for _, sh := range s.shards {
		sh.mu.Lock()
		for k, v := range sh.data {
			if v.expiry > 0 && v.expiry < now {
				delete(sh.data, k)
				cleaned++
			}
		}
		sh.mu.Unlock()
	}

	return cleaned
}

func main() {
	// Создаём хранилище с 16 шардами
	cache := NewSharded(16)

	// ======== Базовые операции ========
	fmt.Println("=== Базовые операции ===")

	// Запись
	cache.Set("user:1", "Alice", 0)                     // без TTL
	cache.Set("session:abc", "token123", 2*time.Second) // с TTL 2 сек

	// Чтение
	if val, ok := cache.Get("user:1"); ok {
		fmt.Printf("user:1 = %v\n", val)
	}

	// Демонстрация TTL
	fmt.Println("\n=== Демонстрация TTL ===")
	if val, ok := cache.Get("session:abc"); ok {
		fmt.Printf("session:abc = %v (до истечения)\n", val)
	}

	time.Sleep(3 * time.Second)

	if _, ok := cache.Get("session:abc"); !ok {
		fmt.Println("session:abc истёк и удалён")
	}

	// Удаление
	cache.Delete("user:1")
	if _, ok := cache.Get("user:1"); !ok {
		fmt.Println("user:1 успешно удалён")
	}

	fmt.Printf("\nТекущий размер: %d\n", cache.Len())
}
