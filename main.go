package main

import "astgraph/gen"

func main() {
	path := "./exp/main.go"
	dpath := "./"
	gen.GenSvg(path, dpath, "tree")
}
