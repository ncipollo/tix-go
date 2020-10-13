package body

type Segment interface {
	SegmentTraversal
	Attributes() Attributes
	Value() string
}