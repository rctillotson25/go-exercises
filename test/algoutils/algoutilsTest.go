package main

import (
	"fmt"
	"algorithms/algoutils"
	"random"
)



func main() {
	a := random.RandIntN(11,100)
	l,r := algoutils.Split(a)
	fmt.Println(a)
	fmt.Println(l)
	fmt.Println(r)
}
