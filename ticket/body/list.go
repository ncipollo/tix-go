package body

type ListStartSegment struct {
	isOrdered      bool
	level          int
	marker         string
	startingNumber int
}

func NewListStartSegment(isOrdered bool, level int, marker string, startingNumber int) *ListStartSegment {
	return &ListStartSegment{isOrdered: isOrdered, level: level, marker: marker, startingNumber: startingNumber}
}

func (l ListStartSegment) Attributes() Attributes {
	return Attributes{}
}

func (l ListStartSegment) Value() string {
	return ""
}
