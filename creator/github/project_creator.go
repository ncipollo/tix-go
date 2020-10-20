package github

import (
	"errors"
	"fmt"
	"github.com/google/go-github/v29/github"
	"strconv"
	"tix/ticket"
)

type ProjectCreator interface {
	CreateProject(ticket *ticket.Ticket) (*github.Project, error)
	UpdateProject(ticket *ticket.Ticket, updateKey string) (*github.Project, error)
}

type ApiProjectCreator struct {
	api     Api
	cache   *Cache
	options *Options
}

func NewProjectCreator(api Api, cache *Cache, options *Options) ProjectCreator {
	return &ApiProjectCreator{
		api:     api,
		cache:   cache,
		options: options,
	}
}

func (c ApiProjectCreator) CreateProject(ticket *ticket.Ticket) (*github.Project, error) {
	fields := c.createFields(ticket)
	project, err := c.createProject(ticket)
	if err != nil {
		return nil, err
	}

	err = c.createColumns(fields, project)

	if err != nil {
		return nil, err
	}

	return project, nil
}

func (c ApiProjectCreator) createFields(ticket *ticket.Ticket) *Fields {
	return NewFields(c.cache, ticket)
}

func (c ApiProjectCreator) createProject(ticket *ticket.Ticket) (*github.Project, error) {
	projectOptions := c.options.Project(ticket)
	project, err := c.api.CreateProject(projectOptions)
	if err != nil {
		return nil, err
	}

	c.cache.Project.AddProject(project)

	return project, nil
}

func (c ApiProjectCreator) createColumns(fields *Fields, project *github.Project) error {
	columnNames := fields.ProjectColumns()
	for _, name := range columnNames {
		columnOptions := c.options.ProjectColumn(name)
		column, err := c.api.CreateProjectColumn(columnOptions, project)
		if err != nil {
			return err
		}

		columnCache, err := c.cache.Project.ColumnCacheById(*project.ID)

		if err != nil {
			return err
		}

		columnCache.AddColumn(column)
	}
	return nil
}

func (c ApiProjectCreator) UpdateProject(ticket *ticket.Ticket, updateKey string) (*github.Project, error) {
	projectNumber, err := strconv.Atoi(updateKey)
	if err != nil {
		return nil, err
	}

	project, err := c.cache.Project.ProjectByNumber(projectNumber)
	if err != nil {
		return nil, err
	}
	if project == nil {
		return nil, errors.New(fmt.Sprintf("No project #%d", projectNumber))
	}

	projectOptions := c.options.Project(ticket)
	updatedProject, err := c.api.UpdateProject(project, projectOptions)
	if err != nil {
		return nil, err
	}

	c.cache.Project.AddProject(updatedProject)

	return project, nil
}