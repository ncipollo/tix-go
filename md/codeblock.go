package md

import (
	"bytes"
	"github.com/yuin/goldmark/ast"
	"strings"
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
}

func (c CodeBlockSegmentParser) parseFencedBlock(state *State, node *ast.FencedCodeBlock) error {
	code := c.textFromBlock(state, node.BaseBlock)
	languageData := node.Language(state.SourceData)
	language := string(languageData)
	if strings.ToLower(language) == "tix" {
		c.addTicketMetadata(state, code)
	} else {
		c.addCodeBlockSegment(state, code, language)
	}
	return nil
}

func (c CodeBlockSegmentParser) parseNormalBlock(state *State, node *ast.CodeBlock) error {
	code := c.textFromBlock(state, node.BaseBlock)
	c.addCodeBlockSegment(state, code, "")
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

func (c CodeBlockSegmentParser) addCodeBlockSegment(state *State, code string, language string) {
	currentTicket := state.CurrentTicket()
	codeBlock := body.NewCodeBlockSegment(code, language)
	currentTicket.AddBodySegment(codeBlock)
	currentTicket.AddBodyLineBreak()
}

func (c CodeBlockSegmentParser) addTicketMetadata(state *State, code string) {
	currentTicket := state.CurrentTicket()
	currentTicket.Metadata = code
}