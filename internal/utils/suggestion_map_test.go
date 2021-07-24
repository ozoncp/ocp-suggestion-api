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
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 3},
		{4, 4, 4},
		{5, 5, 5},
	}
	want := map[uint64]models.Suggestion{
		1: {1, 1, 1},
		2: {2, 2, 2},
		3: {3, 3, 3},
		4: {4, 4, 4},
		5: {5, 5, 5},
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
		{1, 1, 1},
		{2, 2, 2},
		{5, 53, 53},
		{4, 4, 4},
		{5, 5, 5},
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
