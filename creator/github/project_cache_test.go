package github

import (
	"errors"
	"github.com/google/go-github/v29/github"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProjectCache_ColumnCacheById_CacheHit(t *testing.T) {
	api := newMockApi()
	number := 1
	id := int64(number)
	project := &github.Project{ID: &id, Number: &number}
	cache := NewProjectCache(api)

	cache.AddProject(project)
	result, err := cache.ColumnCacheById(id)

	expected := NewColumnCache(api, project)
	assert.Equal(t, expected, result)
	assert.NoError(t, err)
}

func TestProjectCache_ColumnCacheById_CacheMiss(t *testing.T) {
	api := newMockApi()
	number := 1
	id := int64(number)
	cache := NewProjectCache(api)
	api.On("ListRepoProjects").Return([]*github.Project{}, nil)

	result, err := cache.ColumnCacheById(id)

	assert.Nil(t, result)
	assert.NoError(t, err)
}

func TestProjectCache_ColumnCacheById_PopulatesCache_Error(t *testing.T) {
	api := newMockApi()
	number := 1
	id := int64(number)
	cache := NewProjectCache(api)
	err := errors.New("i am err")
	api.On("ListRepoProjects").Return(nil, err)

	result, err := cache.ColumnCacheById(id)

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestProjectCache_ColumnCacheById_PopulatesCache_Success(t *testing.T) {
	api := newMockApi()
	number := 1
	id := int64(number)
	project := &github.Project{ID: &id, Number: &number}
	cache := NewProjectCache(api)
	api.On("ListRepoProjects").Return([]*github.Project{project}, nil)

	result, err := cache.ColumnCacheById(id)

	expected := NewColumnCache(api, project)
	assert.Equal(t, expected, result)
	assert.NoError(t, err)
}

func TestProjectCache_ColumnCacheById_PreservesColumnCacheWhenProjectUpdated(t *testing.T) {
	api := newMockApi()
	number := 1
	id := int64(number)
	name := "name"
	project := &github.Project{ID: &id, Number: &number}
	cache := NewProjectCache(api)
	column := &github.ProjectColumn{ID: &id, Name: &name}

	cache.AddProject(project)
	columnCache, _ := cache.ColumnCacheById(id)
	columnCache.AddColumn(column)
	cache.AddProject(project)

	columnCache, _ = cache.ColumnCacheById(id)
	result, err := columnCache.GetByName(name)

	assert.Equal(t, column, result)
	assert.NoError(t, err)
}

func TestProjectCache_ProjectByNumber_CacheMiss(t *testing.T) {
	api := newMockApi()
	number := 1
	cache := NewProjectCache(api)
	api.On("ListRepoProjects").Return([]*github.Project{}, nil)

	result, err := cache.ProjectByNumber(number)

	assert.Nil(t, result)
	assert.NoError(t, err)
}

func TestProjectCache_ProjectByNumber_CacheHit(t *testing.T) {
	api := newMockApi()
	number := 1
	id := int64(2)
	project := &github.Project{ID: &id, Number: &number}
	cache := NewProjectCache(api)

	cache.AddProject(project)
	result, err := cache.ProjectByNumber(number)

	assert.Equal(t, project, result)
	assert.NoError(t, err)
}

func TestProjectCache_ProjectByNumber_PopulatesCache_Error(t *testing.T) {
	api := newMockApi()
	number := 1
	id := int64(2)
	project := &github.Project{ID: &id, Number: &number}
	cache := NewProjectCache(api)
	api.On("ListRepoProjects").Return([]*github.Project{project}, nil)

	result, err := cache.ProjectByNumber(number)

	assert.Equal(t, project, result)
	assert.NoError(t, err)
}

func TestProjectCache_ProjectByNumber_PopulatesCache_Success(t *testing.T) {
	api := newMockApi()
	number := 1
	cache := NewProjectCache(api)
	err := errors.New("i am err")
	api.On("ListRepoProjects").Return(nil, err)

	result, err := cache.ProjectByNumber(number)

	assert.Nil(t, result)
	assert.Error(t, err)
}
