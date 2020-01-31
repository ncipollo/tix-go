package md

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
	return f.mergeFields(f.defaultFields, f.fieldsByLevel[level])
}

func (f *FieldState) SetDefaultFields(defaultFields map[string]interface{}) {
	f.defaultFields = defaultFields
}

func (f *FieldState) SetFieldsForLevel(fields map[string]interface{}, level int) {
	f.fieldsByLevel[level] = fields
}

func (f FieldState) mergeFields(
	baseFields map[string]interface{},
	overlayFields map[string]interface{}) map[string]interface{} {
	combinedFields := make(map[string]interface{})
	if baseFields != nil {
		for key, value := range baseFields {
			combinedFields[key] = value
		}
	}
	if overlayFields != nil {
		for key, value := range overlayFields {
			combinedFields[key] = value
		}
	}

	return combinedFields
}
