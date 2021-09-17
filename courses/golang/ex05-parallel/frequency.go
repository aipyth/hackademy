package letter

import (
	"sync"
)

func Frequency(text string) map[rune]uint {
	frequencies := make(map[rune]uint)
	for _, letter := range text {
		frequencies[letter]++
	}
	return frequencies
}

func ConcurrentFrequency(texts []string) map[rune]uint {
	var wg sync.WaitGroup
	results := make([]map[rune]uint, len(texts))
	for index, text := range texts {
		wg.Add(1)
		go func(index int, text string) {
			results[index] = Frequency(text)
			wg.Done()
		}(index, text)
	}
	wg.Wait()
	return mergeMaps(results...)
}

func mergeMaps(maps ...map[rune]uint) map[rune]uint {
	result := make(map[rune]uint)
	for _, m := range maps {
		for key, value := range m {
			result[key] += value
		}
	}
	return result
}
