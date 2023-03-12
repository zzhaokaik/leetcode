package leetcode

// 同209
//给你一个整数数组 nums 和一个整数 k ，请你返回子数组内所有元素的乘积严格小于 k 的连续子数组的数目。

func numSubarrayProductLessThanK(nums []int, k int) int {
	l,r:=0,0
	res:=0
	sum:=1
	for r<len(nums){
		sum*=nums[r]
		for l<=r&&sum>k{
			sum/=nums[l]
			l++
		}
		res+=r-l+1
		r++
	}
	return res
}
