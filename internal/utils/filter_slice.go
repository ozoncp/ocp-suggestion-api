package utils

import "errors"

//Захардкоженный фильтр
var in_filter = map[int]struct{}{
	1: {},
	2: {},
	3: {},
	4: {},
	5: {},
}

//FilterSlice фильтрует слайс in_slice по входному списку:
//по критерию отсутствия элемента в фильтре in_filter (захардкожен - см. выше)
func FilterSlice(in_slice []int) ([]int, error) {
	if in_slice == nil {
		return nil, errors.New("Slice cannot be nil")
	}

	//Аллоцируем исходный capacity/2 для уменьшения количества будущих аллокаций
	out_slice := make([]int, 0, len(in_slice)/2)
	for _, elem := range in_slice {
		if _, found := in_filter[elem]; !found {
			out_slice = append(out_slice, elem)
		}
	}

	return out_slice, nil
}
