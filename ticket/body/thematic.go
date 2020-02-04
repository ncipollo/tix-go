package body

type ThematicBreakSegment struct {}

func NewThematicBreakSegment() *ThematicBreakSegment {
	return &ThematicBreakSegment{}
}

func (t ThematicBreakSegment) Attributes() Attributes {
	return Attributes{}
}

func (t ThematicBreakSegment) Value() string {
	return ""
}
