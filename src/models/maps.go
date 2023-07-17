package models

func CountOccurrences(input []string) map[string]int {
	occurrences := make(map[string]int)

	for _, str := range input {
		occurrences[str]++
	}

	return occurrences
}

func MergeMaps(map1, map2 map[string]int) map[string]int {
	mergedMap := make(map[string]int)

	for key, value := range map1 {
		mergedMap[key] = value
	}

	for key, value := range map2 {
		mergedMap[key] += value
	}

	return mergedMap
}
