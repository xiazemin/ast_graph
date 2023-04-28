package main

import "ast_graph/gen"

func main() {
	gen.GenSvg("exp/main.go", "./", "tree")
}
