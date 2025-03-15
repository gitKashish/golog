package parser

import (
	"fmt"
	"os"
	"regexp"

	"github.com/gitKashish/golog/internal/core/models"
	"gopkg.in/yaml.v3"
)

// LogParser defines the interface for parsing logs
type LogParser interface {
	// Parse parses a log entry using the template
	Parse(sourceLog string) string
	// LoadTemplate loads a template from a file
	LoadTemplate(templatePath string) error
	// SetTemplate sets the template directly
	SetTemplate(sourceTemplate, targetTemplate string) error
}

// TemplateParser implements the LogParser interface
type TemplateParser struct {
	template *models.Template
}

// NewTemplateParser creates a new TemplateParser
func NewTemplateParser() *TemplateParser {
	return &TemplateParser{
		template: models.NewTemplate(),
	}
}

// Parse parses a log entry using the template
func (p *TemplateParser) Parse(sourceLog string) string {
	if p.template == nil || p.template.SourceRegex == nil {
		return sourceLog
	}

	// Retrieve field values from source log
	fields := p.template.SourceRegex.FindStringSubmatch(sourceLog)

	// Check if log matches format or not
	if len(fields) <= 1 {
		return sourceLog
	}

	// Map Field Names and their corresponding values
	fieldValues := map[string]string{}
	count := 1
	for _, field := range p.template.Fields {
		if count < len(fields) {
			fieldValues[field.Name] = field.Format(fields[count])
			count++
		}
	}

	formattedLog := p.template.Literals.Target
	for _, field := range p.template.Fields {
		fieldRegex := regexp.MustCompile("@" + field.Name + "@")
		if value, ok := fieldValues[field.Name]; ok {
			formattedLog = fieldRegex.ReplaceAllString(formattedLog, value)
		}
	}
	return formattedLog
}

// LoadTemplate loads a template from a file
func (p *TemplateParser) LoadTemplate(templatePath string) error {
	literals, err := loadTemplateLiterals(templatePath)
	if err != nil {
		return err
	}

	return p.setupTemplate(literals)
}

// SetTemplate sets the template directly
func (p *TemplateParser) SetTemplate(sourceTemplate, targetTemplate string) error {
	literals := models.TemplateLiterals{
		Source: sourceTemplate,
		Target: targetTemplate,
	}

	return p.setupTemplate(literals)
}

// setupTemplate sets up the template with the given literals
func (p *TemplateParser) setupTemplate(literals models.TemplateLiterals) error {
	p.template = models.NewTemplate()
	p.template.Literals = literals

	// Compile regex for fields in source and target template
	p.template.SourceFieldRegex = regexp.MustCompile(`@(\w+)-(\w+)@`)
	p.template.TargetFieldRegex = regexp.MustCompile(`@(\w+)@`)

	// Parse template literal and retrieve fields
	if err := p.parseSourceTemplate(); err != nil {
		return err
	}

	if err := p.parseTargetTemplate(); err != nil {
		return err
	}

	// Get regex to parse logs
	p.template.SourceRegex = getSourceRegex(p.template.Literals.Source, *p.template.SourceFieldRegex)
	return nil
}

// parseSourceTemplate parses the source template and extracts fields
func (p *TemplateParser) parseSourceTemplate() error {
	// Extract fields from literal
	matches := p.template.SourceFieldRegex.FindAllStringSubmatch(p.template.Literals.Source, -1)
	for _, match := range matches {
		// Check if a field name was already mentioned in log
		if p.template.FieldNames[match[1]] {
			return fmt.Errorf("duplicate field name `%s` in source template", match[1])
		}
		p.template.FieldNames[match[1]] = true
		fieldType, err := models.FieldTypeFromString(match[2]) // Get corresponding FieldType
		if err != nil {
			return err
		}
		p.template.Fields = append(p.template.Fields, &models.Field{Name: match[1], Type: fieldType}) // Add new field
	}
	return nil
}

// parseTargetTemplate parses the target template and validates it
func (p *TemplateParser) parseTargetTemplate() error {
	// Extract fields from literal
	matches := p.template.TargetFieldRegex.FindAllStringSubmatch(p.template.Literals.Target, -1)
	for _, match := range matches {
		if !p.template.FieldNames[match[1]] {
			return fmt.Errorf("field `%s` not found in source template", match[1])
		}
	}
	return nil
}

// getSourceRegex creates a regex for extracting field values from source logs
func getSourceRegex(sourceTemplate string, sourceFieldRegex regexp.Regexp) *regexp.Regexp {
	sourceTemplate = regexp.QuoteMeta(sourceTemplate) // Escape all regex meta-characters

	// Replace field patterns with capture groups
	regexPattern := sourceFieldRegex.ReplaceAllString(sourceTemplate, `(.*?)`)

	// Replace whitespace with flexible whitespace pattern
	spaceRegex := regexp.MustCompile(`\s+`)
	regexPattern = spaceRegex.ReplaceAllString(regexPattern, `\s+`)

	// Compile the regex with start and end anchors
	sourceRegex := regexp.MustCompile("^" + regexPattern + "$")
	return sourceRegex
}

// loadTemplateLiterals loads template literals from a file
func loadTemplateLiterals(templatePath string) (models.TemplateLiterals, error) {
	literals := models.TemplateLiterals{}

	// Read template file
	data, err := os.ReadFile(templatePath)
	if err != nil {
		return literals, fmt.Errorf("error reading template file: %w", err)
	}

	// Unmarshal the yaml data into templateLiterals struct
	err = yaml.Unmarshal(data, &literals)
	if err != nil {
		return literals, fmt.Errorf("error parsing template file: %w", err)
	}

	return literals, nil
}
