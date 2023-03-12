package leetcode

import (
	"fmt"
	"math"
	"testing"
)

// 双指针+简化前缀和

//209. 长度最小的子数组
//中等
//1.5K
//相关企业
//给定一个含有 n 个正整数的数组和一个正整数 target 。
//
//找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，
//并返回其长度。如果不存在符合条件的子数组，返回 0 。
// 99 1 2 3
func minSubArrayLen(target int, nums []int) int {
	res:=math.MaxInt64
	sum:=0
	l,r:=0,0
	for r<=len(nums)-1{
		sum+=nums[r]
		for sum >=target{
			res=min(res,r-l+1)
			sum-=nums[l]
			l++
		}
		r++
	}
	if res==math.MaxInt64{
		res=0
	}
	return res
}

func min00(a,b int) int{
	if a>b{
		return b
	}
	return a
}

func Testminsub11(t *testing.T){
	nums:=[]int{2,3,1,2,4,3,}
	target:=7
	res:=minSubArrayLen(target,nums)
	fmt.Println(res)
}
