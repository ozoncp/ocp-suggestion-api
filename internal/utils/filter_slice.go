package utils

//Фильтрация слайса in_slice по входному списку:
//по критерию отсутствия элемента в фильтре in_filter (захардкожен)
func FilterSlice(in_slice []int) ([]int, error) {
	var in_filter = []int{1, 2, 3, 4, 5} //Захардкоженный фильтр

	isFound := func(key int) bool {
		for _, elem := range in_filter {
			if elem == key {
				return true
			}
		}
		return false
	}

	out_slice := make([]int, 0)
	for _, elem := range in_slice {
		if !isFound(elem) {
			out_slice = append(out_slice, elem)
		}
	}

	return out_slice, nil
}
