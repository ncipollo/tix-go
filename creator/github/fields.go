package github

import (
	"errors"
	"fmt"
	"github.com/google/go-github/v29/github"
	"tix/ticket"
)

const (
	KeyAssignee = "assignee"
	KeyAssignees = "assignees"
	KeyColumn    = "column"
	KeyColumns   = "columns"
	KeyLabels    = "labels"
	KeyMilestone = "milestone"
	KeyProject   = "project"
)

type Fields struct {
	cache  *Cache
	ticket *ticket.Ticket
}

func NewFields(cache *Cache, ticket *ticket.Ticket) *Fields {
	return &Fields{
		cache:  cache,
		ticket: ticket,
	}
}

func (f *Fields) Assignees() *[]string {
	assignees := f.stringSliceField(KeyAssignees)
	if assignees != nil {
		 return assignees
	}

	assignee, ok := f.ticket.Fields("github")[KeyAssignee].(string)
	if ok {
		return &[]string{assignee}
	}
	return nil
}

func (f *Fields) IssueColumn(projectId int64) (*github.ProjectColumn, error) {
	column, ok := f.ticket.Fields("github")[KeyColumn].(string)
	if !ok {
		column = "To do"
	}

	columnCache, err := f.cache.Project.ColumnCacheById(projectId)

	if err != nil {
		return nil, err
	}

	if columnCache == nil {
		return nil, errors.New(fmt.Sprintf(":scream: No open project with id: %d", projectId))
	}

	return columnCache.GetByName(column)
}

func (f *Fields) Labels() *[]string {
	return f.stringSliceField(KeyLabels)
}

func (f *Fields) stringSliceField(key string) *[]string {
	rawValue, ok := f.ticket.Fields("github")[key].([]interface{})
	if ok {
		values := make([]string, 0, len(rawValue))
		for _, value := range rawValue {
			valueString, ok := value.(string)
			if ok {
				values = append(values, valueString)
			}
		}
		return &values
	} else {
		return nil
	}
}

func (f *Fields) Milestone() (*github.Milestone, error) {
	column, ok := f.ticket.Fields("github")[KeyMilestone].(string)
	if ok {
		return f.cache.Milestone.GetByName(column)
	}
	return nil, nil
}

func (f *Fields) ProjectColumns() []string {
	columns := f.stringSliceField(KeyColumns)
	if columns != nil {
		return *columns
	}
	return []string{"To do", "In progress", "Done"}
}

func (f *Fields) Project() (*github.Project, error) {
	number, ok := f.ticket.Fields("github")[KeyProject].(int)
	if ok {
		project, err := f.cache.Project.ProjectByNumber(number)
		if err != nil {
			return nil, err
		}
		if project == nil {
			return nil, errors.New(fmt.Sprintf(":scream: No open project with number: %d", number))
		}
		return project, nil
	} else {
		return nil, nil
	}
}
