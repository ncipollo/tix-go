package md

import "tix/ticket"

type FieldState struct {
	defaultFields map[string]interface{}
	fieldsByLevel map[int]map[string]interface{}
}

func NewFieldState() *FieldState {
	return &FieldState{
		defaultFields: make(map[string]interface{}),
		fieldsByLevel: make(map[int]map[string]interface{}),
	}
}

func (f FieldState) FieldsForLevel(level int) map[string]interface{} {
	return ticket.MergeFields(f.defaultFields, f.fieldsByLevel[level])
}

func (f *FieldState) SetDefaultFields(defaultFields map[string]interface{}) {
	f.defaultFields = defaultFields
}

func (f *FieldState) SetFieldsForLevel(fields map[string]interface{}, level int) {
	f.fieldsByLevel[level] = fields
}
