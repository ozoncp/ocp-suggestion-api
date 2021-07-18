package utils

import "errors"

//Разделение на слайса на батчи:
//исходный слайс in_slice конвертируется в слайс слайсов out_slice с чанками одинкового размера (кроме последнего)
func SliceToChunk(in_slice []int, chunk_size int) ([][]int, error) {
	if chunk_size <= 0 {
		return nil, errors.New("Chunk size must be greater than zero")
	}
	if len(in_slice) == 0 {
		return make([][]int, 0), nil
	}

	num_chunks := (len(in_slice)-1)/chunk_size + 1
	out_slice := make([][]int, num_chunks)
	for i := 0; i < num_chunks-1; i++ {
		out_slice[i] = in_slice[i*chunk_size : (i+1)*chunk_size]
	}
	out_slice[num_chunks-1] = in_slice[(num_chunks-1)*chunk_size:]

	return out_slice, nil
}
