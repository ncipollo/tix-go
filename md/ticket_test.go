package md

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTicketParser_Parse_ErrorWhenHeadingLevelSkipped(t *testing.T) {
	text := `
# Ticket 1
### Error Ticket
`
	parser := NewTicketParser()
	state, node := setupTextParser(text)

	err := parser.Parse(state, node)

	assert.Error(t, err)
}

func TestTicketParser_Parse_ErrorWhenTitleMissing(t *testing.T) {
	text := `
#
`
	parser := NewTicketParser()
	state, node := setupTextParser(text)

	err := parser.Parse(state, node)

	assert.Error(t, err)
}

func TestTicketParser_Parse_ErrorWhenTitleNonText(t *testing.T) {
	text := `
#  **foo**
`
	parser := NewTicketParser()
	state, node := setupTextParser(text)

	err := parser.Parse(state, node)

	assert.Error(t, err)
}

func TestTicketParser_Parse_ErrorWhenTextOutsideOfHeading(t *testing.T) {
	text := `
Hello, I'm an issue!

# Ticket 1
`
	parser := NewTicketParser()
	state, node := setupTextParser(text)

	err := parser.Parse(state, node)

	assert.Error(t, err)
}

func TestTicketParser_Parse_CreatesMultipleRootTickets(t *testing.T) {
	text := `
# Ticket 1
# Ticket 2
# Ticket 3
`
	parser := NewTicketParser()
	state, node := setupTextParser(text)

	err := parser.Parse(state, node)

	assert.NoError(t, err)
	assert.Len(t, state.RootTickets, 3)
	assert.Equal(t, "Ticket 1", state.RootTickets[0].Title)
	assert.Equal(t, "Ticket 2", state.RootTickets[1].Title)
	assert.Equal(t, "Ticket 3", state.RootTickets[2].Title)
}

func TestTicketParser_Parse_CreatesNestedTickets(t *testing.T) {
	text := `
# Ticket 1
## Sub 1
### Sub 2
`
	parser := NewTicketParser()
	state, node := setupTextParser(text)

	err := parser.Parse(state, node)

	assert.NoError(t, err)
	assert.Len(t, state.RootTickets, 1)
	assert.Equal(t, "Ticket 1", state.RootTickets[0].Title)
	assert.Equal(t, "Sub 1", state.RootTickets[0].Subtickets[0].Title)
	assert.Equal(t, "Sub 2", state.RootTickets[0].Subtickets[0].Subtickets[0].Title)
}

func TestTicketParser_Parse_CompletesNestedTicketsOnNewRootTicket(t *testing.T) {
	text := `
# Ticket 1
## Sub 1
### Sub 2
# Ticket 2
`
	parser := NewTicketParser()
	state, node := setupTextParser(text)

	err := parser.Parse(state, node)

	assert.NoError(t, err)
	assert.Len(t, state.RootTickets, 2)
	assert.Equal(t, "Ticket 1", state.RootTickets[0].Title)
	assert.Equal(t, "Sub 1", state.RootTickets[0].Subtickets[0].Title)
	assert.Equal(t, "Sub 2", state.RootTickets[0].Subtickets[0].Subtickets[0].Title)
	assert.Equal(t, "Ticket 2", state.RootTickets[1].Title)
}