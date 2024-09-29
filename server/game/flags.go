package game

import (
	"math/rand"
	"reflect"
)

// GetRandomFromList returns a random element from the given slice
func GetRandomFromList[T any](list []T) T {
	return list[GetRandomInt(0, len(list)-1)]
}

// GetRandomPairs generates random pairs from the given list
func GetRandomPairs[T any](quantity int, list []T, matchingFactor float64, label1, label2 string) []map[string]T {
	randomList := make([]map[string]T, quantity)
	for i := range randomList {
		randomList[i] = map[string]T{
			label1: GetRandomFromList(list),
			label2: GetRandomFromList(list),
		}
	}

	matchingLength := countMatching(randomList, label1, label2)

	for float64(matchingLength) < float64(quantity)*matchingFactor {
		value := GetRandomFromList(list)
		index := GetRandomInt(0, quantity-1)
		randomList[index] = map[string]T{
			label1: value,
			label2: value,
		}
		matchingLength = countMatching(randomList, label1, label2)
	}

	return randomList
}

// GetRandomInt returns a random integer between min and max (inclusive)
func GetRandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// countMatching counts the number of matching pairs in the list
func countMatching[T any](list []map[string]T, label1, label2 string) int {
	count := 0
	for _, v := range list {
		if reflect.DeepEqual(v[label1], v[label2]) {
			count++
		}
	}
	return count
}
