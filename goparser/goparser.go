package main

import (
	"fmt"
	"go/build"
	"go/parser"
	"go/token"
	"os"
	"runtime"
	"strings"

	_ "github.com/kisielk/gotool"
)

// GOROOT returns the $GOROOT environment variable
func GOROOT() string {
	return runtime.GOROOT()
}

// GOPATH returns the $GOPATH environment variable
func GOPATH() string {
	return os.Getenv("GOPATH")
}

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "goparser.go", nil, parser.ImportsOnly)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nImports")
	for _, s := range f.Imports {
		fmt.Println(s.Path.Value)
		pkg, err := build.Import(strings.Trim(s.Path.Value, "\""), "", 0)
		if err != nil {
			fmt.Println("Not found")
		} else {
			fmt.Println(pkg.Goroot, pkg.Name, pkg.Dir)
		}
	}

	/*
		fmt.Println("\nImport Paths")
		for _, s := range gotool.ImportPaths([]string{"all"}) {
			fmt.Println(s)
			pkg, err := build.Import(s, "", 0)
			if err != nil {
				fmt.Println("Not found")
			} else {
				fmt.Println(pkg.Goroot, pkg.Name, pkg.Dir)
			}
		}
	*/

	/*
		fmt.Println("\nGO Env")
		fmt.Println(GOROOT(), GOPATH())

		fmt.Println("\nBuild")
		for _, s := range build.Default.SrcDirs() {
			fmt.Println(s)
			pkg, err := build.Import("golang.org/x/tools/oracle/serial", "", 0)
			if err != nil {
				fmt.Println("Not found")
			} else {
				fmt.Println(pkg.Goroot, pkg.Name, pkg.Dir)
			}
		}
	*/
}
