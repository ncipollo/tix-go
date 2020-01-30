package jira

import (
	"github.com/andygrunwald/go-jira"
	"strings"
	"tix/ticket"
)

const (
	KeyComponents  = "components"
	KeyLabels      = "labels"
	KeyProject     = "project"
	KeyType        = "type"
)

type FieldInfo struct {
	ID          string
	Name        string
	UseValueKey bool
}

func fieldInfoByName(jiraFields []jira.Field) map[string]*FieldInfo {
	fieldMap := make(map[string]*FieldInfo)
	for _, jiraField := range jiraFields {
		info := &FieldInfo{
			ID:          jiraField.ID,
			Name:        jiraField.Name,
			UseValueKey: jiraField.Schema.Type == "option", // There may also be others
		}
		key := strings.ToLower(info.Name)
		fieldMap[key] = info
	}
	return fieldMap
}

type IssueFields struct {
	ticket    *ticket.Ticket
	fieldInfo map[string]*FieldInfo
}

func NewIssueFields(jiraFields []jira.Field, ticket *ticket.Ticket) *IssueFields {
	return &IssueFields{
		ticket:    ticket,
		fieldInfo: fieldInfoByName(jiraFields),
	}
}

func (i *IssueFields) Components() []*jira.Component {
	components, ok := i.ticket.Fields[KeyComponents].([]string)
	if !ok {
		components = make([]string, 0)
	}
	var jiraComps []*jira.Component
	for _, component := range components {
		jiraComp := &jira.Component{Name: component}
		jiraComps = append(jiraComps, jiraComp)
	}
	return jiraComps
}

func (i *IssueFields) EpicType() jira.IssueType {
	issueType, ok := i.ticket.Fields[KeyType].(string)
	if !ok {
		issueType = "Epic"
	}
	return jira.IssueType{
		Name: issueType,
	}
}

func (i *IssueFields) Labels() []string {
	return i.ticket.Fields[KeyLabels].([]string)
}

func (i *IssueFields) IssueType() jira.IssueType {
	issueType, ok := i.ticket.Fields[KeyType].(string)
	if !ok {
		issueType = "Story"
	}
	return jira.IssueType{
		Name: issueType,
	}
}

func (i *IssueFields) Project() jira.Project {
	project := i.ticket.Fields[KeyProject].(string)
	return jira.Project{
		Key: project,
	}
}

func (i *IssueFields) Unknowns() map[string]interface{} {
	keysToSkip := map[string]bool{
		KeyComponents:  true,
		KeyLabels:      true,
		KeyProject:     true,
		KeyType:        true,
	}

	unknown := make(map[string]interface{})
	for key, field := range i.ticket.Fields {
		lowerKey := strings.ToLower(key)
		if keysToSkip[lowerKey] {
			continue
		}

		info := i.fieldInfo[lowerKey]
		if info != nil{
			if info.UseValueKey {
				unknown[info.ID] = map[string]interface{}{"value": field}
			} else {
				unknown[info.ID] = field
			}
		} else {
			unknown[key] = field
		}
	}
	return unknown
}