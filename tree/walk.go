package tree

import (
	"fmt"
	"github.com/xiazemin/golang/ast/ast_graph/graph"
	"go/ast"
)

func Walk(v ast.Visitor, node ast.Node) {
	if v = v.Visit(node); v == nil {
		return
	}

	// walk children
	// (the order of the cases matches the order
	// of the corresponding node types in ast.go)
	switch n := node.(type) {
	// Comments and fields
	case *ast.Comment:
		no := graph.NewNode(n, "Comment", n.Text, fmt.Sprintf("%T", n), "Comment")
		graph.AddNode(no)
		// nothing to do

	case *ast.CommentGroup:
		nos := graph.NewNode(n, "CommentGroup", fmt.Sprint(len(n.List)), fmt.Sprintf("%T", n), "CommentGroup")
		graph.AddNode(nos)
		for _, c := range n.List {
			Walk(v, c)
			nod := graph.NewNode(c, "Comment", c.Text, fmt.Sprintf("%T", c), "Comment")
			graph.AddEdage(nos, nod)
		}

	case *ast.Field:
		nos := graph.NewNode(n, "Field", "Doc_Names_Type_Tag_Comment", fmt.Sprintf("%T", n), "Field")
		graph.AddNode(nos)
		if n.Doc != nil {
			Walk(v, n.Doc)
			nod := graph.NewNode(n.Doc, "Field.Doc", fmt.Sprint(len(n.Doc.List)), fmt.Sprintf("%T", n.Doc), "CommentGroup")
			graph.AddEdage(nos, nod)
		}
		for _, x := range n.Names {
			Walk(v, x)
			nod := graph.NewNode(x, x.Name, x.Obj.Name, fmt.Sprintf("%T", x), "Ident")
			graph.AddEdage(nos, nod)
		}
		Walk(v, n.Type)
		nod := graph.NewNode(n.Type, "Expr", "Field.Type", fmt.Sprintf("%T", n.Type), "Expr")
		graph.AddEdage(nos, nod)
		if n.Tag != nil {
			Walk(v, n.Tag)
			nod := graph.NewNode(n.Tag, n.Tag.Value, "Field.Tag", fmt.Sprintf("%T", n.Tag), "BasicLit")
			graph.AddEdage(nos, nod)
		}
		if n.Comment != nil {
			Walk(v, n.Comment)
			nod := graph.NewNode(n, "Field.Comment", fmt.Sprint(len(n.Comment.List)), fmt.Sprintf("%T", n.Comment), "CommentGroup")
			graph.AddEdage(nos, nod)
		}

	case *ast.FieldList:
		nos := graph.NewNode(n, "FieldList", fmt.Sprint(len(n.List)), fmt.Sprintf("%T", n), "FieldList")
		graph.AddNode(nos)
		for _, f := range n.List {
			Walk(v, f)
			nod := graph.NewNode(f, "FieldList.Field", fmt.Sprint(len(f.Names)), fmt.Sprintf("%T", f), "Field")
			graph.AddEdage(nos, nod)
		}

	// Expressions
	case *ast.BadExpr:
		nos := graph.NewNode(n, "BadExpr", fmt.Sprint(n.From.IsValid()), fmt.Sprintf("%T", n), "BadExpr")
		graph.AddNode(nos)
	case *ast.Ident:
		nos := graph.NewNode(n, n.Name, n.Name, fmt.Sprintf("%T", n), "Ident")
		graph.AddNode(nos)
	case *ast.BasicLit:
		//nos:=graph.NewNode(n,n.Kind.String(),n.Value,"BasicLit")
		nos := graph.NewNode(n, n.Kind.String(), "BasicLit", fmt.Sprintf("%T", n), "BasicLit")
		graph.AddNode(nos)
		// nothing to do

	case *ast.Ellipsis:
		nos := graph.NewNode(n, "Ellipsis", "...", fmt.Sprintf("%T", n), "Ellipsis")
		graph.AddNode(nos)
		if n.Elt != nil {
			Walk(v, n.Elt)
			nod := graph.NewNode(n.Elt, "Ellipsis.Expr", "Expr", fmt.Sprintf("%T", n.Elt), "Expr")
			graph.AddEdage(nos, nod)
		}

	case *ast.FuncLit:
		nos := graph.NewNode(n, "FuncLit", "FuncLit", fmt.Sprintf("%T", n), "FuncLit")
		graph.AddNode(nos)
		Walk(v, n.Type)
		nod := graph.NewNode(n.Type, "FuncLit.Type", "FuncType", fmt.Sprintf("%T", n.Type), "FuncType")
		graph.AddEdage(nos, nod)
		Walk(v, n.Body)
		nodb := graph.NewNode(n.Body, "FuncLit.Body", "BlockStmt", fmt.Sprintf("%T", n.Body), "BlockStmt")
		graph.AddEdage(nos, nodb)

	case *ast.CompositeLit:
		nos := graph.NewNode(n, "CompositeLit", "CompositeLit", fmt.Sprintf("%T", n), "CompositeLit")
		graph.AddNode(nos)
		if n.Type != nil {
			Walk(v, n.Type)
			nod := graph.NewNode(n.Type, "CompositeLit.Type", "Expr", fmt.Sprintf("%T", n.Type), "Expr")
			graph.AddEdage(nos, nod)
		}
		for _, x := range n.Elts {
			Walk(v, x)
			nod := graph.NewNode(x, "CompositeLit.Elts", "Expr", fmt.Sprintf("%T", x), "Expr")
			graph.AddEdage(nos, nod)
		}

	case *ast.ParenExpr:
		nos := graph.NewNode(n, "ParenExpr", "ParenExpr", fmt.Sprintf("%T", n), "ParenExpr")
		graph.AddNode(nos)
		Walk(v, n.X)

	case *ast.SelectorExpr:
		nos := graph.NewNode(n, "SelectorExpr", "SelectorExpr", fmt.Sprintf("%T", n), "SelectorExpr")
		graph.AddNode(nos)
		Walk(v, n.X)
		nod := graph.NewNode(n.X, "SelectorExpr.x", "Expr", fmt.Sprintf("%T", n.X), "Expr")
		graph.AddEdage(nos, nod)
		Walk(v, n.Sel)
		nod2 := graph.NewNode(n.Sel, "Ellipsis.Sel", n.Sel.Name, fmt.Sprintf("%T", n.Sel), "Ident")
		graph.AddEdage(nos, nod2)

	case *ast.IndexExpr:
		nos := graph.NewNode(n, "IndexExpr", "IndexExpr", fmt.Sprintf("%T", n), "IndexExpr")
		graph.AddNode(nos)
		Walk(v, n.X)
		nod := graph.NewNode(n.X, "IndexExpr.X", "Expr", fmt.Sprintf("%T", n.X), "Expr")
		graph.AddEdage(nos, nod)
		Walk(v, n.Index)
		nod2 := graph.NewNode(n.Index, "IndexExpr.Index", "Expr", fmt.Sprintf("%T", n.Index), "Expr")
		graph.AddEdage(nos, nod2)

	case *ast.SliceExpr:
		nos := graph.NewNode(n, "SliceExpr", "SliceExpr", fmt.Sprintf("%T", n), "SliceExpr")
		graph.AddNode(nos)
		Walk(v, n.X)
		nod := graph.NewNode(n.X, "SliceExpr.X", "Expr", fmt.Sprintf("%T", n.X), "Expr")
		graph.AddEdage(nos, nod)
		if n.Low != nil {
			Walk(v, n.Low)
			nod := graph.NewNode(n.Low, "IndexExpr.Low", "Expr", fmt.Sprintf("%T", n.Low), "Expr")
			graph.AddEdage(nos, nod)
		}
		if n.High != nil {
			Walk(v, n.High)
			nod := graph.NewNode(n.High, "IndexExpr.High", "Expr", fmt.Sprintf("%T", n.High), "Expr")
			graph.AddEdage(nos, nod)
		}
		if n.Max != nil {
			Walk(v, n.Max)
			nod := graph.NewNode(n.Max, "IndexExpr.High", "Expr", fmt.Sprintf("%T", n.Max), "Expr")
			graph.AddEdage(nos, nod)
		}

	case *ast.TypeAssertExpr:
		nos := graph.NewNode(n, "TypeAssertExpr", "TypeAssertExpr", fmt.Sprintf("%T", n), "TypeAssertExpr")
		graph.AddNode(nos)
		Walk(v, n.X)
		nod := graph.NewNode(n.X, "TypeAssertExpr.X", "Expr", fmt.Sprintf("%T", n.X), "Expr")
		graph.AddEdage(nos, nod)
		if n.Type != nil {
			Walk(v, n.Type)
			nod := graph.NewNode(n.Type, "TypeAssertExpr.Type", "Expr", fmt.Sprintf("%T", n.Type), "Expr")
			graph.AddEdage(nos, nod)
		}

	case *ast.CallExpr:
		nos := graph.NewNode(n, "CallExpr", "CallExpr", fmt.Sprintf("%T", n), "CallExpr")
		graph.AddNode(nos)
		Walk(v, n.Fun)
		nod := graph.NewNode(n.Fun, "CallExpr.Fun", "Expr", fmt.Sprintf("%T", n.Fun), "Expr")
		graph.AddEdage(nos, nod)
		for _, x := range n.Args {
			Walk(v, x)
			nod := graph.NewNode(x, "CallExpr.Args", "Expr", fmt.Sprintf("%T", x), "Expr")
			graph.AddEdage(nos, nod)
		}

	case *ast.StarExpr:
		nos := graph.NewNode(n, "StarExpr", "StarExpr", fmt.Sprintf("%T", n), "StarExpr")
		graph.AddNode(nos)
		Walk(v, n.X)
		nod := graph.NewNode(n.X, "StarExpr.X", "Expr", fmt.Sprintf("%T", n.X), "Expr")
		graph.AddEdage(nos, nod)

	case *ast.UnaryExpr:
		nos := graph.NewNode(n, "UnaryExpr", n.Op.String(), fmt.Sprintf("%T", n), "UnaryExpr")
		graph.AddNode(nos)
		Walk(v, n.X)
		nod := graph.NewNode(n.X, "UnaryExpr.X", "Expr", fmt.Sprintf("%T", n.X), "Expr")
		graph.AddEdage(nos, nod)

	case *ast.BinaryExpr:
		//nos:=graph.NewNode(n,"BinaryExpr",n.Op.String(),"BinaryExpr")
		nos := graph.NewNode(n, "BinaryExpr", "Expr", fmt.Sprintf("%T", n), "BinaryExpr")
		graph.AddNode(nos)
		Walk(v, n.X)
		nod := graph.NewNode(n.X, "BinaryExpr.X", "Expr", fmt.Sprintf("%T", n.X), "Expr")
		graph.AddEdage(nos, nod)
		Walk(v, n.Y)
		nod2 := graph.NewNode(n.Y, "BinaryExpr.Y", "Expr", fmt.Sprintf("%T", n.Y), "Expr")
		graph.AddEdage(nos, nod2)

	case *ast.KeyValueExpr:
		nos := graph.NewNode(n, "KeyValueExpr", "KeyValueExpr", fmt.Sprintf("%T", n), "KeyValueExpr")
		graph.AddNode(nos)
		Walk(v, n.Key)
		nod := graph.NewNode(n.Key, "KeyValueExpr.Key", "Expr", fmt.Sprintf("%T", n.Key), "Expr")
		graph.AddEdage(nos, nod)
		Walk(v, n.Value)
		nod2 := graph.NewNode(n.Value, "KeyValueExpr.Value", "Expr", fmt.Sprintf("%T", n.Value), "Expr")
		graph.AddEdage(nos, nod2)

	// Types
	case *ast.ArrayType:
		nos := graph.NewNode(n, "ArrayType", "ArrayType", fmt.Sprintf("%T", n), "ArrayType")
		graph.AddNode(nos)
		if n.Len != nil {
			Walk(v, n.Len)
			nod := graph.NewNode(n.Len, "ArrayType.Len", "Expr", fmt.Sprintf("%T", n.Len), "Expr")
			graph.AddEdage(nos, nod)
		}
		Walk(v, n.Elt)
		nod := graph.NewNode(n.Elt, "ArrayType.Elt", "Expr", fmt.Sprintf("%T", n.Elt), "Expr")
		graph.AddEdage(nos, nod)

	case *ast.StructType:
		nos := graph.NewNode(n, "StructType", "StructType", fmt.Sprintf("%T", n), "StructType")
		graph.AddNode(nos)
		Walk(v, n.Fields)
		nod := graph.NewNode(n.Fields, "StructType.Fields", fmt.Sprint(len(n.Fields.List)), fmt.Sprintf("%T", n.Fields), "FieldList")
		graph.AddEdage(nos, nod)

	case *ast.FuncType:
		nos := graph.NewNode(n, "FuncType", "FuncType", fmt.Sprintf("%T", n), "FuncType")
		graph.AddNode(nos)
		if n.Params != nil {
			Walk(v, n.Params)
			nod := graph.NewNode(n.Params, "FuncType.Params", fmt.Sprint(len(n.Params.List)), fmt.Sprintf("%T", n.Params), "FieldList")
			graph.AddEdage(nos, nod)
		}
		if n.Results != nil {
			Walk(v, n.Results)
			nod := graph.NewNode(n.Results, "FuncType.Results", fmt.Sprint(len(n.Results.List)), fmt.Sprintf("%T", n.Results), "FieldList")
			graph.AddEdage(nos, nod)
		}

	case *ast.InterfaceType:
		nos := graph.NewNode(n, "InterfaceType", "InterfaceType", fmt.Sprintf("%T", n), "InterfaceType")
		graph.AddNode(nos)
		Walk(v, n.Methods)
		nod := graph.NewNode(n.Methods, "InterfaceType.Methods", fmt.Sprint(len(n.Methods.List)), fmt.Sprintf("%T", n.Methods), "FieldList")
		graph.AddEdage(nos, nod)

	case *ast.MapType:
		nos := graph.NewNode(n, "MapType", "MapType", fmt.Sprintf("%T", n), "MapType")
		graph.AddNode(nos)
		Walk(v, n.Key)
		nod := graph.NewNode(n.Key, "MapType.Key", "Expr", fmt.Sprintf("%T", n.Key), "Expr")
		graph.AddEdage(nos, nod)
		Walk(v, n.Value)
		nod2 := graph.NewNode(n.Value, "MapType.Value", "Expr", fmt.Sprintf("%T", n.Value), "Expr")
		graph.AddEdage(nos, nod2)

	case *ast.ChanType:
		nos := graph.NewNode(n, "ChanType", "ChanType", fmt.Sprintf("%T", n), "ChanType")
		graph.AddNode(nos)
		Walk(v, n.Value)
		nod := graph.NewNode(n.Value, "ChanType.Value", "Expr", fmt.Sprintf("%T", n.Value), "Expr")
		graph.AddEdage(nos, nod)

	// Statements
	case *ast.BadStmt:
		nos := graph.NewNode(n, "BadStmt", "BadStmt", fmt.Sprintf("%T", n), "BadStmt")
		graph.AddNode(nos)
		// nothing to do

	case *ast.DeclStmt:
		nos := graph.NewNode(n, "DeclStmt", "DeclStmt", fmt.Sprintf("%T", n), "DeclStmt")
		graph.AddNode(nos)
		Walk(v, n.Decl)
		nod := graph.NewNode(n.Decl, "DeclStmt.Decl", "Decl", fmt.Sprintf("%T", n.Decl), "Decl")
		graph.AddEdage(nos, nod)

	case *ast.EmptyStmt:
		nos := graph.NewNode(n, "CompositeLit", "CompositeLit", fmt.Sprintf("%T", n), "CompositeLit")
		graph.AddNode(nos)
		// nothing to do

	case *ast.LabeledStmt:
		nos := graph.NewNode(n, "LabeledStmt", "LabeledStmt", fmt.Sprintf("%T", n), "LabeledStmt")
		graph.AddNode(nos)
		Walk(v, n.Label)
		nod := graph.NewNode(n.Label, "LabeledStmt.Label", "Ident", fmt.Sprintf("%T", n.Label), "Ident")
		graph.AddEdage(nos, nod)
		Walk(v, n.Stmt)
		nod2 := graph.NewNode(n.Stmt, "LabeledStmt.Stmt", "Stmt", fmt.Sprintf("%T", n.Stmt), "Stmt")
		graph.AddEdage(nos, nod2)

	case *ast.ExprStmt:
		nos := graph.NewNode(n, "ExprStmt", "ExprStmt", fmt.Sprintf("%T", n), "ExprStmt")
		graph.AddNode(nos)
		Walk(v, n.X)
		nod := graph.NewNode(n.X, "ExprStmt.X", "Expr", fmt.Sprintf("%T", n.X), "Expr")
		graph.AddEdage(nos, nod)

	case *ast.SendStmt:
		nos := graph.NewNode(n, "SendStmt", "SendStmt", fmt.Sprintf("%T", n), "SendStmt")
		graph.AddNode(nos)
		Walk(v, n.Chan)
		nod := graph.NewNode(n.Chan, "SendStmt.Chan", "Expr", fmt.Sprintf("%T", n.Chan), "Expr")
		graph.AddEdage(nos, nod)
		Walk(v, n.Value)
		nod2 := graph.NewNode(n.Value, "SendStmt.Value", "Expr", fmt.Sprintf("%T", n.Value), "Expr")
		graph.AddEdage(nos, nod2)

	case *ast.IncDecStmt:
		nos := graph.NewNode(n, "IncDecStmt", "IncDecStmt", fmt.Sprintf("%T", n), "IncDecStmt")
		graph.AddNode(nos)
		Walk(v, n.X)
		nod := graph.NewNode(n.X, "SendStmt.X", "Expr", fmt.Sprintf("%T", n.X), "Expr")
		graph.AddEdage(nos, nod)

	case *ast.AssignStmt:
		nos := graph.NewNode(n, "AssignStmt", "AssignStmt", fmt.Sprintf("%T", n), "AssignStmt")
		graph.AddNode(nos)
		for _, x := range n.Lhs {
			Walk(v, x)
			nod := graph.NewNode(x, "AssignStmt.Lhs", "Expr", fmt.Sprintf("%T", x), "Expr")
			graph.AddEdage(nos, nod)
		}
		for _, x := range n.Rhs {
			Walk(v, x)
			nod := graph.NewNode(x, "AssignStmt.Rhs", "Expr", fmt.Sprintf("%T", x), "Expr")
			graph.AddEdage(nos, nod)
		}

	case *ast.GoStmt:
		nos := graph.NewNode(n, "GoStmt", "GoStmt", fmt.Sprintf("%T", n), "GoStmt")
		graph.AddNode(nos)
		Walk(v, n.Call)
		nod := graph.NewNode(n.Call, "GoStmt.Call", "CallExpr", fmt.Sprintf("%T", n.Call), "CallExpr")
		graph.AddEdage(nos, nod)

	case *ast.DeferStmt:
		nos := graph.NewNode(n, "DeferStmt", "DeferStmt", fmt.Sprintf("%T", n), "DeferStmt")
		graph.AddNode(nos)
		Walk(v, n.Call)
		nod := graph.NewNode(n.Call, "DeferStmt.Call", "CallExpr", fmt.Sprintf("%T", n.Call), "CallExpr")
		graph.AddEdage(nos, nod)

	case *ast.ReturnStmt:
		nos := graph.NewNode(n, "ReturnStmt", "ReturnStmt", fmt.Sprintf("%T", n), "ReturnStmt")
		graph.AddNode(nos)
		for _, x := range n.Results {
			Walk(v, x)
			nod := graph.NewNode(x, "ReturnStmt.Results", fmt.Sprintf("%T", x), "Expr", "Expr")
			graph.AddEdage(nos, nod)
		}

	case *ast.BranchStmt:
		nos := graph.NewNode(n, "BranchStmt", "BranchStmt", fmt.Sprintf("%T", n), "BranchStmt")
		graph.AddNode(nos)
		if n.Label != nil {
			Walk(v, n.Label)
			nod := graph.NewNode(n.Label, "BranchStmt.Label", n.Label.Name, fmt.Sprintf("%T", n.Label), "Ident")
			graph.AddEdage(nos, nod)
		}

	case *ast.BlockStmt:
		nos := graph.NewNode(n, "BlockStmt", "BlockStmt", fmt.Sprintf("%T", n), "BlockStmt")
		graph.AddNode(nos)
		for _, x := range n.List {
			Walk(v, x)
			nod := graph.NewNode(x, "BlockStmt.List", "Stmt", fmt.Sprintf("%T", x), "Stmt")
			graph.AddEdage(nos, nod)
		}

	case *ast.IfStmt:
		nos := graph.NewNode(n, "IfStmt", "IfStmt", fmt.Sprintf("%T", n), "IfStmt")
		graph.AddNode(nos)
		if n.Init != nil {
			Walk(v, n.Init)
			nod := graph.NewNode(n.Init, "IfStmt.Init", "Stmt", fmt.Sprintf("%T", n.Init), "Stmt")
			graph.AddEdage(nos, nod)
		}
		Walk(v, n.Cond)
		nod := graph.NewNode(n.Cond, "IfStmt.Cond", "Expr", fmt.Sprintf("%T", n.Cond), "Expr")
		graph.AddEdage(nos, nod)
		Walk(v, n.Body)
		nod2 := graph.NewNode(n.Body, "IfStmt.Body", "BlockStmt", fmt.Sprintf("%T", n.Body), "BlockStmt")
		graph.AddEdage(nos, nod2)
		if n.Else != nil {
			Walk(v, n.Else)
			nod := graph.NewNode(n.Else, "IfStmt.Else", "Stmt", fmt.Sprintf("%T", n.Else), "Stmt")
			graph.AddEdage(nos, nod)
		}

	case *ast.CaseClause:
		nos := graph.NewNode(n, "CaseClause", "CaseClause", fmt.Sprintf("%T", n), "CaseClause")
		graph.AddNode(nos)
		for _, x := range n.List {
			Walk(v, x)
			nod := graph.NewNode(x, "CaseClause.List", "Expr", fmt.Sprintf("%T", x), "Expr")
			graph.AddEdage(nos, nod)
		}
		for _, x := range n.Body {
			Walk(v, x)
			nod := graph.NewNode(x, "CaseClause.Body", "Stmt", fmt.Sprintf("%T", x), "Stmt")
			graph.AddEdage(nos, nod)
		}
	case *ast.SwitchStmt:
		nos := graph.NewNode(n, "SwitchStmt", "SwitchStmt", fmt.Sprintf("%T", n), "SwitchStmt")
		graph.AddNode(nos)
		if n.Init != nil {
			Walk(v, n.Init)
			nod := graph.NewNode(n.Init, "SwitchStmt.Init", "Stmt", fmt.Sprintf("%T", n.Init), "Stmt")
			graph.AddEdage(nos, nod)
		}
		if n.Tag != nil {
			Walk(v, n.Tag)
			nod := graph.NewNode(n.Tag, "SwitchStmt.Tag", "Expr", fmt.Sprintf("%T", n.Tag), "Expr")
			graph.AddEdage(nos, nod)
		}
		Walk(v, n.Body)
		nod := graph.NewNode(n.Body, "SwitchStmt.Body", "BlockStmt", fmt.Sprintf("%T", n.Body), "BlockStmt")
		graph.AddEdage(nos, nod)

	case *ast.TypeSwitchStmt:
		nos := graph.NewNode(n, "TypeSwitchStmt", "TypeSwitchStmt", fmt.Sprintf("%T", n), "TypeSwitchStmt")
		graph.AddNode(nos)
		if n.Init != nil {
			Walk(v, n.Init)
			nod := graph.NewNode(n.Init, "TypeSwitchStmt.Init", "Stmt", fmt.Sprintf("%T", n.Init), "Stmt")
			graph.AddEdage(nos, nod)
		}
		Walk(v, n.Assign)
		nod := graph.NewNode(n.Assign, "TypeSwitchStmt.Assign", "Stmt", fmt.Sprintf("%T", n.Assign), "Stmt")
		graph.AddEdage(nos, nod)
		Walk(v, n.Body)
		nod2 := graph.NewNode(n.Body, "TypeSwitchStmt.Body", "BlockStmt", fmt.Sprintf("%T", n.Body), "BlockStmt")
		graph.AddEdage(nos, nod2)

	case *ast.CommClause:
		nos := graph.NewNode(n, "CommClause", "CommClause", fmt.Sprintf("%T", n), "CommClause")
		graph.AddNode(nos)
		if n.Comm != nil {
			Walk(v, n.Comm)
			nod := graph.NewNode(n.Comm, "CommClause.Comm", "Stmt", fmt.Sprintf("%T", n.Comm), "Stmt")
			graph.AddEdage(nos, nod)
		}
		for _, x := range n.Body {
			Walk(v, x)
			nod := graph.NewNode(x, "SwitchStmt.Body", "Stmt", fmt.Sprintf("%T", x), "Stmt")
			graph.AddEdage(nos, nod)
		}
	case *ast.SelectStmt:
		nos := graph.NewNode(n, "SelectStmt", "SelectStmt", fmt.Sprintf("%T", n), "SelectStmt")
		graph.AddNode(nos)
		Walk(v, n.Body)
		nod := graph.NewNode(n.Body, "SelectStmt.Body", "BlockStmt", fmt.Sprintf("%T", n.Body), "BlockStmt")
		graph.AddEdage(nos, nod)

	case *ast.ForStmt:
		nos := graph.NewNode(n, "ForStmt", "ForStmt", fmt.Sprintf("%T", n), "ForStmt")
		graph.AddNode(nos)
		if n.Init != nil {
			Walk(v, n.Init)
			nod := graph.NewNode(n.Init, "ForStmt.Init", "Stmt", fmt.Sprintf("%T", n.Init), "Stmt")
			graph.AddEdage(nos, nod)
		}
		if n.Cond != nil {
			Walk(v, n.Cond)
			nod := graph.NewNode(n.Cond, "ForStmt.Cond", "Expr", fmt.Sprintf("%T", n.Cond), "Expr")
			graph.AddEdage(nos, nod)
		}
		if n.Post != nil {
			Walk(v, n.Post)
			nod := graph.NewNode(n.Post, "ForStmt.Post", "Stmt", fmt.Sprintf("%T", n.Post), "Stmt")
			graph.AddEdage(nos, nod)
		}
		Walk(v, n.Body)
		nod := graph.NewNode(n.Body, "ForStmt.Body", "BlockStmt", fmt.Sprintf("%T", n.Body), "BlockStmt")
		graph.AddEdage(nos, nod)
	case *ast.RangeStmt:
		nos := graph.NewNode(n, "RangeStmt", "RangeStmt", fmt.Sprintf("%T", n), "RangeStmt")
		graph.AddNode(nos)
		if n.Key != nil {
			Walk(v, n.Key)
			nod := graph.NewNode(n.Key, "RangeStmt.Key", "Expr", fmt.Sprintf("%T", n.Key), "Expr")
			graph.AddEdage(nos, nod)
		}
		if n.Value != nil {
			Walk(v, n.Value)
			nod := graph.NewNode(n.Value, "RangeStmt.Body", "Expr", fmt.Sprintf("%T", n.Value), "Expr")
			graph.AddEdage(nos, nod)
		}
		Walk(v, n.X)
		nod := graph.NewNode(n.X, "RangeStmt.Body", "Expr", fmt.Sprintf("%T", n.X), "Expr")
		graph.AddEdage(nos, nod)
		Walk(v, n.Body)
		nod2 := graph.NewNode(n.Body, "RangeStmt.Body", "BlockStmt", fmt.Sprintf("%T", n.Body), "BlockStmt")
		graph.AddEdage(nos, nod2)

	// Declarations
	case *ast.ImportSpec:
		nos := graph.NewNode(n, "ImportSpec", "ImportSpec", fmt.Sprintf("%T", n), "ImportSpec")
		graph.AddNode(nos)
		if n.Doc != nil {
			Walk(v, n.Doc)
			nod := graph.NewNode(n.Doc, "ImportSpec.Doc", "CommentGroup", fmt.Sprintf("%T", n.Doc), "CommentGroup")
			graph.AddEdage(nos, nod)
		}
		if n.Name != nil {
			Walk(v, n.Name)
			nod := graph.NewNode(n.Name, "ImportSpec.Name", n.Name.Name, fmt.Sprintf("%T", n.Name), "Ident")
			graph.AddEdage(nos, nod)
		}
		Walk(v, n.Path)
		nod := graph.NewNode(n.Path, "ImportSpec.Path", n.Path.Value, fmt.Sprintf("%T", n.Path), "BasicLit")
		graph.AddEdage(nos, nod)
		if n.Comment != nil {
			Walk(v, n.Comment)
			nod := graph.NewNode(n.Comment, "ImportSpec.Comment", "CommentGroup", fmt.Sprintf("%T", n.Comment), "CommentGroup")
			graph.AddEdage(nos, nod)
		}

	case *ast.ValueSpec:
		nos := graph.NewNode(n, "ValueSpec", "ValueSpec", fmt.Sprintf("%T", n), "ValueSpec")
		graph.AddNode(nos)
		if n.Doc != nil {
			Walk(v, n.Doc)
			nod := graph.NewNode(n.Doc, "ValueSpec.Doc", "CommentGroup", fmt.Sprintf("%T", n.Doc), "CommentGroup")
			graph.AddEdage(nos, nod)
		}
		for _, x := range n.Names {
			Walk(v, x)
			nod := graph.NewNode(x, "ValueSpec.Names", x.Name, fmt.Sprintf("%T", x), "Ident")
			graph.AddEdage(nos, nod)
		}

		if n.Type != nil {
			Walk(v, n.Type)
			nod := graph.NewNode(n.Type, "ValueSpec.Body", "Expr", fmt.Sprintf("%T", n.Type), "Expr")
			graph.AddEdage(nos, nod)
		}
		for _, x := range n.Values {
			Walk(v, x)
			nod := graph.NewNode(x, "ValueSpec.Values", "Expr", fmt.Sprintf("%T", x), "Expr")
			graph.AddEdage(nos, nod)
		}

		if n.Comment != nil {
			Walk(v, n.Comment)
			nod := graph.NewNode(n.Comment, "ValueSpec.Comment", "CommentGroup", fmt.Sprintf("%T", n.Comment), "CommentGroup")
			graph.AddEdage(nos, nod)
		}

	case *ast.TypeSpec:
		nos := graph.NewNode(n, "TypeSpec", "TypeSpec", fmt.Sprintf("%T", n), "TypeSpec")
		graph.AddNode(nos)
		if n.Doc != nil {
			Walk(v, n.Doc)
			nod := graph.NewNode(n.Doc, "TypeSpec.Doc", "CommentGroup", fmt.Sprintf("%T", n.Doc), "CommentGroup")
			graph.AddEdage(nos, nod)
		}
		Walk(v, n.Name)
		nod := graph.NewNode(n.Name, "TypeSpec.Name", n.Name.Name, fmt.Sprintf("%T", n.Name), "Ident")
		graph.AddEdage(nos, nod)
		Walk(v, n.Type)
		nod2 := graph.NewNode(n.Type, "TypeSpec.Type", "Expr", fmt.Sprintf("%T", n.Type), "Expr")
		graph.AddEdage(nos, nod2)
		if n.Comment != nil {
			Walk(v, n.Comment)
			nod := graph.NewNode(n.Comment, "TypeSpec.Comment", "CommentGroup", fmt.Sprintf("%T", n.Comment), "CommentGroup")
			graph.AddEdage(nos, nod)
		}

	case *ast.BadDecl:
		nos := graph.NewNode(n, "BadDecl", "BadDecl", fmt.Sprintf("%T", n), "BadDecl")
		graph.AddNode(nos)
		// nothing to do

	case *ast.GenDecl:
		nos := graph.NewNode(n, "GenDecl", "GenDecl", fmt.Sprintf("%T", n), "GenDecl")
		graph.AddNode(nos)
		if n.Doc != nil {
			Walk(v, n.Doc)
			nod := graph.NewNode(n.Doc, "GenDecl.Doc", "CommentGroup", fmt.Sprintf("%T", n.Doc), "CommentGroup")
			graph.AddEdage(nos, nod)
		}
		for _, s := range n.Specs {
			Walk(v, s)
			nod := graph.NewNode(s, "GenDecl.Specs", "Spec", fmt.Sprintf("%T", s), "Spec")
			graph.AddEdage(nos, nod)
		}

	case *ast.FuncDecl:
		nos := graph.NewNode(n, "FuncDecl", "FuncDecl", fmt.Sprintf("%T", n), "FuncDecl")
		graph.AddNode(nos)
		if n.Doc != nil {
			Walk(v, n.Doc)
			nod := graph.NewNode(n.Doc, "FuncDecl.Doc", "CommentGroup", fmt.Sprintf("%T", n.Doc), "CommentGroup")
			graph.AddEdage(nos, nod)
		}
		if n.Recv != nil {
			Walk(v, n.Recv)
			nod := graph.NewNode(n.Recv, "FuncDecl.Recv", "FieldList", fmt.Sprintf("%T", n.Recv), "FieldList")
			graph.AddEdage(nos, nod)
		}
		Walk(v, n.Name)
		nod := graph.NewNode(n.Name, "FuncDecl.Name", n.Name.Name, fmt.Sprintf("%T", n.Name), "Ident")
		graph.AddEdage(nos, nod)
		Walk(v, n.Type)
		nod2 := graph.NewNode(n.Type, "FuncDecl.Type", "FuncType", fmt.Sprintf("%T", n.Type), "FuncType")
		graph.AddEdage(nos, nod2)
		if n.Body != nil {
			Walk(v, n.Body)
			nod := graph.NewNode(n.Body, "FuncDecl.Body", "BlockStmt", fmt.Sprintf("%T", n.Body), "BlockStmt")
			graph.AddEdage(nos, nod)
		}

	// Files and packages
	case *ast.File:
		nos := graph.NewNode(n, "File", "File", fmt.Sprintf("%T", n), "File")
		graph.AddNode(nos)
		if n.Doc != nil {
			Walk(v, n.Doc)
			nod := graph.NewNode(n.Doc, "File.Doc", "CommentGroup", fmt.Sprintf("%T", n.Doc), "CommentGroup")
			graph.AddEdage(nos, nod)
		}
		Walk(v, n.Name)
		nod := graph.NewNode(n.Name, "File.Name", n.Name.Name, fmt.Sprintf("%T", n.Name), "Ident")
		graph.AddEdage(nos, nod)
		for _, x := range n.Decls {
			Walk(v, x)
			nod := graph.NewNode(x, "File.Decls", "Decl", fmt.Sprintf("%T", x), "Decl")
			graph.AddEdage(nos, nod)
		}

		// don't walk n.Comments - they have been
		// visited already through the individual
		// nodes

	case *ast.Package:
		nos := graph.NewNode(n, "Package", "Package", fmt.Sprintf("%T", n), "Package")
		graph.AddNode(nos)
		for _, f := range n.Files {
			Walk(v, f)
			nod := graph.NewNode(f, "Package.Files", f.Name.Name, fmt.Sprintf("%T", f), "File")
			graph.AddEdage(nos, nod)
		}

	default:
		nos := graph.NewNode(n, "CompositeLit", "CompositeLit", fmt.Sprintf("%T", n), "CompositeLit")
		graph.AddNode(nos)
		panic(fmt.Sprintf("ast.Walk: unexpected node type %T", n))
	}

	v.Visit(nil)
}
