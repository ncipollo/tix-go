package dryrun

import (
	"fmt"
	"strings"
)

type TicketStats struct {
	levelLabels  []*LevelLabel
	countByLevel map[int]int
}

func NewTicketStats(levelLabels []*LevelLabel) *TicketStats {
	return &TicketStats{
		levelLabels:  levelLabels,
		countByLevel: make(map[int]int),
	}
}

func (t *TicketStats) CountTicket(level int) {
	t.countByLevel[level]++
}

func (t TicketStats) Render(builder *strings.Builder) {
	builder.WriteString("Ticket Stats:\n")
	builder.WriteString(fmt.Sprintf("- Total Tickets: %d\n", t.total()))
	for level, label := range t.levelLabels {
		builder.WriteString(fmt.Sprintf("- %s: %d\n", label.Plural(), t.countByLevel[level]))
	}
}

func (t *TicketStats) total() int {
	sum := 0
	for _, value := range t.countByLevel {
		sum += value
	}
	return sum
}
