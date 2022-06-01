package problem0209

func MinSubArrayLen(target int, nums []int) int {
	currentSum := 0
	windowLen := 0
	//create a slice to hold the current window
	m := make([]int, 1)
	for i := 0; i < len(nums); {
		if currentSum < target {
			m = append(m, nums[i])
			currentSum = currentSum + nums[i]
			i++
		}

		for currentSum >= target {
			//reduce the current sum by the first element
			currentSum = currentSum - m[0]
			//remove the first element from the slice
			m = removeIndex(m, 0)
			temp := len(m)
			if temp != 0 && currentSum >= target {
				windowLen = temp
			}
		}
	}

	//return the length of the map which fits
	return windowLen

}

func removeIndex(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
