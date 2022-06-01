package main

import (
	"fmt"
	problem0209 "go-grokking-the-coding-interview-patterns/sliding_window/0209.minimum-size-subarray-sum"
	"strconv"
)

func main() {
	fmt.Println("testing problem0560.SubarraySum")
	nums := []int{1, 2, 3, 4, 5}
	k := 11
	result := problem0209.MinSubArrayLen(k, nums)
	fmt.Println("result from  problem0209.MinSubArrayLen is :" + strconv.FormatInt(int64(result), 10))

	nums = []int{2, 3, 1, 2, 4, 3}
	k = 7
	result = problem0209.MinSubArrayLen(k, nums)
	fmt.Println("result from  problem0209.MinSubArrayLen is :" + strconv.FormatInt(int64(result), 10))
}
