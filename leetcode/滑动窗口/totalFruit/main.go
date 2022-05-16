package main

func main() {

}

// 水果成篮
// 地址：https://leetcode.cn/problems/fruit-into-baskets/
// 难度：中等
func totalFruit(fruits []int) int {
	// 建立一个hash表，记录已经在篮子里的水果的个数
	hashMap := make(map[int]int)
	j, res := 0, 0
	for i := 0; i < len(fruits); i++ {
		// 我们想得到以i为结尾的区间，所以i必须取到
		hashMap[fruits[i]]++
		// 当hash表长度为3时，表示篮子中右3种水果不符合题意，j指针需要往右移动
		for j <= i && len(hashMap) == 3 {
			// j指针向右移动，将j指向的水果数目-1
			hashMap[fruits[j]]--
			// 当hashMap[fruits[j]] == 0时表示篮子里没有这种水果了，直接delete
			if hashMap[fruits[j]] == 0 {
				delete(hashMap, fruits[j])
			}
			j++
		}
		if res < i-j+1 {
			res = i - j + 1
		}
	}
	return res
}
