package dp

import (
	"fmt"
	"testing"
)

//
//给你一个整数数组 nums 。玩家 1 和玩家 2 基于这个数组设计了一个游戏。
//
//玩家 1 和玩家 2 轮流进行自己的回合，玩家 1 先手。开始时，两个玩家的初始分值都是 0
//。每一回合，玩家从数组的任意一端取一个数字（即，nums[0] 或 nums[nums.length - 1]），
//取到的数字将会从数组中移除（数组长度减 1 ）。玩家选中的数字将会加到他的得分上。当数组中没有剩余数字可取时，游戏结束。
//
//如果玩家 1 能成为赢家，返回 true 。如果两个玩家得分相等，同样认为玩家 1 是游戏的赢家，
//也返回 true 。你可以假设每个玩家的玩法都会使他的分数最大化。



// 转动态规划
//注意dfs 条件
//两个dp 表
//在表1 先手  表2 后手
//l>r 不存在
//表一l=r dp[l][r]=nums[l]
//表2 l=r dp[l][r]=0
//dp1[i][j]=max(nums[l]+dp2[l+1][j],nums[r]+dp2[l][j-1])
//dp2[i][j]=min(dp1[i+1][j],dp1[i][j-
// 动态规划 -1都省了
// 1 5 2              0,1      1,1  0 0
//  1  *  *       0   *  *
//  *  5  *       *   0
//  *  *  2       *      0
func PredictTheWinner2(nums []int) bool {
	dp1:=make([][]int,len(nums))
	dp2:=make([][]int,len(nums))
	for i:=0;i<len(nums);i++{
		dp1[i]=make([]int,len(nums))
		dp2[i]=make([]int,len(nums))
		dp1[i][i]=nums[i]
	}
	// i=r j=l
	for i:=1;i<len(nums);i++{
		col,row:=i,0
		for col<len(nums){
			dp1[row][col]=max(nums[row]+dp2[row+1][col],nums[col]+dp2[row][col-1])
			dp2[row][col]=min(dp1[row+1][col],dp1[row][col-1])
			col++
			row++
		}


	}


	fmt.Println(dp1)
	fmt.Println(dp2)
	resulta:=dp1[0][len(nums)-1]
	resultb:=dp2[0][len(nums)-1]
	if resulta>resultb{
		return true
	}
	return false

}





// 记忆化搜索方式
func PredictTheWinner1(nums []int) bool {
	a1:=make([][]int,len(nums))
	b2:=make([][]int,len(nums))
	for i:=0;i<len(nums);i++{
		a1[i]=make([]int,len(nums))
		b2[i]=make([]int,len(nums))
		for j:=0;j<len(nums);j++{
			a1[i][j]=-1
			b2[i][j]=-1
		}
	}

	fmt.Println(a1,b2)

	a:=predfs(nums,0,len(nums)-1,a1,b2)
	b:=afterdfs(nums,0,len(nums)-1,a1,b2)
	fmt.Println(a,b)
	if a>=b{
		return true
	}
	return false
}

// 先手 先手模式下可以取的最多数值
func predfs(nums []int,l,r int,dp1,dp2 [][]int) int{
	if dp1[l][r]!=-1{
		return dp1[l][r]
	}

	if l==r{
		return nums[l]
	}
	ans:=0
	p:=nums[l]+afterdfs(nums,l+1,r,dp1,dp2)
	q:=nums[r]+afterdfs(nums,l,r-1,dp1,dp2)
	fmt.Println(nums[l:r],p,q)
	ans=max(p,q)
	dp1[l][r]=ans
	return ans
}
// 后手 后手模式下可以取的最多数值
func afterdfs(nums []int,l,r int,dp1,dp2 [][]int)int{
	if dp2[l][r] !=-1{
		return dp2[l][r]
	}

	if l==r{
		return 0
	}
	ans:=0
	p:=predfs(nums,l+1,r,dp1,dp2)
	q:=predfs(nums,l,r-1,dp1,dp2)
	ans=min(p,q)
	dp2[l][r]=ans
	return ans

}


func max(a,b int)int {
	if a>b{
		return a
	}
	return b
}


func min(a,b int )int{
	if a>b{
		return b
	}
	return a
}


func TestAaad(t *testing.T){
	nums:=[]int{1,5,233,7,}
	res:=PredictTheWinner2(nums)
	fmt.Println(res)
}