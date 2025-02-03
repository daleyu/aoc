package math_utils

func Max(nums ...int) int {
	result := nums[0]
	for _, num := range nums {
		if num > result {
			result = num
		}
	}
	return result
}

func MaxIndex(nums ...int) (int, int) {
	result := nums[0]
	index := 0
	for i, num := range nums {
		if num > result {
			result = num
			index = i
		}
	}

	return index, result
}
