package techtask

import (
	"context"
	"sync"
)
// Необходимо имплементировать интерфейс Manager так, чтобы он:
// - перманентно и синхронно принимал данные из Reader
// - обрабатывал все полученные данные на каждом из списка Processor
// - в случае отсутствия ошибок при обработке, полученные в процессе обработки данные передавал в Writer
// - при возникновении ошибки при обработке, ничего не отправлял во Writer

// ManagerImpl реализует интерфейс Manager
type ManagerImpl struct {
	reader     Reader       // Источник данных
	processors []Processor  // Список процессоров для обработки
	writer     Writer       // Получатель обработанных данных
	ctx        context.Context // Контекст для управления жизненным циклом
}

func main() {
	// Создаем зависимости
	reader := NewSomeReader()
	processors := []Processor{NewProcessor1(), NewProcessor2()}
	writer := NewSomeWriter()
	
	// Создаем контекст с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	// Создаем и запускаем менеджер
	manager := NewManager(reader, processors, writer, ctx)
	go manager.Manage()
	
	// ... остальная логика приложения ...
}

// NewManager создает новый экземпляр менеджера
func NewManager(r Reader, p []Processor, w Writer, ctx context.Context) Manager {
	return &ManagerImpl{
		reader:     r,
		processors: p,
		writer:     w,
		ctx:        ctx,
	}
}

// Manage - основной метод, реализующий логику управления
func (m *ManagerImpl) Manage() {
	for {
		select {
		case <-m.ctx.Done(): // Проверяем не был ли отменен контекст
			return
		default:
			// Читаем данные из источника
			dataSlice := m.reader.Read()
			if len(dataSlice) == 0 {
				continue
			}

			var wg sync.WaitGroup
			errorChan := make(chan error, len(dataSlice)*len(m.processors))
			resultChan := make(chan *Data, len(dataSlice)*len(m.processors))

			// Обрабатываем каждое данное через каждый процессор
			for _, data := range dataSlice {
				for _, processor := range m.processors {
					wg.Add(1)
					go func(d Data, p Processor) {
						defer wg.Done()
						results, err := p.Process(d)
						if err != nil {
							errorChan <- err
							return
						}
						for _, result := range results {
							resultChan <- result
						}
					}(*data, processor)
				}
			}

			// Ждем завершения всех обработчиков
			go func() {
				wg.Wait()
				close(errorChan)
				close(resultChan)
			}()

			// Проверяем наличие ошибок
			hasErrors := false
			for err := range errorChan {
				if err != nil {
					hasErrors = true
					break
				}
			}

			// Если ошибок нет - записываем результаты
			if !hasErrors {
				var results []*Data
				for result := range resultChan {
					results = append(results, result)
				}
				if len(results) > 0 {
					m.writer.Write(results)
				}
			}
		}
	}
}