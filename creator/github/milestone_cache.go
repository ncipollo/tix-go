package github

import (
	"errors"
	"fmt"
	"github.com/google/go-github/v29/github"
)

type MilestoneCache struct {
	api              Api
	milestonesById   map[int64]*github.Milestone
	milestonesByName map[string]*github.Milestone
}

func NewMilestoneCache(api Api) *MilestoneCache {
	return &MilestoneCache{
		api:              api,
		milestonesById:   make(map[int64]*github.Milestone),
		milestonesByName: make(map[string]*github.Milestone),
	}
}

func (c *MilestoneCache) AddMilestone(milestone *github.Milestone) {
	if milestone == nil || milestone.ID == nil || milestone.Title == nil {
		return
	}
	c.milestonesById[*milestone.ID] = milestone
	c.milestonesByName[*milestone.Title] = milestone
}

func (c *MilestoneCache) GetById(id int64) (*github.Milestone, error) {
	milestone := c.milestonesById[id]
	if milestone != nil {
		return milestone, nil
	}

	err := c.populateCache()
	if err != nil {
		return nil, err
	}

	milestone = c.milestonesById[id]
	if milestone != nil {
		return milestone, nil
	}

	return nil, errors.New(fmt.Sprintf(":scream: No milestone for id: %d", id))
}

func (c *MilestoneCache) GetByName(name string) (*github.Milestone, error) {
	milestone := c.milestonesByName[name]
	if milestone != nil {
		return milestone, nil
	}

	err := c.populateCache()
	if err != nil {
		return nil, err
	}

	milestone = c.milestonesByName[name]
	if milestone != nil {
		return milestone, nil
	}

	return c.creatMilestone(name)
}

func (c *MilestoneCache) populateCache() error {
	milestones, err := c.api.ListMilestones()
	if err != nil {
		return err
	}
	for _, milestone := range milestones {
		c.AddMilestone(milestone)
	}
	return nil
}

func (c *MilestoneCache) creatMilestone(name string) (*github.Milestone, error) {
	milestoneRequest := &github.Milestone{Title: &name}
	milestone, err := c.api.CreateMilestone(milestoneRequest)
	if err != nil {
		return nil, err
	}
	c.AddMilestone(milestone)
	return milestone, nil
}
