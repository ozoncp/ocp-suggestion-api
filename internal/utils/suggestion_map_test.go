package utils

import (
	"errors"
	"testing"
	"github.com/ozoncp/ocp-suggestion-api/internal/models"
	"github.com/stretchr/testify/assert"
)

//TestSplitToBulksSuccess - тесты ConvertSliceToMap без ожидаемых ошибок
func TestConvertSliceToMapSuccess(t *testing.T) {
	data := []models.Suggestion{
		{ID: 1, UserID: 1, CourseID: 1},
		{ID: 2, UserID: 2, CourseID: 2},
		{ID: 3, UserID: 3, CourseID: 3},
		{ID: 4, UserID: 4, CourseID: 4},
		{ID: 5, UserID: 5, CourseID: 5},
	}
	want := map[uint64]models.Suggestion{
		1: {ID: 1, UserID: 1, CourseID: 1},
		2: {ID: 2, UserID: 2, CourseID: 2},
		3: {ID: 3, UserID: 3, CourseID: 3},
		4: {ID: 4, UserID: 4, CourseID: 4},
		5: {ID: 5, UserID: 5, CourseID: 5},
	}

	{
		result, err := ConvertSliceToMap(data)
		assert.Equal(t, nil, err)
		assert.Equal(t, want, result)
		assert.Equal(t, len(want), len(result))
	}
}

//TestSplitToBulksFail - тесты ConvertSliceToMap с ожидаемыми ошибками
func TestConvertSliceToMapFail(t *testing.T) {
	data := []models.Suggestion{
		{ID: 1, UserID: 1, CourseID: 1},
		{ID: 2, UserID: 2, CourseID: 2},
		{ID: 5, UserID: 53, CourseID: 53},
		{ID: 4, UserID: 4, CourseID: 4},
		{ID: 5, UserID: 5, CourseID: 5},
	}
	var dataNil []models.Suggestion

	{
		result, err := ConvertSliceToMap(data)
		assert.Equal(t, true, errors.Is(err, ErrDuplicateKey))
		assert.Equal(t, map[uint64]models.Suggestion(nil), result)
	}
	{
		result, err := ConvertSliceToMap(dataNil)
		assert.Equal(t, true, errors.Is(err, ErrSliceIsNil))
		assert.Equal(t, map[uint64]models.Suggestion(nil), result)
	}
}
