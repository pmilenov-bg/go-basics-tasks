package models

func RemoveDuplicates(input []int) ([]int, []int) {
	uniqueSet := make(map[int]bool)
	repeated := make(map[int]bool)
	result, result1 := []int{}, []int{}

	for _, num := range input {
		if !uniqueSet[num] {
			uniqueSet[num] = true
			result = append(result, num)
		} else {
			if !repeated[num] {
				repeated[num] = true
				result1 = append(result1, num)
			}
		}
	}

	return result, result1
}
