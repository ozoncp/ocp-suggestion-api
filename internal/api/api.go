package api

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ozoncp/ocp-suggestion-api/internal/models"
	desc "github.com/ozoncp/ocp-suggestion-api/pkg/ocp-suggestion-api"
)

type suggestionAPI struct {
	desc.UnimplementedOcpSuggestionApiServer
}

// NewSuggestionAPI возвращает suggestionAPI
func NewSuggestionAPI() *suggestionAPI {
	return &suggestionAPI{}
}

// CreateSuggestionV1 создаёт предложение курса и возвращает его ID
func (r *suggestionAPI) CreateSuggestionV1(ctx context.Context, req *desc.CreateSuggestionV1Request) (*desc.CreateSuggestionV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Printf("CreateSuggestionV1 request: %v", req)

	// TODO: заменить на возвращаемый ID из БД, когда будет реализована работа с БД
	createdID := uint64(time.Now().Unix())*1000000000 + uint64(time.Now().Nanosecond())
	log.Printf("CreateSuggestionV1 response suggestionId: %d", createdID)

	return &desc.CreateSuggestionV1Response{
		SuggestionId: createdID,
	}, nil
}

// DescribeSuggestionV1 возвращает описание предложения с заданным id
func (r *suggestionAPI) DescribeSuggestionV1(ctx context.Context, req *desc.DescribeSuggestionV1Request) (*desc.DescribeSuggestionV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Printf("DescribeSuggestionV1 request: %v", req)

	// TODO: заменить на возвращаемую информацию из БД, когда будет реализована работа с БД
	suggestion := models.Suggestion{
		ID:       req.SuggestionId,
		CourseID: 2,
		UserID:   3,
	}
	log.Printf("DescribeSuggestionV1 response: %v", suggestion)

	return &desc.DescribeSuggestionV1Response{
		Suggestion: &desc.Suggestion{
			Id:       suggestion.ID,
			UserId:   suggestion.UserID,
			CourseId: suggestion.CourseID,
		},
	}, nil
}

// ListSuggestionV1 возвращает список предложений
func (r *suggestionAPI) ListSuggestionV1(ctx context.Context, req *desc.ListSuggestionV1Request) (*desc.ListSuggestionV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Printf("ListSuggestionV1 request: %v", req)

	// TODO: заменить на возвращаемую информацию из БД, когда будет реализована работа с БД
	suggestions := []models.Suggestion{
		{ID: 1, UserID: 2, CourseID: 3},
		{ID: 2, UserID: 3, CourseID: 4},
	}
	log.Printf("ListSuggestionV1 response:")
	Suggestions := make([]*desc.Suggestion, len(suggestions))
	for i, suggestion := range suggestions {
		log.Printf("suggestions[%d]: %v", i, suggestion)
		Suggestions[i] = new(desc.Suggestion)
		Suggestions[i].Id = suggestion.ID
		Suggestions[i].CourseId = suggestion.CourseID
		Suggestions[i].UserId = suggestion.UserID
	}

	return &desc.ListSuggestionV1Response{
		Suggestions: Suggestions,
	}, nil
}

// UpdateSuggestionV1 обновляет предложение
func (r *suggestionAPI) UpdateSuggestionV1(ctx context.Context, req *desc.UpdateSuggestionV1Request) (*desc.UpdateSuggestionV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Printf("UpdateSuggestionV1 request: %v", req)
	// TODO: реализовать обновление информации в БД, когда будет реализована работа с БД
	log.Printf("UpdateSuggestionV1 response: {}")

	return &desc.UpdateSuggestionV1Response{}, nil
}

// RemoveSuggestionV1 удаляет предложение с заданным id
func (r *suggestionAPI) RemoveSuggestionV1(ctx context.Context, req *desc.RemoveSuggestionV1Request) (*desc.RemoveSuggestionV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Printf("RemoveSuggestionV1 request: %v", req)
	// TODO: реализовать обновление информации в БД, когда будет реализована работа с БД
	log.Printf("RemoveSuggestionV1 response: {}")

	return &desc.RemoveSuggestionV1Response{}, nil
}
