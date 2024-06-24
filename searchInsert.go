package main

import "fmt"

func searchInsert(nums []int, target int) int {
    start := 0
    end := len(nums) - 1
    
    for start <= end {
        mid := (start + end) / 2
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            end = mid - 1
        } else {
            start = mid + 1
        }
    }
    
    return start
}

func main(){
	nums:= []int {1,3,5,6}
	target := 5

	fmt.Println(searchInsert(nums,target))
}