package md

import (
	"bytes"
	"github.com/yuin/goldmark/ast"
	"gopkg.in/yaml.v2"
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

	var err error
	switch strings.ToLower(language) {
	case "tix":
		err = c.addDefaultFields(state, code)
	case "jira":
		err = c.addJiraFields(state, code)
	case "github":
		err = c.addGithubFields(state, code)
	default:
		c.addCodeBlockSegment(state, code, language)
	}

	return err
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

func (c CodeBlockSegmentParser) addDefaultFields(state *State, code string) error {
	currentTicket := state.CurrentTicket()
	fields, err := c.fieldsFromMetadata(code)
	if err == nil {
		currentTicket.UpdateDefaultFields(fields)
	}
	return err
}

func (c CodeBlockSegmentParser) addJiraFields(state *State, code string) error {
	currentTicket := state.CurrentTicket()
	fields, err := c.fieldsFromMetadata(code)
	if err == nil {
		currentTicket.AddFieldsForTicketSystem(fields, "jira")
	}
	return err
}

func (c CodeBlockSegmentParser) addGithubFields(state *State, code string) error {
	currentTicket := state.CurrentTicket()
	fields, err := c.fieldsFromMetadata(code)
	if err == nil {
		currentTicket.AddFieldsForTicketSystem(fields, "github")
	}
	return err
}

func (c CodeBlockSegmentParser) fieldsFromMetadata(code string) (map[string]interface{}, error) {
	fields := make(map[string]interface{})
	err := yaml.Unmarshal([]byte(code), fields)
	return fields, err
}
