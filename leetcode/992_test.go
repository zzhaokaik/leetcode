package leetcode

import "testing"

//
//992. K 个不同整数的子数组
//困难
//408
//相关企业
//给定一个正整数数组 nums和一个整数 k ，返回 num 中 「好子数组」 的数目。
//
//如果 nums 的某个子数组中不同整数的个数恰好为 k，则称 nums 的这个连续、不一定不同的子数组为 「好子数组 」。
//
//例如，[1,2,3,1,2] 中有 3 个不同的整数：1，2，以及 3。

// 前缀和+滑动窗口
//求k 个不同整数子数组  --->  求至多k 不同整数子数组 -求至多 k-1 不同整数子数组
//前缀和 求前n 个数组的 不同子数组
// sum 子数组个数    n  长度
//sum[i] = sum[i-1]+len(n)
//[1]
//1           1
//[1,2]
//1 1,2        2
//[1,2,3]
//1 1,2 1,2,3 2,3 3   5


func subarraysWithKDistinct(nums []int, k int) int {
	return  knums(nums,k)-knums(nums,k-1)
}
// 统计子数组 前缀和
// sum[i] = sum[i-1]+ 区间长度（len(nums)）

// 返回至少k个子数组个数
func knums(nums []int,k int )int{
	res:=0
	count:=0
	i,j:=0,0
	tmp:=make(map[int]int)
	for j<len(nums){
		r:=nums[j]
		j++
		tmp[r]++
		if tmp[r]==1{
			count++
		}
		for count>k{
			l:= nums[i]
			i++
			tmp[l]--
			if tmp[l]==0{
				count--
			}
		}
		res +=j-i
	}
	return res
}


func TestGetK(t *testing.T){
	k:=2
	nums:=[]int{
		1,2,2,3,4,
	}
	subarraysWithKDistinct(nums,k)
}