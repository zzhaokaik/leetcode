package offer

import (
	"fmt"
	"testing"
)

// 生产一个 波菲纳妾数的第 k 位
//1.使用矩阵将波菲纳妾数 优化为计算M的k次幂
//1.使用快速幂算法将将复杂度降低为 nlogn
const mod int = 1e9 + 7

type matrix [2][2]int

func multiply(a, b matrix) (c matrix) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = (a[i][0]*b[0][j] + a[i][1]*b[1][j]) % mod
		}
	}
	return
}

func pow(a matrix, n int) matrix {
	ret := matrix{{1, 0}, {0, 1}}
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			ret = multiply(ret, a)
		}
		a = multiply(a, a)
	}
	return ret
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	res := pow(matrix{{1, 1}, {1, 0}}, n-1)
	return res[0][0]
}

func TestFib(t *testing.T){
	a:=1
	b:=fib(a)
	fmt.Println(b)
}