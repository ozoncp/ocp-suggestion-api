package utils

import "errors"

//ConvertMap конвертирует исходное отображение (“ключ-значение“) in_map в отображение (“значение-ключ“) out_map
func ConvertMap(in_map map[string]string) (map[string]string, error) {
	if in_map == nil {
		return nil, errors.New("Map cannot be nil")
	}

	out_map := make(map[string]string, len(in_map))
	for key, value := range in_map {
		out_map[value] = key
	}

	return out_map, nil
}
