package body

type Segment interface {
	Attributes() Attributes
	Value() string
}