package md

import "testing"

const md = `
# Epic Name

- List item 1
   - Nested Item
- List item 2
`

func TestParse(t *testing.T) {
	source := []byte(md)
	parser := NewParser()
	parser.Parse(source)
}

