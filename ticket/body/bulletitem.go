package body

type BulletListItemSegment struct {
	level  uint
	marker string
}

func NewBulletListItemSegment(level uint, marker string) *BulletListItemSegment {
	return &BulletListItemSegment{level: level, marker: marker}
}

func (b BulletListItemSegment) Attributes() Attributes {
	return Attributes{Level: b.level}
}

func (b BulletListItemSegment) Value() string {
	return b.marker
}
