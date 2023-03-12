package leetcode

//给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
//子数组 是数组中的一个连续部分。


// 前缀最大和 = s[i]-s[j] 最大  s[j] -s[i]最小
func maxSubArray(nums []int) int {
	// 前缀和
	num:=make([]int,len(nums)+1)
	// 前缀和最小值,基于前缀和构造
	minList:=make([]int,len(nums)+1)
	n:=len(nums)
	for k,v:=range nums{
		num[k+1]  = num[k]+v

	}
	minList[0] = num[0]
	for i:=1;i<n+1;i++{
		minList[i] = min(minList[i-1],num[i])

	}


	res:=-10000
	for i:=1;i<n+1;i++{
		tmp :=num[i]-minList[i-1]
		res = max(res,tmp)
	}
	return res
}

func min(a,b int)int{
	if a>b{
		return b
	}
	return a
}


func max1(a,b int)int{
	if a<b{
		return b
	}
	return a
}


// 方法二 动态规划
// 滚动数组
func maxSubArray2(nums []int) int {
	res:=nums[0]
	//滚动更新数组
	for i:=1;i<len(nums)+1;i++{
		if nums[i]+nums[i-1]>nums[i]{
			nums[i] +=nums[i-1]
		}
		if nums[i]>res{
			res = nums[i]
		}
	}
	return res
}