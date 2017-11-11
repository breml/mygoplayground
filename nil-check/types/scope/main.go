package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"regexp"
)

func main() {
	// Parse the source files for a package.
	fset := token.NewFileSet()
	var files []*ast.File
	for _, file := range []struct{ name, input string }{
		{"main.go", `
package main
import "fmt"
func main() {
	freezing := FToC(-18)
	fmt.Println(freezing, Boiling) }
`},
		{"celsius.go", `
package main
import "fmt"
type Celsius float64
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func FToC(f float64) Celsius { return Celsius(f - 32 / 9 * 5) }
const Boiling Celsius = 100
`},
	} {
		f, err := parser.ParseFile(fset, file.name, file.input, 0)
		if err != nil {
			log.Fatal(err)
		}
		files = append(files, f)
	}

	// Type-check a package consisting of these files.
	// Type information for the imported "fmt" package
	// comes from $GOROOT/pkg/$GOOS_$GOOARCH/fmt.a.
	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("temperature", fset, files, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Print the tree of scopes.
	// For determinism, we redact addresses.
	var buf bytes.Buffer
	pkg.Scope().WriteTo(&buf, 0, true)
	rx := regexp.MustCompile(` 0x[a-fA-F0-9]*`)
	fmt.Println(rx.ReplaceAllString(buf.String(), ""))
}
