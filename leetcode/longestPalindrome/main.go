package main

func main() {
	s := "dabaabf"
	println(longestPalindrome2(s))
}

// 中心扩散法，边界1与边界2不容易理解，需要找个例子在纸上画一下，能画明白但说不出原理
func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		l1 := expandAroundCenter(s, i, i)
		l2 := expandAroundCenter(s, i, i+1) // 2, 3
		l := max(l1, l2)
		if l > end-start {
			//这里为什么要i-1?,这里说明一下,因为for循环是从0开始的,
			//如果是奇数回文,假设有个回文是3个,那么len=3,此时中心i是下标1(从0开始),那么(len-1)/2和len/2的结果都是1,因为整型会向下取整
			//但是如果是偶数回文,假设有个回文是4个,那么len=4,此时的中心是一条虚线,但是i的位置却在1,(因为S是从左向右遍历的,如果从右向左,
			//i的位置就会在2.)这时候,(len-1)/2=1,len/2=2.很明显为了保证下标正确,我们需要的是(len-1)/2.原因其实是i在中心线的左边一位,
			//所以要少减个1.
			start = i - (l-1)/2 //边界1
			end = i + l/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, l, r int) int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	// 为什么不是-2而是减1
	//这里其实是right-left+1-2,意思就是right-left+1是本来的长度,但是由于上面最后一次判断肯定false,所以最后一次left--和right++
	//其实不属于回文的一部分,所以要减去2
	return r - l - 1 //边界2
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 动态规划
// 状态转移方程：dp[i][j] =  (s[i]==s[j] && dp[i+1][j-1])。即如果之前的字符为回文且当前左右边界字符相等，那么当前字符也是回文
func longestPalindrome2(s string) string {
	l := len(s)
	if l < 2 {
		return s //单字符都是回文
	}
	start, maxLen := 0, 1
	dp := make([][]bool, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]bool, l)
	}

	// 从2开始枚举回文长度
	for length := 2; length <= l; length++ {
		// 从0开始枚举左边界
		for i := 0; i < l; i++ {
			// 计算右边界，由于length := r - i + 1，所以r := length + i - 1
			j := length + i - 1
			// 右边界越位，推出循环
			if j >= l {
				break
			}

			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				// 边界条件： (j-1) - (i+1) + 1 < 2 => j-i < 3
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				start = i
			}
		}
	}
	return s[start : start+maxLen]
}
