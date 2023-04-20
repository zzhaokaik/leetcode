package dp
//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
//子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

func lengthOfLIS(nums []int) int {
	if len(nums)==0{
		return 0
	}
	// dp i nums[i]结尾最大子序列个数
	dp:=make([]int,len(nums))
	dp[0]=1
	for i:=1;i<len(nums);i++{
		dp[i]=1
		for j:=0;j<i;j++{
			// i 的值比子序列最大的大
			if nums[i]>nums[j]{
				dp[i]=max(dp[i],dp[j]+1)
			}

		}
	}
	res:=0
	for _,v:=range dp{
		res=max(v,res)
	}
	return res
}

//func max(a,b int)int{
//	if a>b{
//		return a
//	}
//	return b
//}