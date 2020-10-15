package github

import (
	"errors"
	"github.com/google/go-github/v29/github"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestColumnCache_GetByName_CacheHit(t *testing.T) {
	api := newMockApi()
	cache := createColumnCache(api)

	column := createColumn(1, "name")
	cache.AddColumn(column)
	result, err := cache.GetByName("name")

	assert.Equal(t, column, result)
	assert.NoError(t, err)
}

func TestColumnCache_GetByName_PopulatesCache_Error(t *testing.T) {
	api := newMockApi()
	cache := createColumnCache(api)
	api.On("ListProjectColumns", cache.project).Return(nil, errors.New("error"))

	result, err := cache.GetByName("name")

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestColumnCache_GetByName_PopulatesCache_Success(t *testing.T) {
	api := newMockApi()
	cache := createColumnCache(api)
	columns := []*github.ProjectColumn{createColumn(1, "name")}

	api.On("ListProjectColumns", cache.project).Return(columns, nil)

	result, err := cache.GetByName("name")

	assert.Equal(t, columns[0], result)
	assert.NoError(t, err)
}

func TestColumnCache_GetByName_CreatesColumn_Error(t *testing.T) {
	api := newMockApi()
	cache := createColumnCache(api)
	columns := []*github.ProjectColumn{createColumn(1, "name")}
	newName := "name2"
	columnOptions := &github.ProjectColumnOptions{Name: newName}
	api.On("ListProjectColumns", cache.project).Return(columns, nil)
	api.On("CreateProjectColumn", columnOptions, cache.project).Return(nil, errors.New("err"))

	result, err := cache.GetByName(newName)

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestColumnCache_GetByName_CreatesColumn_Success(t *testing.T) {
	api := newMockApi()
	cache := createColumnCache(api)
	columns := []*github.ProjectColumn{createColumn(1, "name")}
	newName := "name2"
	newColumn := &github.ProjectColumn{Name: &newName}
	columnOptions := &github.ProjectColumnOptions{Name: newName}
	api.On("ListProjectColumns", cache.project).Return(columns, nil)
	api.On("CreateProjectColumn", columnOptions, cache.project).Return(newColumn, nil)

	result, err := cache.GetByName(newName)

	assert.Equal(t, newColumn, result)
	assert.NoError(t, err)
}

func createColumnCache(api Api) *ColumnCache {
	id := int64(1)
	project := &github.Project{ID: &id}
	return NewColumnCache(api, project)
}

func createColumn(id int64, name string) *github.ProjectColumn {
	return &github.ProjectColumn{
		ID:   &id,
		Name: &name,
	}
}
