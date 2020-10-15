package dryrun

import "strings"

type LevelLabel struct {
	singular string
	plural   string
}

func NewLevelLabel(singular string, plural string) *LevelLabel {
	return &LevelLabel{
		singular: singular,
		plural:   plural,
	}
}

func (l LevelLabel) Singular() string {
	return strings.Title(l.singular)
}

func (l LevelLabel) Plural() string {
	return strings.Title(l.plural)
}
