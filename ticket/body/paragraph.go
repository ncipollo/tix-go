package body

type ParagraphSegment struct {
	text string
}

func NewTParagraphSegment() *ParagraphSegment {
	return &ParagraphSegment{}
}

func (t ParagraphSegment) Attributes() Attributes {
	return Attributes{}
}

func (t ParagraphSegment) Value() string {
	return t.text
}
