package models

import (
	"strings"
	"testing"
)

func TestFieldTypeFromString(t *testing.T) {
	tests := []struct {
		name     string
		typeName string
		want     FieldType
		wantErr  bool
	}{
		{
			name:     "Raw type",
			typeName: "raw",
			want:     Raw,
			wantErr:  false,
		},
		{
			name:     "Number type",
			typeName: "number",
			want:     Number,
			wantErr:  false,
		},
		{
			name:     "String type",
			typeName: "string",
			want:     String,
			wantErr:  false,
		},
		{
			name:     "Timestamp type",
			typeName: "timestamp",
			want:     Timestamp,
			wantErr:  false,
		},
		{
			name:     "JSON type",
			typeName: "json",
			want:     JSON,
			wantErr:  false,
		},
		{
			name:     "Invalid type",
			typeName: "invalid",
			want:     -1,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FieldTypeFromString(tt.typeName)
			if (err != nil) != tt.wantErr {
				t.Errorf("FieldTypeFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FieldTypeFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldType_String(t *testing.T) {
	tests := []struct {
		name      string
		fieldType FieldType
		want      string
	}{
		{
			name:      "Raw type",
			fieldType: Raw,
			want:      "raw",
		},
		{
			name:      "Number type",
			fieldType: Number,
			want:      "number",
		},
		{
			name:      "String type",
			fieldType: String,
			want:      "string",
		},
		{
			name:      "Timestamp type",
			fieldType: Timestamp,
			want:      "timestamp",
		},
		{
			name:      "JSON type",
			fieldType: JSON,
			want:      "json",
		},
		{
			name:      "Invalid type",
			fieldType: -1,
			want:      "unknown",
		},
		{
			name:      "Out of range type",
			fieldType: 100,
			want:      "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fieldType.String(); got != tt.want {
				t.Errorf("FieldType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestField_Format(t *testing.T) {
	tests := []struct {
		name  string
		field *Field
		value string
		want  string
	}{
		{
			name:  "Format raw value",
			field: &Field{Name: "test", Type: Raw},
			value: "raw value",
			want:  "raw value",
		},
		{
			name:  "Format number value",
			field: &Field{Name: "test", Type: Number},
			value: "123",
			want:  "123",
		},
		{
			name:  "Format string value",
			field: &Field{Name: "test", Type: String},
			value: "string value",
			want:  "string value",
		},
		{
			name:  "Format timestamp value",
			field: &Field{Name: "test", Type: Timestamp},
			value: "Wed, 15 Mar 2023 14:30:45 +0000",
			want:  "2023-03-15 14:30:45 +0000 UTC",
		},
		{
			name:  "Format invalid timestamp value",
			field: &Field{Name: "test", Type: Timestamp},
			value: "invalid timestamp",
			want:  "invalid timestamp",
		},
		{
			name:  "Format JSON value",
			field: &Field{Name: "test", Type: JSON},
			value: "{\"user\":\"john\",\"id\":123}",
			want:  "{\n  \"user\": \"john\",\n  \"id\": 123\n}",
		},
		{
			name:  "Format invalid JSON value",
			field: &Field{Name: "test", Type: JSON},
			value: "invalid json",
			want:  "invalid json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.field.Format(tt.value)

			if tt.name == "Format timestamp value" {
				// For timestamp, just check that it contains the expected date and time parts
				if !strings.Contains(result, "2023-03-15") || !strings.Contains(result, "14:30:45") {
					t.Errorf("Field.Format() = %v, does not contain expected date/time", result)
				}
			} else if tt.name == "Format JSON value" {
				// For JSON, check that it contains the expected JSON content
				if !strings.Contains(result, "\"user\": \"john\"") || !strings.Contains(result, "\"id\": 123") {
					t.Errorf("Field.Format() = %v, does not contain expected JSON content", result)
				}
			} else if result != tt.want {
				t.Errorf("Field.Format() = %v, want %v", result, tt.want)
			}
		})
	}
}
