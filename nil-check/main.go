package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

var ()

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "_testdata/nil-panic2.go", nil, 0)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println("Decls")
	for _, d := range f.Decls {
		printNode(d)
	}
}

func printNode(n ast.Node) {
	//fmt.Printf("%#v\n", n)
	switch t := n.(type) {
	default:
		// fmt.Printf("Default: %#v\n", t)
	// case *ast.GenDecl:
	// 	fmt.Printf("%s, GenDecl: %v\n", t.Tok.String(), t)
	// 	for _, s := range t.Specs {
	// 		printNode(s)
	// 	}
	case *ast.FuncDecl:
		fmt.Printf("FuncDecl: %s\n", t.Name.String())
		for _, s := range t.Body.List {
			walkFunc(s)
		}

		// case *ast.TypeSpec:
		// 	fmt.Printf("%s, TypeSpec: %v\n", t.Name.Name, t)
		// 	printNode(t.Type)
		// case *ast.InterfaceType:
		// 	fmt.Printf("InterfaceType: %v", t)
	}
}

func walkFunc(stmt ast.Node) {
	switch s := stmt.(type) {
	default:
		fmt.Println("Default")
		fmt.Printf("%#v\n", s)
	case *ast.DeclStmt:
		fmt.Println("DeclStmt")
		walkFunc(s.Decl)
	case *ast.GenDecl:
		fmt.Println("GenDecl: ", s.Tok.String())
		for _, sp := range s.Specs {
			walkFunc(sp)
		}
	case *ast.ValueSpec:
		fmt.Println("ValueSpec")
		walkFunc(s.Type)
		for _, n := range s.Names {
			fmt.Println(n.Name)
		}
		for _, v := range s.Values {
			walkFunc(v)
		}
	case *ast.MapType:
		fmt.Println("MapType")
		walkFunc(s.Key)
		walkFunc(s.Value)
	case *ast.ExprStmt:
		fmt.Println("ExprStmt")
		walkFunc(s.X)
	case *ast.CallExpr:
		fmt.Println("CallExpr")
		walkFunc(s.Fun)
	case *ast.SelectorExpr:
		fmt.Println("SelectorExpr")
		walkFunc(s.X)
		fmt.Printf("Selector: %s\n", s.Sel.Name)
	case *ast.Ident:
		fmt.Printf("Ident: %s\n", s.String())
	case *ast.AssignStmt:
		fmt.Println("AssignStmt")
		for _, l := range s.Lhs {
			walkFunc(l)
		}
		for _, r := range s.Rhs {
			walkFunc(r)
		}
	case *ast.IndexExpr:
		fmt.Println("IndexExpr")
		walkFunc(s.X)
	}
}
