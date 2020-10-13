package body

type LinkSegment struct {
	linkedSegmentTraversal
	text string
	url  string
}

func NewLinkSegment(text string, url string) *LinkSegment {
	return &LinkSegment{text: text, url: url}
}

func (l LinkSegment) Attributes() Attributes {
	return Attributes{Url: l.url}
}

func (l LinkSegment) Value() string {
	return l.text
}
