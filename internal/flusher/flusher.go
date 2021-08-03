package flusher

import (
	"context"
	"errors"
	"fmt"

	"github.com/ozoncp/ocp-suggestion-api/internal/models"
	"github.com/ozoncp/ocp-suggestion-api/internal/repo"
	"github.com/ozoncp/ocp-suggestion-api/internal/utils"
)

// Flusher - интерфейс для сброса задач в хранилище
type Flusher interface {
	Flush(ctx context.Context, suggestion []models.Suggestion) ([]models.Suggestion, error)
}

// NewFlusher возвращает Flusher с поддержкой батчевого сохранения
func NewFlusher(chunkSize uint, suggestionRepo repo.Repo) Flusher {
	return &flusher{
		chunkSize:      chunkSize,
		suggestionRepo: suggestionRepo,
	}
}

type flusher struct {
	chunkSize      uint
	suggestionRepo repo.Repo
}

//Flush сбрасывает слайс Suggestion в хранилище частями заданного размера (чанками).
//Если при сбросе возникает ошибка, то несохранённый остаток возвращается функцией
func (f *flusher) Flush(ctx context.Context, suggestions []models.Suggestion) ([]models.Suggestion, error) {
	if f == nil {
		return nil, errors.New("interface is nil")
	}
	splitSuggestions, err := utils.SplitToBulks(suggestions, f.chunkSize)
	if err != nil {
		return suggestions, fmt.Errorf("SplitToBulks : %w", err)
	}
	for i, chunk := range splitSuggestions {
		if err := f.suggestionRepo.AddSuggestions(ctx, chunk); err != nil {
			return suggestions[uint(i)*f.chunkSize:], fmt.Errorf("AddSuggestions : %w", err)
		}
	}
	return nil, nil
}
