package main

import (
	"fmt"
	problem0560 "go-grokking-the-coding-interview-patterns/sliding_window/0560.subarray-sum-equals-k"
	"strconv"
)

func main() {
	fmt.Println("testing problem0560.SubarraySum")
	nums := []int{2, 1, 5, 1, 3, 2}
	k := 3
	result := problem0560.SubArraySum(nums, k)
	fmt.Printf("result from  problem0560.SubarraySum is :" + strconv.FormatInt(int64(result), 10))
}
