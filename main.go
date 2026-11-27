package main

import (
	"fmt"
	"math/rand"
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
		return []int{}
	}

	const minRandomValue = 1
	const maxRandomValue = 1_000_000
	var slice = make([]int, size)

	for index := range slice {
		slice[index] = rand.Intn(maxRandomValue) + minRandomValue
	}

	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]
	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	chunkSize := len(data) / CHUNKS
	chunkMaximum := make([]int, CHUNKS)

	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}

		go func(index int, chunk []int) {
			defer wg.Done()
			chunkMaximum[index] = maximum(chunk)
		}(i, data[start:end])
	}

	wg.Wait()

	return maximum(chunkMaximum)
}

func runSearch(title string, data []int, search func([]int) int) {
	fmt.Println(title)

	timeNow := time.Now()
	maxValue := search(data)
	elapsed := time.Since(timeNow).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", maxValue, elapsed)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n\n", SIZE)
	data := generateRandomElements(SIZE)

	runSearch("Ищем максимальное значение в один поток", data, maximum)
	fmt.Println()
	runSearch(fmt.Sprintf("Ищем максимальное значение в %d потоков", CHUNKS), data, maxChunks)
}
