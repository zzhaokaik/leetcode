package leetcode

import (
	"fmt"
	"testing"
)



func containsNearbyDuplicate(nums []int, k int) bool {
	left,right:=0,0
	n:=len(nums)
	for right<n{
		for left<right{
			if nums[left]==nums[right] &&right!=left {
				if abs(left,right)<=k{
					return true
				}
				left++
			}
			right++
		}

	}
	return false
}

	func abs(a,b int) int{
		c:=a-b
		if c<=0{
		return 0-c
	}
		return c
	}
func TestList1(t *testing.T){
	listA:=[]int{1,0,1,1}
	res:=containsNearbyDuplicate(listA,1)
	fmt.Println(res)
}