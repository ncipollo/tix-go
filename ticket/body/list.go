package body

type ListStartSegment struct {
	isOrdered      bool
	level          int
	marker         string
	startingNumber int
}

func NewBulletListStartSegment(level int, marker string) *ListStartSegment {
	return &ListStartSegment{level: level, marker: marker}
}

func NewOrderedListStartSegment(level int, startingNumber int) *ListStartSegment {
	return &ListStartSegment{isOrdered: true, level: level, startingNumber: startingNumber}
}

func (l ListStartSegment) Attributes() Attributes {
	return Attributes{}
}

func (l ListStartSegment) Value() string {
	return ""
}
