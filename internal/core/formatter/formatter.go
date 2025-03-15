package formatter

import (
	"github.com/gitKashish/golog/internal/core/parser"
)

// LogFormatter defines the interface for formatting logs
type LogFormatter interface {
	// FormatLog formats a single log entry
	FormatLog(log string) string
	// FormatLogs formats multiple log entries
	FormatLogs(logs []string) []string
	// SetParser sets the parser to use for formatting
	SetParser(parser parser.LogParser)
}

// TemplateFormatter implements the LogFormatter interface
type TemplateFormatter struct {
	parser parser.LogParser
}

// NewTemplateFormatter creates a new TemplateFormatter
func NewTemplateFormatter(parser parser.LogParser) *TemplateFormatter {
	return &TemplateFormatter{
		parser: parser,
	}
}

// FormatLog formats a single log entry
func (f *TemplateFormatter) FormatLog(log string) string {
	return f.parser.Parse(log)
}

// FormatLogs formats multiple log entries
func (f *TemplateFormatter) FormatLogs(logs []string) []string {
	formattedLogs := make([]string, len(logs))
	for i, log := range logs {
		formattedLogs[i] = f.FormatLog(log)
	}
	return formattedLogs
}

// SetParser sets the parser to use for formatting
func (f *TemplateFormatter) SetParser(parser parser.LogParser) {
	f.parser = parser
}
