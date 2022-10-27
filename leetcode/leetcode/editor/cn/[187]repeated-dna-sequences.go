package cn

import (
	"math"
)

//leetcode submit region begin(Prohibit modification and deletion)
func findRepeatedDnaSequences(s string) []string {
	// 先把字符串转化成4进制的数字数组
	nums := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'A':
			nums[i] = 0
		case 'C':
			nums[i] = 1
		case 'G':
			nums[i] = 2
		case 'T':
			nums[i] = 3
		}
	}
	// 记录结果
	res := make(map[int]string, 0)
	// 记录重复出现的哈希值
	seen := make(map[int]struct{})

	// 数字位数
	L := 10
	// 进制
	R := 4
	// 存储 R^(L - 1) 的结果
	RL := math.Pow(float64(R), float64(L-1))
	// 维护滑动窗口中字符串的哈希值
	windowHash := 0
	// 滑动窗口代码框架，时间 O(N)
	left, right := 0, 0

	for right < len(nums) {
		// 扩大窗口，移入字符，并维护窗口哈希值（在最低位添加数字）
		windowHash = R*windowHash + nums[right]
		right++
		// 当子串的长度达到要求
		if right-left == L {
			// 根据哈希值判断是否曾经出现过相同的子串
			if _, ok := seen[windowHash]; ok {
				// 当前窗口中的子串是重复出现的
				res[windowHash] = s[left:right]
			} else {
				// 当前窗口中的子串之前没有出现过，记下来
				seen[windowHash] = struct{}{}
			}
			// 缩小窗口，移出字符，并维护窗口哈希值（删除最高位数字）
			windowHash = windowHash - nums[left]*int(RL)
			left++
		}
	}

	var ret []string
	for _, v := range res {
		ret = append(ret, v)
	}

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
