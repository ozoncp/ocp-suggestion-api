package repo

import "github.com/ozoncp/ocp-suggestion-api/internal/models"

//Repo - интерфейс хранилища для сущности Suggestion
type Repo interface {
	AddSuggestions(suggestion []models.Suggestion) error
	ListSuggestions(limit, offset uint64) ([]models.Suggestion, error)
	DescribeSuggestion(suggestionId uint64) (*models.Suggestion, error)
}
