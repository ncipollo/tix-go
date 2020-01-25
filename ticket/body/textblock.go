package body

type TextBlockSegment struct {
}

func NewTextBlockSegment() *TextBlockSegment {
	return &TextBlockSegment{}
}

func (t TextBlockSegment) Attributes() Attributes {
	return Attributes{}
}

func (t TextBlockSegment) Value() string {
	return ""
}
