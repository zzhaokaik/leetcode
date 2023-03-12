package leetcode

import (
	"fmt"
	"testing"
)

//单调栈
//
//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
//求在该柱状图中，能够勾勒出来的矩形的最大面积
// 方便计算左右补0
func largestRectangleArea(heights []int) int {
	var maxArea int
	var tmp []int
	// 方便计算左右补0
	newheight:=make([]int,len(heights)+2)
	for i:=0;i<len(heights);i++{
		newheight[i+1] = heights[i]
	}

	for k,v:=range newheight{
		for  len(tmp)>0 && v<newheight[tmp[len(tmp)-1]]{
			index := tmp[len(tmp)-1]

			tmp = tmp[:len(tmp)-1]
			h:=newheight[index]
			x:= k-tmp[len(tmp)-1]-1
			area:=h*x
			maxArea = max(maxArea,area)

		}
		tmp = append(tmp,k)
	}

	return maxArea
}

func max(a,b int)int {
	if a>b{
		return a
	}
	return b
}


func TestMaxarea(t *testing.T){
	listA:=[]int{1,1}
	res:=largestRectangleArea(listA)
	fmt.Println(res)
}
