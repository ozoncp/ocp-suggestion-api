package utils

//Фильтрация слайса in_slice по входному списку:
//по критерию отсутствия элемента в фильтре in_filter (захардкожен)
func FilterSlice(in_slice []int) ([]int, error) {
	var in_filter = map[int]struct{}{ //Захардкоженный фильтр
		1: {},
		2: {},
		3: {},
		4: {},
		5: {}}

	out_slice := make([]int, 0)
	for _, elem := range in_slice {
		if _, found := in_filter[elem]; !found {
			out_slice = append(out_slice, elem)
		}
	}

	return out_slice, nil
}
