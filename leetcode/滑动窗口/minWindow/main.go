package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	st := time.Now()
	sum := make(map[int64]int64)
	for i := 0; i < 100; i++ {
		a := rand2()
		sum[a]++
	}
	fmt.Println(sum)
	elapsed := time.Since(st)
	fmt.Println(float64(elapsed.Nanoseconds()) / 1e9)
}

func rand2() int64 {
	a, _ := rand.Int(rand.Reader, big.NewInt(10))
	return a.Int64()
}
