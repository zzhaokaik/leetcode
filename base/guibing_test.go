package base

import (
	"fmt"
	"math"
	"testing"
)

func guibing(nums []int) []int{
	guibingdfs(nums,0,len(nums)-1)
	return nums
}


func guibingdfs(nums []int,l,r int ){
	if l>=r{
		return
	}
	mid :=(l+r)/2
	guibingdfs(nums,l,mid)
	guibingdfs(nums,mid+1,r)
	merge(nums,l,r,mid)



}


func merge(nums []int,l,r,mid int){
	left:=append([]int(nil),nums[l:mid+1]...)
	right:=append([]int(nil),nums[mid+1:r+1]...)
	left = append(left,math.MaxInt64)
	right =append(right,math.MaxInt64)
	i,j:=0,0
	fmt.Println(nums)
	for k:=l;k<=r;k++{
		fmt.Println(left,right,mid,l,r,k,nums)
		if right[i]<left[j]{
			nums[k]=right[i]
			i++
		}else{
			nums[k]=left[j]
			j++
		}
	}


}


func TestGGG(t *testing.T){

	nums:=[]int{11, 5, 3, 4, 1, 2,}
	guibing(nums)
	fmt.Println(nums)
}