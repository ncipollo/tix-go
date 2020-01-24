package md

import "testing"

const md = `
# Epic Name

- List item 1
   - Nested Item
- List item 2

## Issue 1

The first issue

### Subissue 1

Sub issue!

### Subissue 2

Sub issue!

## Issue 2

The second issue

`

func TestParse(t *testing.T) {
	source := []byte(md)
	parser := NewParser()
	parser.Parse(source)
}

