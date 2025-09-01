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
func generateRandomElements(size int) ([]int, error) {
	// ваш код здесь
	if size <= 0 {
		err := fmt.Errorf("wrong slice size on function input, size: %v", size)
		fmt.Println(err)
		return nil, err
	}

	result := make([]int, size)
	for i := range result {
		result[i] = rand.Int()
	}
	if len(result) != size {
		err := fmt.Errorf("wrong slice size on function output: %v", size)
		fmt.Println(err)
		return nil, err
	}
	return result, nil

}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		err := fmt.Errorf("slice is empty on input of maximum func")
		fmt.Println(err)
		return 0
	}
	maximum := data[0]
	for _, v := range data[1:] {
		if maximum < v {
			maximum = v
		}
	}
	return maximum
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	length := len(data)
	if length == 0 {
		err := fmt.Errorf("slice is empty on input of maxChunks func")
		fmt.Println(err)
		return 0
	}

	actualChunks := CHUNKS
	if length < actualChunks {
		actualChunks = length
	}

	maximums := make([]int, actualChunks)
	var wg sync.WaitGroup
	for i := 0; i < actualChunks; i++ {
		wg.Add(1)
		threadStart := i * length / actualChunks
		threadEnd := (i + 1) * length / actualChunks
		go func(i, start, end int) {
			defer wg.Done()
			maximums[i] = maximum(data[start:end])
		}(i, threadStart, threadEnd)
	}
	wg.Wait()
	return maximum(maximums)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	slice, _ := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(slice)
	stop := time.Now()
	elapsed := stop.Sub(start) / time.Microsecond
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(slice)
	stop = time.Now()
	elapsed = stop.Sub(start) / time.Microsecond
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
