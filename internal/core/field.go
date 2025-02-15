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
	fieldName  string
	fieldValue string
	fieldType  FieldType
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
func (field *Field) Format() {
	switch field.fieldType {
	case Number:
		return
	case String:
		return
	case Timestamp:
		formattedTime, err := time.Parse(time.RFC1123Z, field.fieldValue)
		if err != nil {
			return
		}
		field.fieldValue = formattedTime.String()
	case JSON:
		formattedJson := bytes.Buffer{}
		err := json.Indent(&formattedJson, []byte(field.fieldValue), "", "  ")
		if err != nil {
			return
		}
		field.fieldValue = formattedJson.String()
	default:
		return
	}
}
