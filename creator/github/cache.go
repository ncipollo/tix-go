package github

type Cache struct {
	Milestone *MilestoneCache
	Project *ProjectCache
}

func NewCache(api Api) *Cache {
	return &Cache{
		Milestone: NewMilestoneCache(api),
		Project:   NewProjectCache(api),
	}
}