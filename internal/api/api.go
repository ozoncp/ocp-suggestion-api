package api

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ozoncp/ocp-suggestion-api/internal/models"
	"github.com/ozoncp/ocp-suggestion-api/internal/repo"
	desc "github.com/ozoncp/ocp-suggestion-api/pkg/ocp-suggestion-api"
)

type suggestionAPI struct {
	desc.UnimplementedOcpSuggestionApiServer
	repo repo.Repo
}

// NewSuggestionAPI возвращает suggestionAPI
func NewSuggestionAPI(repo repo.Repo) *suggestionAPI {
	return &suggestionAPI{repo: repo}
}

// CreateSuggestionV1 создаёт предложение курса и возвращает его ID
func (r *suggestionAPI) CreateSuggestionV1(ctx context.Context, req *desc.CreateSuggestionV1Request) (*desc.CreateSuggestionV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Printf("CreateSuggestionV1 request: %v", req)
	createdID, err := r.repo.CreateSuggestion(ctx, models.Suggestion{
		ID:       0,
		CourseID: req.CourseId,
		UserID:   req.UserId,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
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
	suggestion, err := r.repo.DescribeSuggestion(ctx, req.SuggestionId)
	if errors.Is(err, repo.ErrSuggestionNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
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
	suggestions, err := r.repo.ListSuggestions(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("ListSuggestionV1 response:")
	Suggestions := make([]*desc.Suggestion, 0, len(suggestions))
	for i, suggestion := range suggestions {
		log.Printf("suggestions[%d]: %v", i, suggestion)
		Suggestions = append(Suggestions, &desc.Suggestion{
			Id:       suggestion.ID,
			CourseId: suggestion.CourseID,
			UserId:   suggestion.UserID})
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
	err := r.repo.UpdateSuggestion(ctx, models.Suggestion{
		ID:       req.Suggestion.Id,
		CourseID: req.Suggestion.CourseId,
		UserID:   req.Suggestion.UserId,
	})

	if errors.Is(err, repo.ErrSuggestionNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Printf("UpdateSuggestionV1 response: {}")

	return &desc.UpdateSuggestionV1Response{}, nil
}

// RemoveSuggestionV1 удаляет предложение с заданным id
func (r *suggestionAPI) RemoveSuggestionV1(ctx context.Context, req *desc.RemoveSuggestionV1Request) (*desc.RemoveSuggestionV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Printf("RemoveSuggestionV1 request: %v", req)
	err := r.repo.RemoveSuggestion(ctx, req.SuggestionId)
	if errors.Is(err, repo.ErrSuggestionNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Printf("RemoveSuggestionV1 response: {}")

	return &desc.RemoveSuggestionV1Response{}, nil
}
