package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}

	elements := make([]int, size)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < size; i++ {
		elements[i] = r.Int()
	}

	return elements
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	if len(data) == 1 {
		return data[0]
	}

	max := slices.Max(data)

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var wg sync.WaitGroup

	if len(data) < CHUNKS {
		return maximum(data)
	}

	chunkSize := (len(data) + CHUNKS - 1) / CHUNKS
	maxValues := make([]int, CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		if start >= len(data) {
			continue
		}

		end := start + chunkSize
		if end > len(data) {
			end = len(data)
		}

		wg.Add(1)
		chunk := data[start:end]

		go func(i int, chunk []int) {
			defer wg.Done()
			maxValues[i] = maximum(chunk)
		}(i, chunk)
	}

	wg.Wait()
	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	elements := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(elements)
	elapsed := time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(elements)
	elapsed = time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
