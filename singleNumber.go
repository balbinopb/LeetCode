package main

import "fmt"

func singleNumber(nums []int) int {
    result := 0
    for i:=0; i<len(nums); i++ {
        result ^= nums[i]
    }
    return result
}

func main(){
	nums:= []int {4,1,2,1,2}
	fmt.Println(singleNumber(nums))
}