package body

type CodeSpanSegment struct {
	linkedSegmentTraversal
	text string
}

func NewCodeSpanSegment(text string) *CodeSpanSegment {
	return &CodeSpanSegment{text: text}
}

func (c CodeSpanSegment) Attributes() Attributes {
	return Attributes{}
}

func (c CodeSpanSegment) Value() string {
	return c.text
}
