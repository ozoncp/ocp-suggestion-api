package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/ozoncp/ocp-suggestion-api/internal/models"
)

var ErrSuggestionNotFound = errors.New("suggestion not found")

//Repo - интерфейс хранилища для сущности Suggestion
type Repo interface {
	AddSuggestions(ctx context.Context, suggestions []models.Suggestion) error
	CreateSuggestion(ctx context.Context, suggestion models.Suggestion) (uint64, error)
	DescribeSuggestion(ctx context.Context, suggestionId uint64) (*models.Suggestion, error)
	ListSuggestions(ctx context.Context, limit, offset uint64) ([]models.Suggestion, error)
	UpdateSuggestion(ctx context.Context, suggestion models.Suggestion) error
	RemoveSuggestion(ctx context.Context, suggestionId uint64) error
}

type repo struct {
	db *sqlx.DB
}

// NewRepo возвращает структуру repo (интерфейс Repo)
func NewRepo(db *sqlx.DB) *repo {
	return &repo{db: db}
}

// CreateSuggestion создает предложение курса в БД
func (r repo) CreateSuggestion(ctx context.Context, suggestion models.Suggestion) (uint64, error) {
	query := squirrel.Insert("suggestions").
		Columns("user_id", "course_id").
		Values(suggestion.UserID, suggestion.CourseID).
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	if sqlStr, args, err := query.ToSql(); err != nil {
		log.Printf("CreateSuggestion ToSql error: %v", err)
	} else {
		log.Printf("CreateSuggestion sql: %s, args: %v", sqlStr, args)
	}

	err := query.QueryRowContext(ctx).Scan(&suggestion.ID)
	if err != nil {
		return 0, err
	}

	return suggestion.ID, nil
}

// AddSuggestions создает несколько предложений курсов в БД
func (r repo) AddSuggestions(ctx context.Context, suggestions []models.Suggestion) error {
	query := squirrel.
		Insert("suggestions").
		Columns("user_id", "course_id").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	for _, suggestion := range suggestions {
		query = query.Values(suggestion.UserID, suggestion.CourseID)
	}

	if sqlStr, args, err := query.ToSql(); err != nil {
		log.Printf("AddSuggestions ToSql error: %v", err)
	} else {
		log.Printf("AddSuggestions sql: %s, args: %v", sqlStr, args)
	}

	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

// DescribeSuggestion возвращает описание предложения курса из БД
func (r repo) DescribeSuggestion(ctx context.Context, suggestionId uint64) (*models.Suggestion, error) {
	query := squirrel.Select("id", "user_id", "course_id").
		From("suggestions").
		Where(squirrel.Eq{"id": suggestionId}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	if sqlStr, args, err := query.ToSql(); err != nil {
		log.Printf("DescribeSuggestion ToSql error: %v", err)
	} else {
		log.Printf("DescribeSuggestion sql: %s, args: %v", sqlStr, args)
	}

	var suggestion models.Suggestion
	err := query.QueryRowContext(ctx).Scan(&suggestion.ID, &suggestion.UserID, &suggestion.CourseID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("id: %d, %w", suggestionId, ErrSuggestionNotFound)
	} else if err != nil {
		return nil, err
	}

	return &suggestion, nil
}

// ListSuggestions возвращает список предложений курсов из БД
func (r repo) ListSuggestions(ctx context.Context, limit, offset uint64) ([]models.Suggestion, error) {
	query := squirrel.Select("id", "user_id", "course_id").
		From("suggestions").
		RunWith(r.db).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(squirrel.Dollar)

	if sqlStr, args, err := query.ToSql(); err != nil {
		log.Printf("ListSuggestions ToSql error: %v", err)
	} else {
		log.Printf("ListSuggestions sql: %s, args: %v", sqlStr, args)
	}

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var suggestions []models.Suggestion
	for rows.Next() {
		var suggestion models.Suggestion
		if err := rows.Scan(&suggestion.ID, &suggestion.UserID, &suggestion.CourseID); err != nil {
			return nil, err
		}
		suggestions = append(suggestions, suggestion)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return suggestions, nil
}

// UpdateSuggestion обновляет описание предложения курса в БД
func (r repo) UpdateSuggestion(ctx context.Context, suggestion models.Suggestion) error {
	query := squirrel.Update("suggestions").
		Set("user_id", suggestion.UserID).
		Set("course_id", suggestion.CourseID).
		Where(squirrel.Eq{"id": suggestion.ID}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	if sqlStr, args, err := query.ToSql(); err != nil {
		log.Printf("UpdateSuggestion ToSql error: %v", err)
	} else {
		log.Printf("UpdateSuggestion sql: %s, args: %v", sqlStr, args)
	}

	result, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected <= 0 {
		return fmt.Errorf("id: %d, %w", suggestion.ID, ErrSuggestionNotFound)
	}

	return nil
}

// RemoveSuggestion удаляет предложение курса из БД
func (r repo) RemoveSuggestion(ctx context.Context, suggestionId uint64) error {
	query := squirrel.Delete("suggestions").
		Where(squirrel.Eq{"id": suggestionId}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	if sqlStr, args, err := query.ToSql(); err != nil {
		log.Printf("RemoveSuggestion ToSql error: %v", err)
	} else {
		log.Printf("RemoveSuggestion sql: %s, args: %v", sqlStr, args)
	}

	result, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected <= 0 {
		return fmt.Errorf("id: %d, %w", suggestionId, ErrSuggestionNotFound)
	}

	return nil
}
