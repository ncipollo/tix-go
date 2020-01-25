package md

import (
	"bytes"
	"github.com/yuin/goldmark/ast"
	"tix/ticket/body"
)

type CodeBlockSegmentParser struct {
	fenced bool
}

func NewCodeBlockSegmentParser(fenced bool) *CodeBlockSegmentParser {
	return &CodeBlockSegmentParser{fenced: fenced}
}

func (c CodeBlockSegmentParser) Parse(state *State, node ast.Node) error {
	if c.fenced {
		return c.parseFencedBlock(state, node.(*ast.FencedCodeBlock))
	} else {
		return c.parseNormalBlock(state, node.(*ast.CodeBlock))
	}
	return nil
}

func (c CodeBlockSegmentParser) parseNormalBlock(state *State, node *ast.CodeBlock) error {
	currentTicket := state.CurrentTicket()

	code  := c.textFromBlock(state, node.BaseBlock)
	codeSpan := body.NewCodeBlockSegment(code, "")
	currentTicket.AddBodySegment(codeSpan)

	return nil
}

func (c CodeBlockSegmentParser) parseFencedBlock(state *State, node *ast.FencedCodeBlock) error {
	currentTicket := state.CurrentTicket()

	code  := c.textFromBlock(state, node.BaseBlock)
	languageData := node.Language(state.SourceData)
	codeSpan := body.NewCodeBlockSegment(code, string(languageData))
	currentTicket.AddBodySegment(codeSpan)

	return nil
}

func (c CodeBlockSegmentParser) textFromBlock(state *State, node ast.BaseBlock) string {

	buffer := bytes.NewBuffer(nil)
	l := node.Lines().Len()
	for ii := 0; ii < l; ii++ {
		line := node.Lines().At(ii)
		buffer.Write(line.Value(state.SourceData))
	}
	data := buffer.Bytes()

	return string(data)
}
