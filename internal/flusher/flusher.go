package flusher

import (
	"github.com/ozoncp/ocp-suggestion-api/internal/models"
	"github.com/ozoncp/ocp-suggestion-api/internal/repo"
	"github.com/ozoncp/ocp-suggestion-api/internal/utils"
)

// Flusher - интерфейс для сброса задач в хранилище
type Flusher interface {
	Flush(suggestion []models.Suggestion) []models.Suggestion
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
func (f *flusher) Flush(suggestions []models.Suggestion) []models.Suggestion {
	if f == nil {
		return nil
	}
	splitSuggestions, err := utils.SplitToBulks(suggestions, f.chunkSize)
	if err != nil {
		return suggestions
	}
	for i, chunk := range splitSuggestions {
		if err := f.suggestionRepo.AddSuggestions(chunk); err != nil {
			return suggestions[uint(i)*f.chunkSize:]
		}
	}
	return nil
}
