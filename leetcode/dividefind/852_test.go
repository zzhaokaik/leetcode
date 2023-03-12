package dividefind

import "testing"

// 符合下列属性的数组 arr 称为 山脉数组 ：
//arr.length >= 3
//存在 i（0 < i < arr.length - 1）使得：
//arr[0] < arr[1] < ... arr[i-1] < arr[i]
//arr[i] > arr[i+1] > ... > arr[arr.length - 1]
//给你由整数组成的山脉数组 arr ，返回任何满足 arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
//的下标 i 。


// 010
// 1
// 二分查找 左右闭模板 爬坡 当nums[i]>nums[i+1] i左边全部去掉
// 复杂度 logn
func peakIndexInMountainArray(arr []int) int {
	res:=0
	l,r:=0,len(arr)-1
	for l<r{
		mid:=(l+r)/2
		if arr[mid]>arr[mid+1]{
			res=mid
			r=mid
		}else{
			l=mid+1
		}
	}

	return res

}

func TestPeak(t *testing.T){
	nums:=[]int{1,2,4,3}
	peakIndexInMountainArray(nums)
}