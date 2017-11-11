package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	// src is the input for which we want to inspect the AST.
	src := `
package main

func main() {
	var m map[string]bool
	m["data"] = true
}`

	src = `
package main

func main() {
	var f func()
	f()
}
`

	//idents := make(map[string]string)

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	// Inspect the AST and print all identifiers and literals.
	for _, d := range f.Decls {
		ast.Print(fset, d)

		v := newVisitor(fset, src)
		ast.Walk(v, d)
		fmt.Printf("%s\n", v.idents)

		// ast.Inspect(d, func(n ast.Node) bool {
		// 	var s string
		// 	switch x := n.(type) {
		// 	case *ast.BasicLit:
		// 		s = x.Value
		// 	case *ast.Ident:
		// 		s = x.Name
		// 		// case *ast.DeclStmt:
		// 		// 	return false
		// 	}
		// 	if s != "" {
		// 		fmt.Printf("%s:\t%s\n", fset.Position(n.Pos()), s)
		// 	}
		// 	return true
		// })
	}

	//ast.Print(fset, f)
}
