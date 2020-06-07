package Week_03

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	var result [][]int

	for i, num := range nums {
		newNums := make([]int, len(nums)-1)
		copy(newNums[0:], nums[0:i])
		copy(newNums[i:], nums[i+1:])
		nexts := permute(newNums)

		for _, next := range nexts {
			result = append(result, append(next, num))
		}
	}

	return result
}
