package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s := Constructor([]int{3, 1, 7, 9})
	for i := 0; i < 10; i++ {
		fmt.Println(s.PickIndex())
	}
}

type Solution struct {
	w   []int
	sum int
}

func Constructor(w []int) Solution {
	sum := 0
	for _, v := range w {
		sum += v
	}
	return Solution{w: w, sum: sum}
}

func (this *Solution) PickIndex() int {
	//rand.Seed(time.Now().Unix())
	r := rand.Intn(this.sum) + 1
	fmt.Println("sum:=", this.sum)
	fmt.Println("r:=", r)
	ans := 0
	for i, v := range this.w {
		ans += v
		if r <= ans {
			return i
		}
	}
	return 0
}
