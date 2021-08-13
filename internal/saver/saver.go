package saver

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ozoncp/ocp-suggestion-api/internal/flusher"
	"github.com/ozoncp/ocp-suggestion-api/internal/models"
)

//Saver - интерфейс для периодического сохранения в хранилище данные через Flusher
//Для запуска цикла периодического сохранения необходим вызов Init()
//Для останова необходимо использовать Close()
//Функции Init() и Close() могут быть использованы повторно
type Saver interface {
	Save(ctx context.Context, suggestion models.Suggestion) error
	Init(ctx context.Context)
	Close(ctx context.Context) error
}

//NewSaver возвращает saver с поддержкой периодического сохранения
func NewSaver(capacity uint, flusher flusher.Flusher, period time.Duration) *saver {
	s := &saver{
		capacity:    capacity,
		flusher:     flusher,
		flushPeriod: period,
		buffer:      make([]models.Suggestion, 0, capacity),
	}
	s.config.isInit = false
	s.config.isClosed = true

	return s
}

type saver struct {
	capacity    uint
	flusher     flusher.Flusher
	flushPeriod time.Duration
	sync.Mutex
	buffer      []models.Suggestion
	configGuard sync.Mutex
	config      struct {
		isInit   bool
		isClosed bool
	}
	closeCh chan struct{}
	doneCh  chan struct{}
}

//Save сохраняет suggestion во внутреннем буфере размера capacity
//При переполнении capacity выполняется внеочередной flush
func (s *saver) Save(ctx context.Context, suggestion models.Suggestion) error {
	s.buffer = append(s.buffer, suggestion)
	if len(s.buffer) > int(s.capacity) {
		if leftNum, err := s.flush(ctx); err != nil {
			return fmt.Errorf("save (%d buffer left): %w",
				leftNum,
				err,
			)
		}
	}

	return nil
}

//flush сбрасывает данные через интерфейс flusher с таймаутом, равным периоду цикла сохранения
func (s *saver) flush(ctx context.Context) (int, error) {
	s.Lock()
	defer s.Unlock()
	if len(s.buffer) > 0 {
		ctx, cancel := context.WithTimeout(ctx, s.flushPeriod)
		defer cancel()

		var err error
		s.buffer, err = s.flusher.Flush(ctx, s.buffer)
		return len(s.buffer), err
	}
	return 0, nil
}

//Init запускает цикл сохранения данных с заданным периодом
func (s *saver) Init(ctx context.Context) {
	s.configGuard.Lock()
	defer s.configGuard.Unlock()
	if s.config.isInit {
		return
	}
	s.doneCh = make(chan struct{})
	s.config.isInit = true

	if s.config.isClosed {
		s.closeCh = make(chan struct{})
		s.config.isClosed = false
	}

	go func() {
		defer func() {
			close(s.doneCh)
			s.configGuard.Lock()
			s.config.isInit = false
			s.configGuard.Unlock()
		}()
		ticker := time.NewTicker(s.flushPeriod)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				fmt.Printf("finish saver by context done: %v\n", ctx.Err())
				return
			case <-s.closeCh:
				return
			case <-ticker.C:
				if leftNum, err := s.flush(ctx); err != nil {
					fmt.Printf("flush when ticker (%d buffer left): %v\n",
						leftNum,
						err,
					)
				}
			}
		}
	}()
}

//Close выполняет сброс всех данных во Flusher и завершает периодическое сохранение данных в Init()
func (s *saver) Close(ctx context.Context) error {
	s.configGuard.Lock()
	defer s.configGuard.Unlock()
	if !s.config.isClosed {
		s.config.isClosed = true
		close(s.closeCh)
	}

	leftNum, err := s.flush(ctx)
	if err != nil {
		return fmt.Errorf("close (%d buffer left): %w",
			leftNum,
			err,
		)
	}
	if s.config.isInit {
		<-s.doneCh
	}

	return nil
}
