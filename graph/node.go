package graph

import (
	"fmt"
	"go/ast"
	"strings"
)

type Node struct {
	FileName string
	Name     string
	Val      string
	Type     string
	Kind     string
	Id       string
	Raw      ast.Node
}

var NodeList = map[*Node]int{}

func (n *Node) getType() string {
	return n.replace(n.Type)
}

func (n *Node) getName() string {
	return n.replace(n.Name)
}

func (n *Node) getValue() string {
	return n.replace(n.Val)
}
func (n *Node) Key() string {
	return n.Id
}

func (n *Node) replace(s string) string {
	s = strings.Replace(s, ".", "_", -1)
	s = strings.Replace(s, "\"", "_", -1)
	s = strings.Replace(s, "/", "_", -1)
	s = strings.Replace(s, "<", "_", -1)
	s = strings.Replace(s, ">", "_", -1)
	return strings.Replace(s, "*", "_", -1)
}
func (n *Node) Dot(m map[*Node]int) string {
	statics := make(map[string]int)
	for k, v := range m {
		statics[k.Key()] += v
	}

	style := "\tnode[shape=record, width=.1, height=.1,color=black,style=solid,fontcolor=black];\n"
	switch raw := n.Raw.(type) {
	case *ast.FuncDecl:
		if raw.Name.Name == "main" {
			style = "\tnode[shape=record, width=.1, height=.1,color=green,style=filled,fontcolor=red];\n"
			return style + fmt.Sprintf("\t%s[label=\"{<%s> %s  |<%s> %d  |<%s> %s |<%s> %s |<%s> %s  }\"];\n",
				n.Key(), n.Kind, n.Kind, "count", statics[n.Key()], n.Key(), n.Key(), n.getType(), n.getType(), raw.Name.Name, raw.Name.Name)
		} else {
			style = "\tnode[shape=record, width=.1, height=.1,color=green,style=filled];\n"
			return style + fmt.Sprintf("\t%s[label=\"{<%s> %s  |<%s> %d  |<%s> %s |<%s> %s |<%s> %s }\"];\n",
				n.Key(), n.Kind, n.Kind, "count", statics[n.Key()], n.Key(), n.Key(), n.getType(), n.getType(), raw.Name.Name, raw.Name.Name)
		}
	case *ast.File:
		style = "\tnode[shape=record, width=.1, height=.1,color=blue,style=filled];\n"
	case *ast.ImportSpec:
		style = "\tnode[shape=record, width=.1, height=.1,color=red,style=filled,fontcolor=black];\n"
		name := "nil"
		if raw.Name != nil {
			name = raw.Name.Name
		}
		return style + fmt.Sprintf("\t%s[label=\"{<%s> %s  |<%s> %d  |<%s> %s |<%s> %s |<%s> %s |<%s> %s }\"];\n",
			n.Key(), n.Kind, n.Kind, "count", statics[n.Key()], n.Key(), n.Key(), n.getType(), n.getType(), name, name, n.replace(raw.Path.Value), n.replace(raw.Path.Value))
	case *ast.CallExpr:
		style = "\tnode[shape=record, width=.1, height=.1,color=green,style=dashed];\n"
	case *ast.Package:
		style = "\tnode[shape=record, width=.1, height=.1,color=red,style=dashed];\n"
	case *ast.GenDecl:
		style = "\tnode[shape=record, width=.1, height=.1,color=pink,style=dashed];\n"
		return style + fmt.Sprintf("\t%s[label=\"{<%s> %s  |<%s> %d  |<%s> %s |<%s> %s |<%s> %s }\"];\n",
			n.Key(), n.Kind, n.Kind, "count", statics[n.Key()], n.Key(), n.Key(), n.getType(), n.getType(), raw.Tok.String(), raw.Tok.String())
	case *ast.BasicLit:
		style = "\tnode[shape=record, width=.1, height=.1,color=yellow,style=dashed];\n"
		return style + fmt.Sprintf("\t%s[label=\"{<%s> %s  |<%s> %d  |<%s> %s |<%s> %s |<%s> %s }\"];\n",
			n.Key(), n.Kind, n.Kind, "count", statics[n.Key()], n.Key(), n.Key(), n.getType(), n.getType(), n.replace(raw.Value), n.replace(raw.Value))
	case *ast.Ident:
		if raw.Name == "main" {
			style = "\tnode[shape=record, width=.1, height=.1,color=black,style=solid,fontcolor=red];\n"
		}
		return style + fmt.Sprintf("\t%s[label=\"{<%s> %s  |<%s> %d  |<%s> %s |<%s> %s |<%s> %s }\"];\n",
			n.Key(), n.Kind, n.Kind, "count", statics[n.Key()], n.Key(), n.Key(), n.getType(), n.getType(), raw.Name, raw.Name)
	}
	return style + fmt.Sprintf("\t%s[label=\"{<%s> %s  |<%s> %d  |<%s> %s |<%s> %s }\"];\n",
		n.Key(), n.Kind, n.Kind, "count", statics[n.Key()], n.Key(), n.Key(), n.getType(), n.getType())
}

func NewNode(n ast.Node, name, value, deatilType, kind string) *Node {
	return &Node{
		Name: name,
		Val:  value,
		Type: deatilType,
		Kind: kind,
		Id:   fmt.Sprintf("POS_%d_%d", n.Pos(), n.End()),
		Raw:  n,
	}
}

func AddNode(nos *Node) {
	NodeList[nos]++
}
