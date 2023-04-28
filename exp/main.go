package main

import (
	"ast_graph/exp/multi"
	"fmt"
)

func main() {
	a := 1
	b := 2
	fmt.Println(a+b, add(a, b), multi.Multi(a, b), devide(a, b))

}

func devide(i, j int) int {
	return i / j
}
