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
	return Attributes{Level: l.level, Number: l.startingNumber}
}

func (l ListStartSegment) Value() string {
	return ""
}

type ListEndSegment struct {
	isOrdered bool
	level     int
}

func NewBulletListEndSegment(level int) *ListEndSegment {
	return &ListEndSegment{level: level}
}

func NewOrderedListEndSegment(level int) *ListEndSegment {
	return &ListEndSegment{isOrdered: true, level: level}
}

func (l ListEndSegment) Attributes() Attributes {
	return Attributes{Level: l.level}
}

func (l ListEndSegment) Value() string {
	return ""
}
