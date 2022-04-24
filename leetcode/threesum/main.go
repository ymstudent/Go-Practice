package main

import (
	"fmt"
	"sort"
)

func main() {
	res := threeSum([]int{3, 0, -2, -1, 1, 2})
	for _, v := range res {
		fmt.Println(v)
	}
}

func threeSum(nums []int) [][]int {
	l := len(nums)
	sort.Ints(nums)
	var ans [][]int

	for first := 0; first < l; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		third := l - 1
		target := -1 * nums[first]

		for second := first + 1; second < l; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			for second < third && nums[second]+nums[third] > target {
				third--
			}

			if second == third {
				break
			}

			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}

	return ans
}
