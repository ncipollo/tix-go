package body

type EmphasisSegment struct {
	text string
}

func NewEmphasisSegment(text string) *EmphasisSegment {
	return &EmphasisSegment{text: text}
}

func (e EmphasisSegment) Attributes() Attributes {
	return Attributes{}
}

func (e EmphasisSegment) Value() string {
	return e.text
}
