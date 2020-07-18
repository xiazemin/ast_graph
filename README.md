1,安装graphviz

2，go get https://github.com/xiazemin/ast_graph

3，run
`
package main

import "github.com/xiazemin/golang/ast/ast_graph/gen"

func main() {
	path := "/Users/didi/goLang/src/github.com/xiazemin/golang/ast/ast_graph/exp/main.go"
	dpath := "/Users/didi/goLang/src/github.com/xiazemin/golang/ast/ast_graph/"
	gen.GenSvg(path, dpath, "tree")
}

`