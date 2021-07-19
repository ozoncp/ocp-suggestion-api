package utils

import "errors"

//SliceToChunk разделяет слайс на чанки:
//исходный слайс in_slice конвертируется в слайс слайсов out_slice с чанками одинкового размера (кроме последнего)
func SliceToChunk(in_slice []int, chunk_size int) ([][]int, error) {
	if in_slice == nil {
		return nil, errors.New("Slice cannot be nil")
	}
	if chunk_size <= 0 {
		return nil, errors.New("Chunk size must be greater than zero")
	}

	out_slice := make([][]int, 0)
	for beg := 0; beg < len(in_slice); beg += chunk_size {
		end := beg + chunk_size
		if end > len(in_slice) {
			end = len(in_slice)
		}
		out_slice = append(out_slice, in_slice[beg:end])
	}

	return out_slice, nil
}
