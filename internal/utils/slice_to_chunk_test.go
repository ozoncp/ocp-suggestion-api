package utils

import "testing"

var s1 = make([]int, 4500)
var s2 = make([]int, 50000)
var s3 = make([]int, 150000)

//DoBenchmarksSlice вспомогательная функция для бенчмарков BenchmarkSliceToChunk/BenchmarkSliceToChunkMake
func DoBenchmarksSlice(f func(in_slice []int, chunk_size int) ([][]int, error)) {
	_, _ = f(s1, 20)
	_, _ = f(s2, 20)
	_, _ = f(s3, 20)
	_, _ = f(s1, 200)
	_, _ = f(s2, 200)
	_, _ = f(s3, 200)
	_, _ = f(s1, 1000)
	_, _ = f(s2, 1000)
	_, _ = f(s3, 1000)
}

//BenchmarkSliceToChunk бенчмарк для SliceToChunk
func BenchmarkSliceToChunk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoBenchmarksSlice(SliceToChunk)
	}
	b.ReportAllocs()
}

//BenchmarkSliceToChunkMake бенчмарк для SliceToChunkMake
func BenchmarkSliceToChunkMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoBenchmarksSlice(SliceToChunkMake)
	}
	b.ReportAllocs()
}
