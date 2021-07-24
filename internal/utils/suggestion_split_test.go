package utils

import (
	"errors"
	"testing"
	"github.com/ozoncp/ocp-suggestion-api/internal/models"
	"github.com/stretchr/testify/assert"
)

//TestSplitToBulksSuccess - тесты SplitToBulks без ожидаемых ошибок
func TestSplitToBulksSuccess(t *testing.T) {
	data := []models.Suggestion{
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 3},
		{4, 4, 4},
		{5, 5, 5},
	}
	batchSize1 := uint(2)
	want1 := [][]models.Suggestion{
		{
			{1, 1, 1},
			{2, 2, 2},
		},
		{
			{3, 3, 3},
			{4, 4, 4},
		},
		{
			{5, 5, 5},
		},
	}
	batchSize2 := uint(1)
	want2 := [][]models.Suggestion{
		{
			{1, 1, 1},
		},
		{
			{2, 2, 2},
		},
		{
			{3, 3, 3},
		},
		{
			{4, 4, 4},
		},
		{
			{5, 5, 5},
		},
	}
	batchSize3 := uint(5)
	want3 := [][]models.Suggestion{
		{
			{1, 1, 1},
			{2, 2, 2},
			{3, 3, 3},
			{4, 4, 4},
			{5, 5, 5},
		},
	}

	{
		result, err := SplitToBulks(data, batchSize1)
		assert.Equal(t, nil, err)
		assert.Equal(t, want1, result)
	}
	{
		result, err := SplitToBulks(data, batchSize2)
		assert.Equal(t, nil, err)
		assert.Equal(t, want2, result)
	}
	{
		result, err := SplitToBulks(data, batchSize3)
		assert.Equal(t, nil, err)
		assert.Equal(t, want3, result)
	}
}

//TestSplitToBulksFail - тесты SplitToBulks с ожидаемыми ошибками
func TestSplitToBulksFail(t *testing.T) {
	data := []models.Suggestion{
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 3},
		{4, 4, 4},
		{5, 5, 5},
	}
	var dataNil []models.Suggestion

	{
		result, err := SplitToBulks(data, 0)
		assert.Equal(t, true, errors.Is(err, ErrBatchSizeIsZero))
		assert.Equal(t, [][]models.Suggestion(nil), result)
	}
	{
		result, err := SplitToBulks(dataNil, 1)
		assert.Equal(t, true, errors.Is(err, ErrSliceIsNil))
		assert.Equal(t, [][]models.Suggestion(nil), result)
	}
}
