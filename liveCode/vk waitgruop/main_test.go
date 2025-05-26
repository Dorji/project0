package main

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func TestLimitedGroup_Do(t *testing.T) {
	const maxConcurrent = 3
	lg := NewLimitedGroup(maxConcurrent)

	var (
		currentRunning int
		maxObserved    int
		mu             sync.Mutex
	)

	// Функция, которая увеличивает счетчик работающих горутин
	job := func() error {
		mu.Lock()
		currentRunning++
		if currentRunning > maxObserved {
			maxObserved = currentRunning
		}
		mu.Unlock()

		time.Sleep(100 * time.Millisecond) // Имитация работы

		mu.Lock()
		currentRunning--
		mu.Unlock()
		return nil
	}

	// Запускаем больше задач, чем разрешено
	for i := 0; i < maxConcurrent*2; i++ {
		lg.Do(job)
	}

	lg.Wait()

	if maxObserved != maxConcurrent {
		t.Errorf("expected max %d concurrent jobs, got %d", maxConcurrent, maxObserved)
	}
}

func TestLimitedGroup_TryDo(t *testing.T) {
	const maxConcurrent = 2
	lg := NewLimitedGroup(maxConcurrent)

	// Заполняем все слоты
	for i := 0; i < maxConcurrent; i++ {
		if !lg.TryDo(func() error {
			time.Sleep(200 * time.Millisecond)
			return nil
		}) {
			t.Error("expected TryDo to succeed when there's capacity")
		}
	}

	// Попытка добавить еще одну задачу (должна вернуть false)
	if lg.TryDo(func() error { return nil }) {
		t.Error("expected TryDo to fail when no capacity")
	}

	lg.Wait()
}

func TestLimitedGroup_ErrorHandling(t *testing.T) {
	lg := NewLimitedGroup(1)

	expectedErr := errors.New("test error")

	lg.Do(func() error {
		return expectedErr
	})

	lg.Do(func() error {
		return nil
	})

	err := lg.Wait()
	if err != expectedErr {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}
}

func TestLimitedGroup_Wait(t *testing.T) {
	lg := NewLimitedGroup(2)

	var counter int
	var mu sync.Mutex

	// Запускаем несколько задач, увеличивающих счетчик
	for i := 0; i < 5; i++ {
		lg.Do(func() error {
			mu.Lock()
			counter++
			mu.Unlock()
			return nil
		})
	}

	err := lg.Wait()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if counter != 5 {
		t.Errorf("expected counter to be 5, got %d", counter)
	}
}

func TestLimitedGroup_CloseChannelOnWait(t *testing.T) {
	lg := NewLimitedGroup(1)

	lg.Do(func() error {
		return nil
	})

	_ = lg.Wait()

	// Проверяем, что канал закрыт
	select {
	case _, ok := <-lg.chanPool:
		if ok {
			t.Error("expected channel to be closed after Wait")
		}
	default:
		t.Error("channel should be closed after Wait")
	}
}