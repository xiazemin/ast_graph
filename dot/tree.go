package dot

import (
	"os/exec"

	"github.com/xiazemin/ast_graph/graph"
)

// GenTreeDot  generate the dot of ast
func GenTreeDot(path, name string) string {
	WriteToFile(path, name+".dot", buildTree())
	cmd := exec.Command("/usr/local/bin/dot", "-Tsvg", path+name+".dot", "-o", path+name+".svg")
	cmd.Run()
	return name
}

func buildTree() string {
	graphs := ""
	for k := range graph.NodeList {
		graphs += k.Dot(graph.NodeList)
	}

	for k := range graph.EdageMap {
		graphs += k.Dot(graph.EdageMap)
	}
	style := "\tsubgraph clustera { \n\tstyle=invis;\n\trank=same;\n" +
		"\trankdir=LR;\n\tnode[shape=record, width=.1, height=.1,color=black,style=solid,fontcolor=black];\n"
	node := "\tdefault[label=\"{<default>default}\",color=black,style=solid,fontcolor=black ];\n" +
		"\tmain[label=\"{<main>mian}\", color=green,style=filled,fontcolor=red  ];\n" +
		"\tFuncDecl[label=\"{<FuncDecl>FuncDecl}\", color=green,style=filled ];\n " +
		"\tFile[label=\"{<File>File}\", color=blue,style=filled ];\n" +
		"\tImportSpec[label=\"{<ImportSpec>ImportSpec}\", color=red,style=filled,fontcolor=black ];\n " +
		"\tCallExpr[label=\"{<CallExpr>CallExpr}\", color=green,style=dashed ];\n" +
		"\tPackage[label=\"{<Package>Package}\", color=red,style=dashed ];\n " +
		"\tGenDecl[label=\"{<GenDecl>GenDecl}\", color=pink,style=dashed];\n" +
		"\tdefault -> main -> FuncDecl -> File -> GenDecl -> ImportSpec -> CallExpr -> Package;\n\t}\n"
	return getDotHead() + graphs + style + node + getDotTail()
}
