package main

import "fmt"

type Flags uint

const (
	FlagUp           Flags = 1 << iota //向上 1
	FlagBroadcast                      //支持广播访问 2
	FlagLoopback                       //是环回接口 4
	FlagPointToPoint                   //属于点对点链路 8
	FlagMulticast                      //支持多路广播访问 16
)

func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}
func TurnDown(v *Flags) {
	*v &^= FlagUp
}

func SetBroadcast(v *Flags) {
	*v |= FlagBroadcast
}

func IsCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMulticast) != 0
}

func main() {
	fmt.Println(FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) //10001 true
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) //10000 false
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   //10010 false
	fmt.Printf("%b %t\n", v, IsCast(v)) //10010 true
}