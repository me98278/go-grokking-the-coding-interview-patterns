package problem0560

import (
	"fmt"
	"strconv"
)

func SubArraySum(nums []int, k int) int {

	//create a for loop with length equal to size of passed in array minus k
	//loop thru the passed in array and read the first k elements
	//calculate the sum
	//save the sum in an max sum variable
	//move to the next element.
	//calculate the sum and deduct the prev element
	//compare this sum with the currently saved value
	//if the current sum is greater than prev value , then update the max sum variable

	maxSum := 0
	loopLength := len(nums)
	currentSum := 0

	for i := 0; i < loopLength; i++ {
		if i < k {
			fmt.Println(" Adding to currentSum " + strconv.FormatInt(int64(nums[i]), 10))
			currentSum = currentSum + nums[i]
			continue
		}

		if currentSum > maxSum {
			maxSum = currentSum
			fmt.Println(" maxSum " + strconv.FormatInt(int64(maxSum), 10))
		}
		fmt.Println(" Adding " + strconv.FormatInt(int64(nums[i]), 10))
		fmt.Println(" Removing " + strconv.FormatInt(int64(nums[i-k]), 10))
		currentSum = currentSum + nums[i] - nums[i-k]
		fmt.Println(" CurrentSum " + strconv.FormatInt(int64(currentSum), 10))

	}
	return maxSum
}
