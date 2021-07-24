package utils

import "testing"

//BenchmarkSliceToChunk бенчмарк для SliceToChunk
func BenchmarkSliceToChunk(b *testing.B) {
	const size_data = 10000000
	const chunk_size = 10
	data := make([]int, 0, size_data)
	b.ResetTimer() // remove time of allocation
	for i := 0; i < b.N; i++ {
		for j := range data { // prepare data
			data[j] = i ^ j ^ 0x2cc
		}
		b.StartTimer()
		SliceToChunk(data, chunk_size)
		b.StopTimer()
	}
	b.ReportAllocs()
}

//BenchmarkSliceToChunkMake бенчмарк для SliceToChunkMake
func BenchmarkSliceToChunkMake(b *testing.B) {
	const size_data = 10000000
	const chunk_size = 10
	data := make([]int, 0, size_data)
	b.ResetTimer() // remove time of allocation
	for i := 0; i < b.N; i++ {
		for j := range data { // prepare data
			data[j] = i ^ j ^ 0x2cc
		}
		b.StartTimer()
		SliceToChunkMake(data, chunk_size)
		b.StopTimer()
	}
	b.ReportAllocs()
}
