package initchecker

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var InitChecker = &analysis.Analyzer{
	Name: "InitChecker",
	Doc:  `This analyzer suggests "good" initilization behaviors.`,
	Run:  runAnalyzer,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func runAnalyzer(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.AssignStmt)(nil),
		(*ast.GenDecl)(nil),
		(*ast.CompositeLit)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch x := n.(type) {
		case *ast.AssignStmt:
			// check for slice init `var a []int` should be `a := []int{}`
			// check for map init with make
			for _, rhs := range x.Rhs {
				if cl, ok := rhs.(*ast.CompositeLit); ok && cl.Type != nil && len(cl.Elts) == 0 {
					switch cl.Type.(type) {
					case *ast.ArrayType:
						pass.Reportf(cl.Pos(), "consider using 'var' for empty slice initialization to avoid unnecessary memory allocation")
					case *ast.MapType:
						pass.Reportf(cl.Pos(), "consider using 'make' for map initialization to be explicit about intent")
					}
				}
			}
		case *ast.GenDecl:
			// Check for common variable initilazation
			if x.Tok == token.VAR {
				for _, spec := range x.Specs {
					if vs, ok := spec.(*ast.ValueSpec); ok && len(vs.Values) > 0 {
						pass.Reportf(vs.Pos(), "consider using ':=' for variable initialization for conciseness")
					}
				}
			}
		case *ast.CompositeLit:
			// Check for empty struct initilazation
			if _, ok := x.Type.(*ast.Ident); ok && len(x.Elts) == 0 {
				pass.Reportf(x.Pos(), "consider using 'var' for zero value struct initialization to improve clarity")
			}
		}
	})

	return nil, nil
}
