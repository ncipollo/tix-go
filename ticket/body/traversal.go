package body

type SegmentTraversal interface {
	Next() Segment
	SetNext(segment Segment)
	Previous() Segment
	SetPrevious(segment Segment)
}

type linkedSegmentTraversal struct {
	next Segment
	previous Segment
}

func (l *linkedSegmentTraversal) Next() Segment {
	return l.next
}

func (l *linkedSegmentTraversal) SetNext(segment Segment) {
	l.next = segment
}

func (l *linkedSegmentTraversal) Previous() Segment {
	return l.previous
}

func (l *linkedSegmentTraversal) SetPrevious(segment Segment) {
	l.previous = segment
}