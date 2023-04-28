1,安装graphviz

2，go get https://github.com/xiazemin/ast_graph

3，go run main.go

`````
package main

import "ast_graph/gen"

func main() {
	gen.GenSvg("exp/main.go", "./", "tree")
}

```