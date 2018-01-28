package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

// 实现一个只支持int,只支持加法,函数只有println,只支持一个参数的Go语言子集....
func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", `
package main
func main(){
	// comments
	x:=1
	println(x+1)
}`, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	// 变量表
	var vars = map[string]int{}
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if ok && fn.Name.Name == "main" {
			// 进入main函数
			for _, stmt := range fn.Body.List {
				// 独立表达式的语句
				exprstmt, ok := stmt.(*ast.ExprStmt)
				if ok {
					call, ok := exprstmt.X.(*ast.CallExpr)
					if ok {
						ident, ok := call.Fun.(*ast.Ident)
						if ok && ident.Name == "println" {
							expr, ok := call.Args[0].(*ast.BinaryExpr)
							if ok && expr.Op == token.ADD {
								x := expr.X.(*ast.Ident)
								y := expr.Y.(*ast.BasicLit)
								yn, _ := strconv.ParseInt(y.Value, 10, 64)
								arg := vars[x.Name] + int(yn)
								println(arg)
							}
						}
					}
				}
				// 语句
				ass, ok := stmt.(*ast.AssignStmt)
				if ok {
					x, ok := ass.Lhs[0].(*ast.Ident)
					if ok {
						y, ok := ass.Rhs[0].(*ast.BasicLit)
						if ok {
							yn, err := strconv.ParseInt(y.Value, 10, 64)
							if err != nil {
								panic(err)
							}
							vars[x.Name] = int(yn)
						}

					}
				}
			}
		}
	}
}