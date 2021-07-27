package utils

import (
	"errors"
	"fmt"
	"github.com/ozoncp/ocp-suggestion-api/internal/models"
)

//ErrDuplicateKey показывает, что при создании map найден дублирующий ключ
var ErrDuplicateKey = errors.New("duplicate key found")

//ConvertSliceToMap конвертирует слайс от структуры Suggestion в отображение,
//где ключ - идентификатор структуры, а значение - сама структура
func ConvertSliceToMap(suggestions []models.Suggestion) (map[uint64]models.Suggestion, error) {
	if len(suggestions) == 0 {
		return nil, ErrSliceIsNil
	}

	outMap := make(map[uint64]models.Suggestion, len(suggestions))
	for _, elem := range suggestions {
		if _, found := outMap[elem.ID]; found {
			return nil, fmt.Errorf("elem1: %v, elem2: %v : %w",
				outMap[elem.ID], elem, ErrDuplicateKey)
		}
		outMap[elem.ID] = elem
	}

	return outMap, nil
}
