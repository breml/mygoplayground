package main

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"strings"
)

type uncheckedNil struct {
	pos      token.Position
	nilstate nilstate
	line     string
}

type varstruct struct {
	vartype  int
	nilstate nilstate
}

type vartype int

const (
	NA = iota
	Func
)

func (vt vartype) String() string {
	switch vt {
	case Func:
		return fmt.Sprint("function")
	default:
		return fmt.Sprint("NA")
	}
}

type nilstate int

const (
	Unknown = iota
	SureNil
	SureNotNil
	MaybeNilExported
	MaybeNilPackageGlobal
)

func (ns nilstate) String() string {
	switch ns {
	case SureNil:
		return fmt.Sprint("nil")
	case SureNotNil:
		return fmt.Sprint("notnil")
	case MaybeNilExported:
		return fmt.Sprint("maybe nil (exported)")
	case MaybeNilPackageGlobal:
		return fmt.Sprint("maybe nil (package global)")
	default:
		return fmt.Sprint("unknown")
	}
}

type visitor struct {
	fset         *token.FileSet
	idents       map[string]varstruct
	uncheckedNil []uncheckedNil
	lines        map[string][]string
}

func newVisitor(fset *token.FileSet, source string) *visitor {
	v := &visitor{}
	v.fset = fset
	v.idents = make(map[string]varstruct)
	v.uncheckedNil = []uncheckedNil{}
	v.lines = make(map[string][]string)
	v.lines[""] = strings.Split(source, "\n")
	return v
}

func (v *visitor) Visit(n ast.Node) ast.Visitor {
	fmt.Printf("%#v\n", n)

	switch x := n.(type) {
	case *ast.ValueSpec:
		//fmt.Printf("+++%#v\n", x)
		relevant := false
		switch x.Type.(type) {
		case *ast.FuncType:
			relevant = true
		}
		if relevant {
			for i, n := range x.Names {
				if x.Values == nil || x.Values[i] == nil {
					v.idents[n.Name] = varstruct{nilstate: SureNil}
				} else {
					v.idents[n.Name] = varstruct{nilstate: SureNotNil}
				}
			}
		}
	case *ast.CallExpr:
		if i, ok := x.Fun.(*ast.Ident); ok {
			if v.idents[i.Name].nilstate != SureNotNil {
				v.addNilAtPosition(x.Pos(), v.idents[i.Name].nilstate) // = append(v.uncheckedNil, uncheckedNil{pos: x.Lparen, nilstate: v.idents[i.Name].nilstate, line: fmt.Sprintf("%s", i.String())})
			}
		}
	}
	return v
}

func (v *visitor) addNilAtPosition(pos token.Pos, nilstate nilstate) {
	position := v.fset.Position(pos)
	fmt.Println("+++", position)
	lines, ok := v.lines[position.Filename]
	fmt.Println("+++", position.Filename, lines, ok)
	if !ok {
		lines = readfile(position.Filename)
		v.lines[position.Filename] = lines
	}

	line := "??"
	if position.Line-1 < len(lines) {
		line = strings.TrimSpace(lines[position.Line-1])
	}
	v.uncheckedNil = append(v.uncheckedNil, uncheckedNil{pos: position, nilstate: nilstate, line: line})
}

func readfile(filename string) []string {
	var f, err = os.Open(filename)
	if err != nil {
		return nil
	}

	var lines []string
	var scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
