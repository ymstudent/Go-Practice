package test2

const N = 1000

var a [N]int

//go:noinline
func g0(a *[N]int) {
	for i := range a {
		a[i] = i // line 12 // 371 ns/op
	}
}

//go:noinline
func g1(a *[N]int) {
	_ = *a // line 18 // 触发gc，防止出现抢占？ 349 ns/op
	for i := range a {
		a[i] = i // line 20
	}
}
