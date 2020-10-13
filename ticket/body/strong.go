package body

type StrongEmphasisSegment struct {
	linkedSegmentTraversal
	text string
}

func NewStrongEmphasisSegment(text string) *StrongEmphasisSegment {
	return &StrongEmphasisSegment{text: text}
}

func (e StrongEmphasisSegment) Attributes() Attributes {
	return Attributes{}
}

func (e StrongEmphasisSegment) Value() string {
	return e.text
}
