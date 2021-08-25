package api

import (
	"context"
	"encoding/binary"
	"errors"

	"github.com/opentracing/opentracing-go"
	tracelog "github.com/opentracing/opentracing-go/log"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ozoncp/ocp-suggestion-api/internal/metrics"
	"github.com/ozoncp/ocp-suggestion-api/internal/models"
	"github.com/ozoncp/ocp-suggestion-api/internal/producer"
	"github.com/ozoncp/ocp-suggestion-api/internal/repo"
	"github.com/ozoncp/ocp-suggestion-api/internal/utils"
	desc "github.com/ozoncp/ocp-suggestion-api/pkg/ocp-suggestion-api"
)

type suggestionAPI struct {
	desc.UnimplementedOcpSuggestionApiServer
	repo      repo.Repo
	batchSize uint
	prod      producer.Producer
}

// NewSuggestionAPI возвращает suggestionAPI
func NewSuggestionAPI(repo repo.Repo, batchSize uint, prod producer.Producer) *suggestionAPI {
	return &suggestionAPI{
		repo:      repo,
		batchSize: batchSize,
		prod:      prod,
	}
}

// CreateSuggestionV1 создаёт предложение курса и возвращает его ID
func (r *suggestionAPI) CreateSuggestionV1(ctx context.Context, req *desc.CreateSuggestionV1Request) (*desc.CreateSuggestionV1Response, error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("CreateSuggestionV1")
	defer span.Finish()

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

	message := producer.NewMessage(createdID, producer.Create)
	err = r.prod.Send("ocp-suggestion-api", message)
	if err != nil {
		log.Error().Err(err).Msgf("CreateSuggestionV1: failed to send message to kafka")
	}

	metrics.CreateCounterInc()
	log.Printf("CreateSuggestionV1 response suggestionId: %d", createdID)

	return &desc.CreateSuggestionV1Response{
		SuggestionId: createdID,
	}, nil
}

// MultiCreateSuggestionV1 создаёт предложение курса и возвращает его ID
func (r *suggestionAPI) MultiCreateSuggestionV1(ctx context.Context, req *desc.MultiCreateSuggestionV1Request) (*desc.MultiCreateSuggestionV1Response, error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("MultiCreateSuggestionV1")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Printf("MultiCreateSuggestionV1 request: %v", req)
	suggestions := make([]models.Suggestion, 0, len(req.NewSuggestion))
	for _, newSuggestion := range req.NewSuggestion {
		suggestions = append(suggestions, models.Suggestion{
			CourseID: newSuggestion.CourseId,
			UserID:   newSuggestion.UserId,
		})
	}

	chunkSuggestions, err := utils.SplitToBulks(suggestions, r.batchSize)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var createdNumberTotal uint64
	for _, chunk := range chunkSuggestions {
		createdNumber, err := func() (uint64, error) {
			childSpan := tracer.StartSpan("MultiCreateSuggestionV1Batch", opentracing.ChildOf(span.Context()))
			childSpan.LogFields(tracelog.Int("batch_size_bytes", binary.Size(chunk)))
			defer childSpan.Finish()

			number, err := r.repo.AddSuggestions(ctx, chunk)
			if err != nil {
				return 0, status.Error(codes.Internal, err.Error())
			}
			return number, nil
		}()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		createdNumberTotal += createdNumber
	}
	log.Printf("MultiCreateSuggestionV1 response createdNumber: %d", createdNumberTotal)

	return &desc.MultiCreateSuggestionV1Response{
		CreatedNumber: createdNumberTotal,
	}, nil
}

// DescribeSuggestionV1 возвращает описание предложения с заданным id
func (r *suggestionAPI) DescribeSuggestionV1(ctx context.Context, req *desc.DescribeSuggestionV1Request) (*desc.DescribeSuggestionV1Response, error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("DescribeSuggestionV1")
	defer span.Finish()

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
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("ListSuggestionV1")
	defer span.Finish()

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
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("UpdateSuggestionV1")
	defer span.Finish()

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

	message := producer.NewMessage(req.Suggestion.Id, producer.Update)
	err = r.prod.Send("ocp-suggestion-api", message)
	if err != nil {
		log.Error().Err(err).Msgf("UpdateSuggestionV1: failed to send message to kafka")
	}

	metrics.UpdateCounterInc()
	log.Printf("UpdateSuggestionV1 response: {}")

	return &desc.UpdateSuggestionV1Response{}, nil
}

// RemoveSuggestionV1 удаляет предложение с заданным id
func (r *suggestionAPI) RemoveSuggestionV1(ctx context.Context, req *desc.RemoveSuggestionV1Request) (*desc.RemoveSuggestionV1Response, error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("RemoveSuggestionV1")
	defer span.Finish()

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

	message := producer.NewMessage(req.SuggestionId, producer.Remove)
	err = r.prod.Send("ocp-suggestion-api", message)
	if err != nil {
		log.Error().Err(err).Msgf("RemoveSuggestionV1: failed to send message to kafka")
	}

	metrics.RemoveCounterInc()
	log.Printf("RemoveSuggestionV1 response: {}")

	return &desc.RemoveSuggestionV1Response{}, nil
}
