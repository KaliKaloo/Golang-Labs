package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"testing"
)

func random(n int) []int32 {
	s := make([]int32, n)
	for i := 0; i < n; i++ {
		s[i] = rand.Int31()
	}
	return s
}

const (
	start = 1048576
	end   = 4194304
)

func BenchmarkSequential(b *testing.B) {
	for size := start; size <= end; size *= 2 {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			os.Stdout = nil // Disable all program output apart from benchmark results
			for i := 0; i < b.N; i++ {
				unsorted := random(size)
				b.StartTimer()
				mergeSort(unsorted)
				b.StopTimer()
			}
		})
	}
}

func Benchmark512(b *testing.B) {
	for size := start; size <= end; size *= 2 {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			os.Stdout = nil // Disable all program output apart from benchmark results
			for i := 0; i < b.N; i++ {
				unsorted := random(size)
				b.StartTimer()
				parallelMergeSort(unsorted, 512)
				b.StopTimer()
			}
		})
	}
}

func Benchmark1024(b *testing.B) {
	for size := start; size <= end; size *= 2 {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			os.Stdout = nil // Disable all program output apart from benchmark results
			for i := 0; i < b.N; i++ {
				unsorted := random(size)
				b.StartTimer()
				parallelMergeSort(unsorted, 1024)
				b.StopTimer()
			}
		})
	}
}

func Benchmark2048(b *testing.B) {
	for size := start; size <= end; size *= 2 {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			os.Stdout = nil // Disable all program output apart from benchmark results
			for i := 0; i < b.N; i++ {
				unsorted := random(size)
				b.StartTimer()
				parallelMergeSort(unsorted, 2048)
				b.StopTimer()
			}
		})
	}
}

func Benchmark4096(b *testing.B) {
	for size := start; size <= end; size *= 2 {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			os.Stdout = nil // Disable all program output apart from benchmark results
			for i := 0; i < b.N; i++ {
				unsorted := random(size)
				b.StartTimer()
				parallelMergeSort(unsorted, 4096)
				b.StopTimer()
			}
		})
	}
}
//
func Benchmark8192(b *testing.B) {
	for size := start; size <= end; size *= 2 {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			os.Stdout = nil // Disable all program output apart from benchmark results
			for i := 0; i < b.N; i++ {
				unsorted := random(size)
				b.StartTimer()
				parallelMergeSort(unsorted, 8192)
				b.StopTimer()
			}
		})
	}
}
//
func Benchmark32768(b *testing.B) {
	for size := start; size <= end; size *= 2 {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			os.Stdout = nil // Disable all program output apart from benchmark results
			for i := 0; i < b.N; i++ {
				unsorted := random(size)
				b.StartTimer()
				parallelMergeSort(unsorted, 32768)
				b.StopTimer()
			}
		})
	}
}

func Benchmark16384(b *testing.B) {
	for size := start; size <= end; size *= 2 {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			os.Stdout = nil // Disable all program output apart from benchmark results
			for i := 0; i < b.N; i++ {
				unsorted := random(size)
				b.StartTimer()
				parallelMergeSort(unsorted, 16384)
				b.StopTimer()
			}
		})
	}
}

func BenchmarkDivGOMAX(b *testing.B) {
	for size := start; size <= end; size *= 2 {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			os.Stdout = nil // Disable all program output apart from benchmark results
			for i := 0; i < b.N; i++ {
				unsorted := random(size)
				b.StartTimer()
				parallelMergeSort(unsorted, size/runtime.NumCPU())
				b.StopTimer()
			}
		})
	}
}
