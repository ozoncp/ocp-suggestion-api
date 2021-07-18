package utils

//Обратный ключ: конвертация исходного отображения (“ключ-значение“) in_map
//в отображение (“значение-ключ“) out_map
func ConvertMap(in_map map[string]string) (map[string]string, error) {
	var out_map = make(map[string]string, len(in_map))

	for key, value := range in_map {
		out_map[value] = key
	}

	return out_map, nil
}
