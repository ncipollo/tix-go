package md

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListState_CompleteAllLists(t *testing.T) {
	state := NewListState()
	state.StartBulletList("-")
	state.StartBulletList("-")
	state.StartBulletList("-")

	state.CompleteAllLists()

	assert.Empty(t, state.ListPath)
}

func TestListState_CompleteList_PopsAllLists(t *testing.T) {
	state := NewListState()
	for i := 0; i < 10; i++ {
		state.StartBulletList("-")
	}
	for i := 0; i < 10; i++ {
		state.CompleteList()
	}

	assert.Empty(t, state.ListPath)
}

func TestListState_CurrentList_ReturnsNilWhenEmpty(t *testing.T) {
	state := NewListState()

	currentList := state.CurrentList()

	assert.Nil(t, currentList)
}

func TestListState_CurrentList_ReturnsList(t *testing.T) {
	state := NewListState()
	state.StartBulletList("-")

	currentList := state.CurrentList()

	assert.Equal(t, state.ListPath[0], currentList)
}

func TestListState_CurrentList_ReturnsSublist(t *testing.T) {
	state := NewListState()
	state.StartBulletList("-")
	state.StartBulletList("-")
	state.ListPath[1].CurrentNumber = 5

	currentList := state.CurrentList()

	assert.Equal(t, state.ListPath[1], currentList)
}

func TestListState_ListLevel(t *testing.T) {
	state := NewListState()
	state.StartBulletList("-")
	state.StartBulletList("-")
	state.StartBulletList("-")

	assert.Equal(t, 3, state.ListLevel())
}

func TestListState_StartBulletList(t *testing.T) {
	state := NewListState()
	state.StartBulletList("-")
	state.StartBulletList("-")
	state.StartBulletList("-")

	assert.Equal(t, &ListInfo{Marker: "-"}, state.CurrentList())
	assert.Len(t, state.ListPath, 3)
}

func TestListState_StartOrderedList(t *testing.T) {
	state := NewListState()
	state.StartOrderedList(1)
	state.StartOrderedList(1)
	state.StartOrderedList(1)

	assert.Equal(t, &ListInfo{CurrentNumber: 1, IsOrdered: true}, state.CurrentList())
	assert.Len(t, state.ListPath, 3)
}
