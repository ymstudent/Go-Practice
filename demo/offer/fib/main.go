package main

func main() {
	println(fib(10))
}

const mod int = 1e9 + 7

type matrix [2][2]int

// 矩阵相乘
func multiply(a, b matrix) (c matrix) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = (a[i][0]*b[0][j] + a[i][1]*b[1][j]) % mod
		}
	}
	return
}

// 矩阵快速幂: 求底数为a 幂次为n 的结果
func pow(a matrix, n int) matrix {
	// 定义变量，存储计算结果，此次定义为单位阵
	ret := matrix{{1, 0}, {0, 1}}
	// 对幂次进行整除
	for ; n > 0; n >>= 1 {
		// 如果为奇数，需要多乘一次底数
		if n&1 == 1 {
			ret = multiply(ret, a)
		}
		a = multiply(a, a)
	}
	return ret
}

// 斐波那切数
func fib(n int) int {
	if n < 2 {
		return n
	}

	// 定义乘积底数
	base := matrix{{1, 1}, {1, 0}}
	// 定义幂次
	power := n - 1

	res := pow(base, power)
	return res[0][0]
}
