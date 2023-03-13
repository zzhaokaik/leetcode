package qianzhuihe

import (
	"fmt"
	"testing"
)

//给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。
func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		// 某一子数组为前缀和i- 前缀和j，判断i,j是否存在
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}
	return count
}

func TestSubSum(t *testing.T){
	a:=[]int{1,1,1,1}
	res:=subarraySum(a,2)
	fmt.Println(res)
}