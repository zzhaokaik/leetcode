package leetcode

import (
	"fmt"
	"math"
	"testing"
)

// 快速排序
// 两种 维护一个pivot 用于分组分治
// a 分治点在结尾
// b 分治点在中间

// 快速排序1
func sort1(list []int) []int {
	quicksort1(list, 0, len(list)-1)
	return list
}

func quicksort1(nums []int, l, r int) {
	if l > r {
		return
	}
	p := partition1(nums, l, r)
	quicksort1(nums, l, p-1)
	quicksort1(nums, p+1, r)
}

func partition1(nums []int, l, r int) int {
	pivot := nums[r]
	i, j := l, l
	fmt.Println(nums,i,j,l,r,pivot)
	for j < r {
		if nums[j] < pivot {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
		j++
		fmt.Println(nums,i,j,l,r,pivot)
	}
	nums[i], nums[r] = nums[r], nums[i]
	fmt.Println(nums,i,j,l,r,pivot)
	return i
}

func TestSort(t *testing.T) {

	nums := []int{
		11, 5, 3, 4, 1, 2,
	}
	sort1(nums)
	fmt.Println(nums)
}

// 快速排序2

// 快速排序2
func sort2(list []int) []int {
	quicksort2(list, 0, len(list)-1)
	return list
}

func quicksort2(nums []int, l, r int) {
	if l > r {
		return
	}
	p := partition2(nums, l, r)
	quicksort2(nums, l, p-1)
	quicksort2(nums, p+1, r)
}

func partition2(nums []int, l, r int) int {

	pivot := nums[(l+r)/2]
	i, j := l, r

	for i < j {
		for nums[i] < pivot {
			i++
		}
		for nums[j] > pivot {
			j--
		}
		nums[j], nums[i] = nums[i], nums[j]

	}

	return i
}

func TestSort2(t *testing.T) {
	nums := []int{
		11, 5, 3, 4, 1, 2,
	}
	sort2(nums)
	fmt.Println(nums)
}

// 归并排序
// 将数组分为2 对数组进行排序
//和快速排序区别
//快速排序设置了默认区分值,快速排序
// 归并没有
func guibing(nums []int) []int {
	guibingmerge(nums, 0, len(nums)-1)
	return nums
}

func guibingmerge(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	guibingmerge(nums, l, mid)
	guibingmerge(nums, mid+1, r)
	merge(nums, l, mid, r)

}
func merge(nums []int, l, mid, r int) {
	left := append([]int(nil), nums[l:mid+1]...)
	right := append([]int(nil), nums[mid+1:r+1]...)
	left = append(left,math.MaxInt64)
	right =append(right,math.MaxInt64)
	i, j := 0, 0
	for k := l; k <= r; k++ {
		if left[i] <= right[j] {
			nums[k] = left[i]
			i++
		} else {
			nums[k] = right[j]
			j++
		}
	}
}

func TestSort3(t *testing.T) {
	nums := []int{
		11, 5, 3, 4, 1, 2,
	}
	guibing(nums)
	fmt.Println(nums)
}


//堆排序 桶排序