package jira

import (
	"github.com/andygrunwald/go-jira"
	"strings"
	"tix/ticket"
)

const (
	KeyAffectsVersions = "affects versions"
	KeyComponents      = "components"
	KeyFixVersions     = "fix versions"
	KeyLabels          = "labels"
	KeyParent          = "parent"
	KeyProject         = "project"
	KeyType            = "type"
	KeyUseParent       = "use_parent"
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

func (i *IssueFields) AddDefaultEpicName(unknowns map[string]interface{}, name string) {
	epicNameKey := i.fieldInfo["epic name"].ID
	epicName, ok := unknowns[epicNameKey].(string)
	if !ok || len(epicName) == 0 {
		unknowns[epicNameKey] = name
	}
}

func (i *IssueFields) AffectsVersions() []*jira.AffectsVersion {
	rawVersions, ok := i.ticket.Fields("jira")[KeyAffectsVersions].([]interface{})
	if ok {
		versions := make([]*jira.AffectsVersion, 0)
		for _, label := range rawVersions {
			versionString, ok := label.(string)
			if ok {
				versions = append(versions, &jira.AffectsVersion{Name: versionString})
			}
		}
		return versions
	} else {
		return nil
	}
}

func (i *IssueFields) Components() []*jira.Component {
	components, ok := i.ticket.Fields("jira")[KeyComponents].([]interface{})
	if !ok {
		return nil
	}
	jiraComps := make([]*jira.Component, 0)
	for _, component := range components {
		compString, ok := component.(string)
		if !ok {
			compString = ""
		}
		jiraComp := &jira.Component{Name: compString}
		jiraComps = append(jiraComps, jiraComp)
	}
	return jiraComps
}

func (i *IssueFields) EpicLinkKey() string {
	return i.fieldInfo["epic link"].ID
}

func (i *IssueFields) EpicType() jira.IssueType {
	issueType, ok := i.ticket.Fields("jira")[KeyType].(string)
	if !ok {
		issueType = "Epic"
	}
	return jira.IssueType{
		Name: issueType,
	}
}

func (i *IssueFields) UseParent() bool {
	useParent, ok := i.ticket.Fields("jira")[KeyUseParent].(bool)
	if ok {
		return useParent
	} else {
		return false
	}
}

func (i *IssueFields) FixVersions() []*jira.FixVersion {
	rawVersions, ok := i.ticket.Fields("jira")[KeyFixVersions].([]interface{})
	if ok {
		versions := make([]*jira.FixVersion, 0)
		for _, label := range rawVersions {
			versionString, ok := label.(string)
			if ok {
				versions = append(versions, &jira.FixVersion{Name: versionString})
			}
		}
		return versions
	} else {
		return nil
	}
}

func (i *IssueFields) Labels() []string {
	labels, ok := i.ticket.Fields("jira")[KeyLabels].([]interface{})
	if ok {
		labelStrings := make([]string, 0)
		for _, label := range labels {
			labelString, ok := label.(string)
			if ok {
				labelStrings = append(labelStrings, labelString)
			}
		}
		return labelStrings
	} else {
		return make([]string, 0)
	}
}

func (i *IssueFields) IssueType() jira.IssueType {
	issueType, ok := i.ticket.Fields("jira")[KeyType].(string)
	if !ok {
		issueType = "Story"
	}
	return jira.IssueType{
		Name: issueType,
	}
}

func (i *IssueFields) Parent() *jira.Parent {
	parentKey, ok := i.ticket.Fields("jira")[KeyParent].(string)
	if !ok {
		return nil
	}

	return &jira.Parent{Key: parentKey}
}

func (i *IssueFields) Project() jira.Project {
	project, ok := i.ticket.Fields("jira")[KeyProject].(string)
	if !ok {
		project = ""
	}
	return jira.Project{
		Key: project,
	}
}

func (i *IssueFields) TaskType() jira.IssueType {
	issueType, ok := i.ticket.Fields("jira")[KeyType].(string)
	if !ok {
		issueType = "Task"
	}
	return jira.IssueType{
		Name: issueType,
	}
}

func (i *IssueFields) Unknowns() map[string]interface{} {
	keysToSkip := map[string]bool{
		KeyAffectsVersions:     true,
		KeyComponents:          true,
		KeyFixVersions:         true,
		KeyLabels:              true,
		KeyParent:              true,
		KeyProject:             true,
		KeyType:                true,
		ticket.KeyUpdateTicket: true,
		KeyUseParent:           true,
	}

	unknown := make(map[string]interface{})
	for key, field := range i.ticket.Fields("jira") {
		lowerKey := strings.ToLower(key)
		if keysToSkip[lowerKey] {
			continue
		}

		info := i.fieldInfo[lowerKey]
		if info != nil {
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
