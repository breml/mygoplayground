package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

var (
	failFunc = `
package test

import "fmt"

type test struct {
	value string
}

func (t *test) fail1() {
	fmt.Printf("in *t: %s\n", t.value)
}
`

	okFunc = `
package test

import "fmt"

type test struct {
	value string
}

func (t *test) fail1() {
    if t != nil {
	    fmt.Printf("in *t: %s\n", t.value)
    }
}`
)

type visitor struct {
	receiver string
}

func (v visitor) Visit(node ast.Node) (w ast.Visitor) {
	if node != nil {
		fmt.Printf("in visitor, node: %#v\n", node)
	}
	switch t := node.(type) {
	case *ast.FuncDecl:
		if t.Recv != nil {
			fmt.Printf("we have a method, receiver: %s\n", t.Recv.List[0].Names[0].Name)
			if _, ok := t.Recv.List[0].Type.(*ast.StarExpr); ok {
				v.receiver = t.Recv.List[0].Names[0].Name
			} else {
				return nil
			}
		} else {
			fmt.Println("we have a function")
		}
	case *ast.Ident:
		if t.Name == v.receiver {
			fmt.Println("pointer receiver accessed")
		}
	case *ast.BinaryExpr:
		if t.Op == token.NEQ {
			fmt.Println("NEQ")
		}
	default:
	}
	return v
}

func main() {
	fset := token.NewFileSet()
	p, err := parser.ParseFile(fset, "failFunc.go", failFunc, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ast.Walk(visitor{}, p)
	// for _, d := range p.Decls {
	// 	var v visitor
	// 	v.Visit(d)
	// }

	fmt.Println("===")

	p, err = parser.ParseFile(fset, "okFunc.go", okFunc, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ast.Walk(visitor{}, p)
	// for _, d := range p.Decls {
	// 	var v visitor
	// 	v.Visit(d)
	// }
}
