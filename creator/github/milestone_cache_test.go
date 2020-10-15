package github

import (
	"errors"
	"github.com/google/go-github/v29/github"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMilestoneCache_GetById_CacheHit(t *testing.T) {
	api := newMockApi()
	cache := NewMilestoneCache(api)

	milestone := createMilestone(1, "name")
	cache.AddMilestone(milestone)
	result, err := cache.GetById(1)

	assert.Equal(t, milestone, result)
	assert.NoError(t, err)
}

func TestMilestoneCache_GetById_CacheMiss(t *testing.T) {
	api := newMockApi()
	milestones := []*github.Milestone{createMilestone(1, "name")}
	api.On("ListMilestones").Return(milestones, nil)
	cache := NewMilestoneCache(api)

	result, err := cache.GetById(2)

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestMilestoneCache_GetById_PopulatesCache_Error(t *testing.T) {
	api := newMockApi()
	api.On("ListMilestones").Return(nil, errors.New("error"))
	cache := NewMilestoneCache(api)

	result, err := cache.GetById(1)

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestMilestoneCache_GetById_PopulatesCache_Success(t *testing.T) {
	api := newMockApi()
	milestones := []*github.Milestone{createMilestone(1, "name")}
	api.On("ListMilestones").Return(milestones, nil)
	cache := NewMilestoneCache(api)

	result, err := cache.GetById(1)

	assert.Equal(t, milestones[0], result)
	assert.NoError(t, err)
}

func TestMilestoneCache_GetByName_CacheHit(t *testing.T) {
	api := newMockApi()
	cache := NewMilestoneCache(api)
	milestone := createMilestone(1, "name")
	cache.AddMilestone(milestone)

	result, err := cache.GetByName("name")

	assert.Equal(t, milestone, result)
	assert.NoError(t, err)
}

func TestMilestoneCache_GetByName_PopulatesCache_Error(t *testing.T) {
	api := newMockApi()
	api.On("ListMilestones").Return(nil, errors.New("error"))
	cache := NewMilestoneCache(api)

	result, err := cache.GetByName("name")

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestMilestoneCache_GetByName_PopulatesCache_Success(t *testing.T) {
	api := newMockApi()
	milestones := []*github.Milestone{createMilestone(1, "name")}
	api.On("ListMilestones").Return(milestones, nil)
	cache := NewMilestoneCache(api)

	result, err := cache.GetByName("name")

	assert.Equal(t, milestones[0], result)
	assert.NoError(t, err)
}

func TestMilestoneCache_GetByName_CreatesMilestone_Error(t *testing.T) {
	api := newMockApi()
	newName := "name2"
	newMilestone := &github.Milestone{Title: &newName}
	milestones := []*github.Milestone{createMilestone(1, "name")}
	api.On("ListMilestones").Return(milestones, nil)
	api.On("CreateMilestone", newMilestone).Return(newMilestone, errors.New("err"))
	cache := NewMilestoneCache(api)

	result, err := cache.GetByName(newName)

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestMilestoneCache_GetByName_CreatesMilestone_Success(t *testing.T) {
	api := newMockApi()
	newName := "name2"
	newMilestone := &github.Milestone{Title: &newName}
	milestones := []*github.Milestone{createMilestone(1, "name")}
	api.On("ListMilestones").Return(milestones, nil)
	api.On("CreateMilestone", newMilestone).Return(newMilestone, nil)
	cache := NewMilestoneCache(api)

	result, err := cache.GetByName(newName)

	assert.Equal(t, newMilestone, result)
	assert.NoError(t, err)
}

func createMilestone(id int64, title string) *github.Milestone {
	return &github.Milestone{
		ID:    &id,
		Title: &title,
	}
}
