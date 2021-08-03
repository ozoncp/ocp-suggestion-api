package repo

import (
	"context"
	
	"github.com/ozoncp/ocp-suggestion-api/internal/models"
)

//Repo - интерфейс хранилища для сущности Suggestion
type Repo interface {
	AddSuggestions(ctx context.Context, suggestion []models.Suggestion) error
	ListSuggestions(ctx context.Context, limit, offset uint64) ([]models.Suggestion, error)
	DescribeSuggestion(ctx context.Context, suggestionId uint64) (*models.Suggestion, error)
}
