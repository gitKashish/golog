package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// FieldType represents the data type stored in a field
// like "number", "string", "json", "timestamp" etc.
// Fields are formatted depending on the field type
type FieldType int

type Field struct {
	fieldName string
	fieldType FieldType
}

const (
	Raw FieldType = iota
	Number
	String
	Timestamp
	JSON
)

var fieldTypeMap = map[string]FieldType{
	"raw":       Raw,
	"number":    Number,
	"string":    String,
	"timestamp": Timestamp,
	"json":      JSON,
}

// Get corresponding field type name
func (fieldType FieldType) String() string {
	// Map field type names to corresponding FieldType
	names := [...]string{"raw", "number", "string", "timestamp", "json"}

	// Return "unknown" if fieldType does not have corresponding name
	if fieldType < 0 || int(fieldType) >= len(names) {
		return "unknown"
	}
	return names[fieldType]
}

// Function to get FieldType from corresponding string
func FieldTypeFromString(name string) (FieldType, error) {
	fieldType, ok := fieldTypeMap[name]
	if !ok {
		return -1, fmt.Errorf("invalid field type `%s`", name)
	}
	return fieldType, nil
}

// Function to format a Field based on fieldType
func (field *Field) Format(value string) string {
	switch field.fieldType {
	case Number:
		return value
	case String:
		return value
	case Timestamp:
		formattedTime, err := time.Parse(time.RFC1123Z, value)
		if err != nil {
			return value
		}
		return formattedTime.String()
	case JSON:
		formattedJson := bytes.Buffer{}
		err := json.Indent(&formattedJson, []byte(value), "", "  ")
		if err != nil {
			return value
		}
		return formattedJson.String()
	default:
		return value
	}
}
