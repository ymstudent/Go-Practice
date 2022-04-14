package main

func main() {

}

//青蛙调台阶问题，与求斐波那契数列第 n 项的值一样，只是初始值不同而已
// f(0) = 1, f(1) = 1, f(2) = 2
func numWays(n int) int {
	// 注意：此时的sum算出来时f(n+1),所以最后需要返回a
	a, b, sum := 1, 1, 0
	for i := 0; i < n; i++ {
		// f(n+1) = f(n-1) + f(n)
		sum = (a + b) % 1000000007
		// 将f(n)赋值给a
		a = b
		b = sum
	}
	return a
}

// f(0) = 0, f(1) = 1, f(2) = 1
func fib(n int) int {
	a, b, sum := 0, 1, 0
	for i := 0; i < n; i++ {
		sum = (a + b) % 1000000007
		a = b
		b = sum
	}
	return a
}
