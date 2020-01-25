package body

type ListStartSegment struct {
}

func NewListStartSegment() *ListStartSegment {
	return &ListStartSegment{}
}

func (l ListStartSegment) Attributes() Attributes {
	return Attributes{}
}

func (l ListStartSegment) Value() string {
	return ""
}
