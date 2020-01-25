package body

type CodeBlockSegment struct {
	code     string
	language string
}

func NewCodeBlockSegment(code string, language string) *CodeBlockSegment {
	return &CodeBlockSegment{code: code, language: language}
}

func (c CodeBlockSegment) Attributes() Attributes {
	return Attributes{Language:c.language}
}

func (c CodeBlockSegment) Value() string {
	return c.code
}
