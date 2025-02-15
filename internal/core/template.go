package core

import (
	"fmt"
	"os"
	"regexp"

	"github.com/gitkashish/golog/internal/helpers"
	"gopkg.in/yaml.v3"
)

// Store the template literals representing the source and target log formats
type templateLiterals struct {
	Source string `yaml:"sourceTemplate"`
	Target string `yaml:"targetTemplate"`
}

// Represents a template for parsing log entries
type Template struct {
	literals    templateLiterals
	sourceRegex *regexp.Regexp // Compiled regex for source log format

	sourceFieldRegex *regexp.Regexp
	targetFieldRegex *regexp.Regexp

	fieldNames map[string]bool
	Fields     []*Field // List of fields
}

// Function to parse log
func (template *Template) Parse(sourceLog string) error {
	// Retrieve field values from source log
	fields := template.sourceRegex.FindStringSubmatch(sourceLog)
	// Check if log matches format or not
	if fields == nil {
		return fmt.Errorf("log does not match source pattern")
	}

	// Setting field values to corresponding capture group values
	count := 1
	for _, field := range template.Fields {
		field.fieldValue = fields[count]
		field.Format()
		count++
		if count >= len(fields) {
			break
		}
	}
	return nil
}

func (template *Template) Execute() string {

	formattedLog := template.literals.Target
	for _, field := range template.Fields {
		fieldRegex := regexp.MustCompile("@" + field.fieldName + "@")
		formattedLog = fieldRegex.ReplaceAllString(formattedLog, field.fieldValue)
	}
	return formattedLog
}

// Function to load a template from template.yaml file
func loadTemplateLiterals() (templateLiterals, error) {
	literals := templateLiterals{}
	// Check if template.yaml exists
	err := helpers.IsFile("template.yaml")
	if err != nil {
		return literals, err
	}

	// Read template.yaml file
	data, err := os.ReadFile("template.yaml")
	if err != nil {
		return literals, err
	}

	// Unmarshal the yaml data into templateLiterals struct
	err = yaml.Unmarshal(data, &literals)
	if err != nil {
		return literals, err
	}

	return literals, nil
}

// Function to parse template into a field map
func (template *Template) parseSourceTemplate() {
	// Extract fields from literal
	matches := template.sourceFieldRegex.FindAllStringSubmatch(template.literals.Source, -1)
	for _, match := range matches {
		// Check if a field name was already mentioned log
		// Exit if field name already exists
		if template.fieldNames[match[1]] {
			fmt.Printf("duplicate field name `%s` in source template\n", match[1])
			os.Exit(1)
		}
		template.fieldNames[match[1]] = true
		fieldType, err := FieldTypeFromString(match[2]) // Get corresponding FieldType
		// Exit if unkown FieldType
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		template.Fields = append(template.Fields, &Field{match[1], "", fieldType}) // Add new field
	}
}

// Function to parse target template and check if it is valid or not
func (template *Template) parseTargetTemplate() {
	// Extract fields from literal
	matches := template.targetFieldRegex.FindAllStringSubmatch(template.literals.Target, -1)
	for _, match := range matches {
		if !template.fieldNames[match[1]] {
			fmt.Printf("field `%s` not found in source template\n", match)
			os.Exit(1)
		}
	}
}

// Function to get regex for extracting field values from
func getSourceRegex(sourceTemplate string, sourceFieldRegex regexp.Regexp) *regexp.Regexp {
	sourceTemplate = regexp.QuoteMeta(sourceTemplate) // Escape all regex meta-characters

	regexPattern := sourceFieldRegex.ReplaceAllString(sourceTemplate, `(.+)`) // Replace Field Regex Pattern with `(.+)`

	spaceRegex := regexp.MustCompile(`\s+`)
	regexPattern = spaceRegex.ReplaceAllString(regexPattern, `\s*`) // Replace all whitespace with `\s+`

	sourceRegex := regexp.MustCompile("^" + regexPattern + "$") // Compile source regex with anchoring to entire line
	return sourceRegex
}

// Function to get the template for parsing log entries
func GetTemplate() (*Template, error) {
	literals, err := loadTemplateLiterals()
	if err != nil {
		return nil, err
	}

	template := Template{
		literals:   literals,
		fieldNames: make(map[string]bool),
	}

	// Compile regex for fields in source and target template
	template.sourceFieldRegex = regexp.MustCompile(`@(\w+)-(\w+)@`)
	template.targetFieldRegex = regexp.MustCompile(`@(\w+)@`)

	// Parse template literal and retrieve fields
	template.parseSourceTemplate()
	template.parseTargetTemplate()

	// Get regex to parse logs
	template.sourceRegex = getSourceRegex(template.literals.Source, *template.sourceFieldRegex)
	return &template, nil
}
