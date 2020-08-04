package main

import "github.com/xiazemin/ast_graph/gen"

func main() {
	path := "./main.go"
	dpath := "./tree"
	gen.GenSvg(path, dpath, "tree")
}
