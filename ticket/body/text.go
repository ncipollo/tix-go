package body

type TextSegment struct {
	text string
}

func (t TextSegment) Attributes() Attributes {
	return Attributes{}
}

func (t TextSegment) Value() string {
	return t.text
}
