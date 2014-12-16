package main

import (
	"fmt"
	"random"
)

func main() {
	fmt.Println(random.RandIntN(10, 100))
	fmt.Println(random.RandIntN(100, 1000))
}
