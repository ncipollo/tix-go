package main

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

const md = `
# Epic Name

Epic descriptions

## Issue 1

Text 1

## Issue 2

Text 2
`

func main() {
	source := []byte(md)
	reader := text.NewReader(source)
	parser := goldmark.DefaultParser()
	node := parser.Parse(reader)
	ast.Walk(node, func(visited ast.Node, entering bool) (ast.WalkStatus, error) {
		status := ast.WalkStatus(ast.WalkContinue)
		kind := visited.Kind()
		println(kind.String())
		return status, nil
	})

	println(node)
}