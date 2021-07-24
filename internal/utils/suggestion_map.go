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
	if suggestions == nil || len(suggestions) == 0 {
		return nil, ErrSliceIsNil
	}

	outMap := make(map[uint64]models.Suggestion, len(suggestions))
	for _, elem := range suggestions {
		if _, found := outMap[elem.Id]; found {
			return nil, fmt.Errorf("%w: elem1: %v, elem2: %v",
				ErrDuplicateKey, outMap[elem.Id], elem)
		}
		outMap[elem.Id] = elem
	}

	return outMap, nil
}
