package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	j := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}

func main() {
	arr := []int{1, 1, 2}

	Length := removeDuplicates(arr)

	// Print the results
	fmt.Println(arr[:Length])
	fmt.Print( Length)

}
