package dp

import (
	"fmt"
	"testing"
)

// 131 分割回文字

// 每个字符间有个逗号
// [aa] [b]  [a] [a] [b]

var res [][]string
func partition(s string) [][]string {
	res=make([][]string,0)
	dfs(s,0,[]string{})
	return res
}

// i 第i 个字符处
func dfs(s string,i int,path []string){
	if i==len(s){
		res = append(res,append([]string(nil),path...))
		return
	}
	for j:=i;j<len(s);j++{
		tmp := s[i:j+1]
		if isPart(tmp){
			path = append(path,tmp) //[a]
			dfs(s,j+1,path)         // [a,a]   //
			// 不明白
			path = path[:len(path)-1]
		}
	}
}

// aa >-a +a 也是
func isPart(nums string) bool{
	l:=0
	r:=len(nums)-1
	if len(nums)==0{
		return true
	}
	for l<r{
		if nums[l]!=nums[r]{
			return false
		}
		l++
		r--
	}
	return true
}

func TestPartest(t *testing.T){
	a:="aab"
	res:=partition(a)
	fmt.Println(res)
}