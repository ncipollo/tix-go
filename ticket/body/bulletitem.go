package body

type BulletListItemSegment struct {
	linkedSegmentTraversal
	level  int
	marker string
}

func NewBulletListItemSegment(level int, marker string) *BulletListItemSegment {
	return &BulletListItemSegment{level: level, marker: marker}
}

func (b BulletListItemSegment) Attributes() Attributes {
	return Attributes{Level: b.level}
}

func (b BulletListItemSegment) Value() string {
	return b.marker
}
