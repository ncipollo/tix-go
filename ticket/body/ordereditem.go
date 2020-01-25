package body

import "strconv"

type OrderedListItemSegment struct {
	level  uint
	number int
}

func NewOrderedListItemSegment(level uint, number int) *OrderedListItemSegment {
	return &OrderedListItemSegment{level: level, number: number}
}

func (o OrderedListItemSegment) Attributes() Attributes {
	return Attributes{Level: o.level, Number: o.number}
}

func (o OrderedListItemSegment) Value() string {
	return strconv.Itoa(o.number)
}
