package leetcode

import (
	"fmt"
	"testing"
)

//
//给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，
//下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。

// 遍历数组 维护中间栈 tmp 存放当前数组元素，遍历栈 如果找到比当前数组大的数据 元素出栈
// 由于有比自己大的就出，所以栈是单调递减的 叫做单调栈
// 73
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	var tmp []int
	for k, v := range temperatures {

		for len(tmp) > 0 && v > temperatures[tmp[len(tmp)-1]] {
			index := tmp[len(tmp)-1]
			res[index] = k - index
			tmp = tmp[:len(tmp)-1]
		}
		tmp = append(tmp, k)
	}
	return res
}

func TestDail(t *testing.T) {
	listA := []int{73, 74, 75, 71, 69, 72, 76, 73}
	res := dailyTemperatures(listA)
	fmt.Println(res)
}
