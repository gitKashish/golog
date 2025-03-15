package models

import (
	"regexp"
)

// TemplateLiterals store the template literals representing the source and target log formats
type TemplateLiterals struct {
	Source string `yaml:"sourceTemplate"`
	Target string `yaml:"targetTemplate"`
}

// Template represents a template for parsing log entries
type Template struct {
	Literals         TemplateLiterals
	SourceRegex      *regexp.Regexp // Compiled regex for source log format
	SourceFieldRegex *regexp.Regexp
	TargetFieldRegex *regexp.Regexp
	FieldNames       map[string]bool
	Fields           []*Field // List of fields
}

// NewTemplate creates a new empty template
func NewTemplate() *Template {
	return &Template{
		FieldNames: make(map[string]bool),
		Fields:     []*Field{},
	}
}
