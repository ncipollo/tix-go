package github

import "github.com/google/go-github/v29/github"

type ColumnCache struct {
	api           Api
	columnsByName map[string]*github.ProjectColumn
	project       *github.Project
}

func NewColumnCache(api Api, project *github.Project) *ColumnCache {
	return &ColumnCache{
		api:           api,
		columnsByName: make(map[string]*github.ProjectColumn),
		project:       project,
	}
}

func (c *ColumnCache) AddColumn(column *github.ProjectColumn) {
	if column == nil || column.Name == nil {
		return
	}
	c.columnsByName[*column.Name] = column
}

func (c *ColumnCache) GetByName(name string) (*github.ProjectColumn, error) {
	column := c.columnsByName[name]
	if column != nil {
		return column, nil
	}

	err := c.populateCache()
	if err != nil {
		return nil, err
	}

	column = c.columnsByName[name]
	if column != nil {
		return column, nil
	}

	return c.createColumn(name)
}

func (c *ColumnCache) populateCache() error {
	milestones, err := c.api.ListProjectColumns(c.project)
	if err != nil {
		return err
	}
	for _, columns := range milestones {
		c.AddColumn(columns)
	}
	return nil
}

func (c *ColumnCache) createColumn(name string) (*github.ProjectColumn, error) {
	columnRequest := &github.ProjectColumnOptions{Name: name}
	column, err := c.api.CreateProjectColumn(columnRequest, c.project)
	if err != nil {
		return nil, err
	}
	c.AddColumn(column)
	return column, nil
}