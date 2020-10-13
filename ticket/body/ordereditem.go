package body

import "strconv"

type OrderedListItemSegment struct {
	linkedSegmentTraversal
	level  int
	number int
}

func NewOrderedListItemSegment(level int, number int) *OrderedListItemSegment {
	return &OrderedListItemSegment{level: level, number: number}
}

func (o OrderedListItemSegment) Attributes() Attributes {
	return Attributes{Level: o.level, Number: o.number}
}

func (o OrderedListItemSegment) Value() string {
	return strconv.Itoa(o.number)
}
