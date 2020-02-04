package body

type BlockQuoteSegment struct {}

func NewBlockQuoteSegment() *BlockQuoteSegment {
	return &BlockQuoteSegment{}
}

func (b BlockQuoteSegment) Attributes() Attributes {
	return Attributes{}
}

func (b BlockQuoteSegment) Value() string {
	return ""
}
