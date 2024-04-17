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
	ErrorString  string `json:"errorString,omitempty"` // Optional field to include error strings
}

func main() {
	path := "../kubelet/kubelet.go"
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		return
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
			if (ident.Name == "fmt" || ident.Name == "klog") &&
				(strings.HasSuffix(selExpr.Sel.Name, "Errorf") ||
					strings.HasSuffix(selExpr.Sel.Name, "ErrorS") ||
					strings.HasSuffix(selExpr.Sel.Name, "InfoS")) {
				position := fset.Position(callExpr.Pos()).String()
				var errorString string
				if len(callExpr.Args) > 0 {
					formatArg, ok := callExpr.Args[0].(*ast.BasicLit)
					if ok && formatArg.Kind == token.STRING {
						formatStr := strings.Trim(formatArg.Value, "\"")
						args := make([]interface{}, len(callExpr.Args)-1)
						for i, arg := range callExpr.Args[1:] {
							args[i] = argToString(arg)
						}
						errorString = fmt.Sprintf(formatStr, args...)
					}
				}

				entry := Entry{
					FunctionName: selExpr.Sel.Name,
					Position:     position,
					ErrorString:  errorString,
				}
				entries = append(entries, entry)
			}
		}
		return true
	})

	jsonData, err := json.MarshalIndent(entries, "", "    ")
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return
	}

	err = os.WriteFile("log_entries.json", jsonData, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
}

func argToString(arg ast.Expr) string {
	switch x := arg.(type) {
	case *ast.BasicLit:
		return x.Value
	case *ast.Ident:
		return x.Name
	default:
		return fmt.Sprintf("%v", arg)
	}
}
