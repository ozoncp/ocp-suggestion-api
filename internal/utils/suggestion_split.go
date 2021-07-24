package utils

import (
	"errors"
	"github.com/ozoncp/ocp-suggestion-api/internal/models"
)

//ErrBatchSizeIsZero показывает, что размер батча равен нулю
var ErrBatchSizeIsZero = errors.New("batch size must be greater than zero")

//ErrSliceIsNil показывает, что слайс нулевой или его размер равен нулю
var ErrSliceIsNil = errors.New("slice cannot be nil or zero length")

//SplitToBulks разделяет слайс Suggestion на слайс слайсов с заданным размером батча
func SplitToBulks(suggestions []models.Suggestion, batchSize uint) ([][]models.Suggestion, error) {
	if suggestions == nil || len(suggestions) == 0 {
		return nil, ErrSliceIsNil
	}
	if batchSize == 0 {
		return nil, ErrBatchSizeIsZero
	}
	length := uint(len(suggestions))

	var outSlice [][]models.Suggestion
	for beg := uint(0); beg < length; beg += batchSize {
		end := beg + batchSize
		if end > length {
			end = length
		}
		outSlice = append(outSlice, suggestions[beg:end])
	}

	return outSlice, nil
}
