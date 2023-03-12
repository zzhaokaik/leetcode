package weiyunsuan

import (
	"fmt"
	"testing"
)

//数组中只出现1次的数
//给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。


// 位运算 map方法省略
// 任何数字 亦或自己 都是0(都是1 则为0)
//


func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}



func TestSingle(t *testing.T){

	nums:=[]int{1,2,1}
	res:=singleNumber(nums)
	fmt.Println(res)
}

func TestWeitu(t *testing.T){
	a:=1
	b:=a^0
	fmt.Println(a,b)
}

func TestWeitu2(t *testing.T){
	a:=13
	b:=int32(a)>>1&1 // &1 相当于取最后1位
	c:=int32(a)
	fmt.Println(a,b,string(c))
	fmt.Printf("% 08b", c)
}