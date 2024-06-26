package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums) + 1 // because we are missing one number
	totalSum := (n * (n - 1)) / 2
	arraySum := 0
	for i := 0; i < len(nums); i++ {
		arraySum += nums[i]
	}
	return totalSum - arraySum
}

func main() {
	arr := []int{3, 0, 1}

	fmt.Println(missingNumber(arr))

}
