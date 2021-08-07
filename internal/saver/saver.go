package saver

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/ozoncp/ocp-suggestion-api/internal/flusher"
	"github.com/ozoncp/ocp-suggestion-api/internal/models"
)

//ErrNilPointer показывает, что передан нулевой указатель
var ErrNilPointer = errors.New("pointer is nil")

//ErrCapacityOverflow показывает, что превышена допустимая ёмкость Saver
var ErrCapacityOverflow = errors.New("overflow capacity")

//Saver - интерфейс для периодического сохранения в хранилище данные через Flusher
type Saver interface {
	Save(suggestion models.Suggestion) error
	Close(ctx context.Context) error
}

//NewSaver возвращает Saver с поддержкой периодического сохранения
func NewSaver(ctx context.Context, capacity uint, flusher flusher.Flusher, period time.Duration) Saver {
	s := &saver{
		capacity:    capacity,
		flusher:     flusher,
		flushPeriod: period,
		suggestions: make([]models.Suggestion, 0, capacity),
		closeCh:     make(chan struct{}),
	}

	go s.run(ctx)
	return s
}

type saver struct {
	capacity    uint
	flusher     flusher.Flusher
	flushPeriod time.Duration
	suggestions []models.Suggestion
	closeCh     chan struct{}
	sync.Mutex
}

//Save сохраняет suggestion во внутренней структуре
//При переполнении capacity возвращается ошибка ErrCapacityOverflow
func (s *saver) Save(suggestion models.Suggestion) error {
	if s == nil {
		return ErrNilPointer
	}
	s.Lock()
	defer s.Unlock()
	if len(s.suggestions) >= int(s.capacity) {
		return ErrCapacityOverflow
	}
	s.suggestions = append(s.suggestions, suggestion)

	return nil
}

//flush сбрасывает данные через интерфейс flusher
func (s *saver) flush(ctx context.Context) (int, error) {
	s.Lock()
	defer s.Unlock()
	if len(s.suggestions) > 0 {
		var err error
		s.suggestions, err = s.flusher.Flush(ctx, s.suggestions)
		return len(s.suggestions), err
	}
	return 0, nil
}

//run запускает цикл сохранения данных с заданным периодом
func (s *saver) run(ctx context.Context) {
	ticker := time.NewTicker(s.flushPeriod)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("finish saver by context done: %v\n", ctx.Err())
			close(s.closeCh)
			return
		case <-s.closeCh:
			return
		case <-ticker.C:
			if leftNum, err := s.flush(ctx); err != nil {
				fmt.Printf("flush when ticker (%d suggestions left): %v\n",
					leftNum,
					err,
				)
			}
		}
	}
}

//Close выполняет сброс всех данных во Flusher и завершает цикл сохранения данных run()
func (s *saver) Close(ctx context.Context) error {
	if s == nil {
		return ErrNilPointer
	}
	close(s.closeCh)

	leftNum, err := s.flush(ctx)
	if err != nil {
		return fmt.Errorf("close (%d suggestions left): %w",
			leftNum,
			err,
		)
	}
	return nil
}
