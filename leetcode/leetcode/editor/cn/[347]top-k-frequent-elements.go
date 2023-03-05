package cn

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	values := make([][]int, 0, len(m))
	for k, v := range m {
		values = append(values, []int{k, v})
	}
	//values = quickSort(values, 0, len(values)-1)
	sort.Slice(values, func(i, j int) bool {
		return values[i][1] > values[j][1]
	})
	res := make([]int, k)
	for i, v := range values[:k] {
		res[i] = v[0]
	}
	return res
}

func quickSort(values [][]int, low, high int) [][]int {
	if low >= high {
		return values
	}
	piovt := values[high][1]
	i := low
	for j := low; j < high; j++ {
		if values[j][1] > piovt {
			values[i], values[j] = values[j], values[i]
			i++
		}
	}

	values[i], values[high] = values[high], values[i]

	quickSort(values, low, i-1)
	quickSort(values,i+1,high)
	return values
}
//leetcode submit region end(Prohibit modification and deletion)
