package main

import (
	"fmt"

	"github.com/xiazemin/ast_graph/exp/multi"
)

func main() {
	a := 1
	b := 2
	fmt.Println(a+b, add(a, b), multi.Multi(a, b), devide(a, b))

}

func devide(i, j int) int {
	return i / j
}
