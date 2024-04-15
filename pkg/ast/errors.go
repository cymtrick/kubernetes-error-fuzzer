package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

type Entry struct {
	FunctionName string `json:"functionName"`
	Position     string `json:"position"`
}

func main() {
	path := "../kubelet/kubelet.go"
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		panic(err)
	}

	var entries []Entry

	ast.Inspect(node, func(n ast.Node) bool {
		callExpr, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}

		if ident, ok := selExpr.X.(*ast.Ident); ok {
			fmt.Println("Found an identifier:", ident.Name)
			if ident.Name == "klog" {
				if strings.HasSuffix(selExpr.Sel.Name, "ErrorS") || strings.HasSuffix(selExpr.Sel.Name, "InfoS") {
					position := fset.Position(callExpr.Pos()).String()
					entry := Entry{
						FunctionName: selExpr.Sel.Name,
						Position:     position,
					}
					entries = append(entries, entry)
				}
			}
			if ident.Name == "fmt" {
				if strings.HasSuffix(selExpr.Sel.Name, "Errorf") {
					position := fset.Position(callExpr.Pos()).String()
					entry := Entry{
						FunctionName: selExpr.Sel.Name,
						Position:     position,
					}
					entries = append(entries, entry)
				}
			}
		}
		return true
	})

	jsonData, err := json.MarshalIndent(entries, "", "    ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("log_entries.json", jsonData, 0644)
	if err != nil {
		panic(err)
	}
}
