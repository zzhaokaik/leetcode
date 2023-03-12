package offer

import (
	"fmt"
	"testing"
)

//
//地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
//一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），
//也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，
//因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

func movingCount(m int, n int, k int) int {
	visit:=make([][]bool,0)
	for i := 0; i < m; i++ {
		list:=make([]bool,n)
		visit = append(visit,list)
	}

	return dfs(0,0,m,n,k,visit)
}

func dfs(x, y int, m, n, k int, visit [][]bool) int {
	if x < 0 || y < 0 || x > m || y > n || visit[x][y] == true {
		return 0

	}
	if getNum(x)+getNum(y)<k{
		visit[x][y]=true
		return 1+dfs(x+1,y,m,n,k,visit)+dfs(x,y+1,m,n,k,visit)+dfs(x-1,y,m,n,k,visit)+dfs(x,y-1,m,n,k,visit)
	}
	return dfs(x+1,y,m,n,k,visit)+dfs(x,y+1,m,n,k,visit)+dfs(x-1,y,m,n,k,visit)+dfs(x,y-1,m,n,k,visit)
}

func getNum(a int)int{
	b:=0
	for a>0{
		b+=a%10
		a /= 10
	}
	return b
}


func TestMoving(t *testing.T) {
	x := 1
	y := 1
	k := 1
	count := movingCount(x, y, k)
	fmt.Println(count)
}

func TestNums(t *testing.T) {
	nums := 12
	fmt.Println(nums % 10)
}
