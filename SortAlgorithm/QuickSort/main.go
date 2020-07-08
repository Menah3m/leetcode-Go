
package main

import (
	"fmt"
)



func QuickSort(nums []int ) []int{

	res := make([]int, 0)
	left := make([]int, 0)
	right := make([]int, 0)
	end := len(nums)-1
	pivot := nums[end]
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < pivot{
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}
	res = append(res, left...)
	res = append(res, pivot)
	res = append(res, right...)
	return res

}

func main(){
	nums := []int{1,3,5,7,0,9,2}
	res := QuickSort(nums)
	fmt.Println(res)

}