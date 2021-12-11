package gen

import (
	"astgraph/dot"
	"astgraph/tree"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os/exec"
	"path/filepath"
	"runtime"
)

type Visitor int

func (this Visitor) Visit(node ast.Node) (w ast.Visitor) {
	return this
}

var commands = map[string]string{
	"windows": "cmd /c start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func GenSvg(spath, dpath, name string) {
	fset := token.NewFileSet()
	path, _ := filepath.Abs(spath)
	f, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
	if err != nil {
		log.Println(err)
		return
	}
	var v Visitor
	tree.Walk(v, f)
	dot.GenTreeDot(dpath, name)

	run, ok := commands[runtime.GOOS]
	if !ok {
		fmt.Println(runtime.GOOS, run)
	}
	//Open calls the OS default program for uri
	cmd := exec.Command("/Applications/Google Chrome.app/Contents/MacOS/Google Chrome", dpath+name+".svg")
	cmd.Start()
}
