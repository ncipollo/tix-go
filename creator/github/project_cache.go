package github

import "github.com/google/go-github/v29/github"

type ProjectCache struct {
	api             Api
	columnCacheById map[int64]*ColumnCache
	projectsByNumber    map[int]*github.Project
}

func NewProjectCache(api Api) *ProjectCache {
	return &ProjectCache{
		api:             api,
		columnCacheById: make(map[int64]*ColumnCache),
		projectsByNumber:    make(map[int]*github.Project),
	}
}

func (p *ProjectCache) AddProject(project *github.Project) {
	if project == nil || project.ID == nil {
		return
	}

	p.columnCacheById[*project.ID] = NewColumnCache(p.api, project)
	p.projectsByNumber[*project.Number] = project
}

func (p *ProjectCache) ColumnCacheById(id int64) (*ColumnCache, error) {
	cache := p.columnCacheById[id]
	if cache != nil {
		return cache, nil
	}

	err := p.populateCache()
	if err != nil {
		return nil, err
	}

	return p.columnCacheById[id], nil
}

func (p *ProjectCache) ProjectByNumber(number int) (*github.Project, error) {
	project := p.projectsByNumber[number]
	if project != nil {
		return project, nil
	}

	err := p.populateCache()
	if err != nil {
		return nil, err
	}

	return p.projectsByNumber[number], nil
}

func (p *ProjectCache) populateCache() error {
	projects, err := p.api.ListRepoProjects()
	if err != nil {
		return err
	}
	for _, project := range projects {
		p.AddProject(project)
	}
	return nil
}