package main

func main() {
	for _, v := range letterCombinations("234") {
		println(v)
	}
}

var (
	phoneMap = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	combinations []string
)

// 电话号码的字母组合，中等
// 回溯算法
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations

}

func backtrack(digits string, index int, tempStr string) {
	if index == len(digits) { //终止条件, 字符串长度等于digits长度
		combinations = append(combinations, tempStr) //收集结果
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ { //单层搜索逻辑
			tempStr = tempStr + string(letters[i]) //拼接
			backtrack(digits, index+1, tempStr)
			tempStr = tempStr[:len(tempStr)-1] // 回溯
		}
	}
}
