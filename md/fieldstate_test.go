package md

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFieldState_FieldsForLevel_ReturnsMergedFields(t *testing.T) {
	defaultFields := map[string]interface{}{
		"default": "default",
	}
	levelFields := map[string]interface{}{
		"level": "level",
	}
	fieldState := NewFieldState()

	fieldState.SetDefaultFields(defaultFields)
	fieldState.SetFieldsForLevel(levelFields, 1)
	combinedFields := fieldState.FieldsForLevel(1)

	expected := map[string]interface{}{
		"default": "default",
		"level": "level",
	}
	assert.Equal(t, expected, combinedFields)
}

func TestFieldState_FieldsForLevel_ReturnsDefaultsWhenNoLevelFields(t *testing.T) {
	defaultFields := map[string]interface{}{
		"default": "default",
	}
	fieldState := NewFieldState()

	fieldState.SetDefaultFields(defaultFields)
	combinedFields := fieldState.FieldsForLevel(1)

	expected := map[string]interface{}{
		"default": "default",
	}
	assert.Equal(t, expected, combinedFields)
}

func TestFieldState_FieldsForLevel_ReturnsDefaultsWhenProvidedLevelIsEmpty(t *testing.T) {
	defaultFields := map[string]interface{}{
		"default": "default",
	}
	levelFields := map[string]interface{}{
		"level": "level",
	}
	fieldState := NewFieldState()

	fieldState.SetDefaultFields(defaultFields)
	fieldState.SetFieldsForLevel(levelFields, 1)
	combinedFields := fieldState.FieldsForLevel(2)

	expected := map[string]interface{}{
		"default": "default",
	}
	assert.Equal(t, expected, combinedFields)
}

func TestFieldState_FieldsForLevel_ReturnsEmptyFields(t *testing.T) {
	fieldState := NewFieldState()
	combinedFields := fieldState.FieldsForLevel(2)

	expected := map[string]interface{}{}
	assert.Equal(t, expected, combinedFields)
}

func TestFieldState_FieldsForLevel_ReturnLevelFieldsWhenNoDefaults(t *testing.T) {
	levelFields := map[string]interface{}{
		"level": "level",
	}
	fieldState := NewFieldState()

	fieldState.SetFieldsForLevel(levelFields, 1)
	combinedFields := fieldState.FieldsForLevel(1)

	expected := map[string]interface{}{
		"level": "level",
	}
	assert.Equal(t, expected, combinedFields)
}