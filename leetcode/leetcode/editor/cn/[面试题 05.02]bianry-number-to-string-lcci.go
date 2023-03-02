package cn

import (
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
// 十进制小数转二进制小数的方法是：小数部分乘以 2，
// 取整数部分作为二进制小数的下一位，小数部分作为下一次乘法的被乘数，
// 直到小数部分为 0 或者二进制小数的长度超过 32 位。
func printBin(num float64) string {
	s := strings.Builder{}
	s.WriteString("0.")
	for s.Len() < 32 && num != 0 {
		num *= 2
		x := int(num)
		s.WriteString(strconv.Itoa(x))
		num -= float64(x)
	}
	if num != 0 {
		return "ERROR"
	}
	return s.String()
}

//leetcode submit region end(Prohibit modification and deletion)
