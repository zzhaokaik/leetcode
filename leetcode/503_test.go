package leetcode

import (
	"fmt"
	"testing"
)
// 单调栈 下一个最大数字


func nextGreaterElements(nums []int) []int {
	// tmp 单调栈存放 nums index
	// 当比较大入栈
	tmp:=make([]int,0)
	res:=make([]int,len(nums))
	for i:=0;i<len(nums);i++{
		res[i]=-1
	}

	for i:=0;i<len(nums)*2 ;i++{
		key:=i%len(nums)
		for len(tmp)>0&& nums[key] > nums[tmp[len(tmp)-1]]{
			res[tmp[len(tmp)-1]] = nums[key]
			tmp = tmp[:len(tmp)-1]
		}
		tmp=append(tmp,key)
	}

	return res
}


func TestNext(t *testing.T){
	testList:=[]int{1,2,1,}
	res:=nextGreaterElements(testList)
	fmt.Println(res)
}
