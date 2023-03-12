package dividefind
// 寻找峰值
// 二分查找
// 二分查找 左右闭模板 爬坡 当nums[i]>nums[i+1] i左边全部去掉
// 复杂度 logn
func findPeakElement(nums []int) int {
	// r 峰顶以及封顶右侧范围
	// 爬坡 找第一个mid 》mid+1的
	l,r:=0,len(nums)-1

	for l<r{
		mid:=(l+r)/2
		if nums[mid]>nums[mid+1]{
			r=mid
		}else{
			l=mid+1
		}
	}
	return r

}
