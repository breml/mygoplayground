package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestWalk(t *testing.T) {
	for _, test := range []struct {
		input string
		// nilerrors []string
	}{
		{
			input: `
package main

func main() {
	var f func()
	f()
}
`,
		},
	} {
		fset := token.NewFileSet() // positions are relative to fset
		f, err := parser.ParseFile(fset, "", test.input, 0)
		if err != nil {
			t.Fatalf("err: %s, failed to parse: %s", err, test.input)
		}

		v := newVisitor(fset, test.input)
		ast.Walk(v, f)

		t.Logf("%s\n\n%s", test.input, v.idents)

		for _, errLine := range v.uncheckedNil {
			t.Logf("Pos: %s, ErrLine: %s\n", errLine.pos, errLine.line)
		}

		//ast.Print(fset, f)

		t.Fatalf("fatal")
	}

}
