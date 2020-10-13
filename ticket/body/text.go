package body

type TextSegment struct {
	linkedSegmentTraversal
	text string
}

func NewTextSegment(text string) *TextSegment {
	return &TextSegment{text: text}
}

func (t TextSegment) Attributes() Attributes {
	return Attributes{}
}

func (t TextSegment) Value() string {
	return t.text
}
