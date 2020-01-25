package body

type LineBreakSegment struct {
}

func NewLineBreakSegment() *LineBreakSegment {
	return &LineBreakSegment{}
}

func (l LineBreakSegment) Attributes() Attributes {
	return Attributes{}
}

func (l LineBreakSegment) Value() string {
	return "\n"
}
