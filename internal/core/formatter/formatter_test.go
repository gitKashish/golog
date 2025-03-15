package formatter

import (
	"strings"
	"testing"

	"github.com/gitKashish/golog/internal/core/parser"
)

// MockParser is a mock implementation of the LogParser interface
type MockParser struct {
	parseFunc func(string) string
}

func (m *MockParser) Parse(sourceLog string) string {
	return m.parseFunc(sourceLog)
}

func (m *MockParser) LoadTemplate(templatePath string) error {
	return nil
}

func (m *MockParser) SetTemplate(sourceTemplate, targetTemplate string) error {
	return nil
}

func TestTemplateFormatter_FormatLog(t *testing.T) {
	// Create a mock parser that returns a formatted log
	mockParser := &MockParser{
		parseFunc: func(sourceLog string) string {
			return "[FORMATTED] " + sourceLog
		},
	}

	// Create a formatter with the mock parser
	formatter := NewTemplateFormatter(mockParser)

	// Test formatting a single log
	sourceLog := "This is a log message"
	expected := "[FORMATTED] This is a log message"
	result := formatter.FormatLog(sourceLog)
	if result != expected {
		t.Errorf("FormatLog() = %v, want %v", result, expected)
	}
}

func TestTemplateFormatter_FormatLogs(t *testing.T) {
	// Create a mock parser that returns a formatted log
	mockParser := &MockParser{
		parseFunc: func(sourceLog string) string {
			return "[FORMATTED] " + sourceLog
		},
	}

	// Create a formatter with the mock parser
	formatter := NewTemplateFormatter(mockParser)

	// Test formatting multiple logs
	sourceLogs := []string{
		"Log message 1",
		"Log message 2",
		"Log message 3",
	}
	expected := []string{
		"[FORMATTED] Log message 1",
		"[FORMATTED] Log message 2",
		"[FORMATTED] Log message 3",
	}
	result := formatter.FormatLogs(sourceLogs)

	// Check that the result has the expected length
	if len(result) != len(expected) {
		t.Errorf("FormatLogs() returned %d logs, want %d", len(result), len(expected))
		return
	}

	// Check each log
	for i, log := range result {
		if log != expected[i] {
			t.Errorf("FormatLogs()[%d] = %v, want %v", i, log, expected[i])
		}
	}
}

func TestTemplateFormatter_SetParser(t *testing.T) {
	// Create two mock parsers
	mockParser1 := &MockParser{
		parseFunc: func(sourceLog string) string {
			return "[PARSER1] " + sourceLog
		},
	}
	mockParser2 := &MockParser{
		parseFunc: func(sourceLog string) string {
			return "[PARSER2] " + sourceLog
		},
	}

	// Create a formatter with the first mock parser
	formatter := NewTemplateFormatter(mockParser1)

	// Test formatting with the first parser
	sourceLog := "This is a log message"
	expected1 := "[PARSER1] This is a log message"
	result1 := formatter.FormatLog(sourceLog)
	if result1 != expected1 {
		t.Errorf("FormatLog() with parser1 = %v, want %v", result1, expected1)
	}

	// Set the second parser
	formatter.SetParser(mockParser2)

	// Test formatting with the second parser
	expected2 := "[PARSER2] This is a log message"
	result2 := formatter.FormatLog(sourceLog)
	if result2 != expected2 {
		t.Errorf("FormatLog() with parser2 = %v, want %v", result2, expected2)
	}
}

// TestTemplateFormatter_WithRealParser_Simple tests the formatter with a real parser
func TestTemplateFormatter_WithRealParser_Simple(t *testing.T) {
	// Create a real parser
	p := parser.NewTemplateParser()
	err := p.SetTemplate("@timestamp-timestamp@ @level-string@ @message-string@", "[@timestamp@] @level@: @message@")
	if err != nil {
		t.Fatalf("Failed to set template: %v", err)
	}

	// Create a formatter with the real parser
	formatter := NewTemplateFormatter(p)

	// Test formatting a single log
	sourceLog := "2023-03-15T14:30:45Z INFO This is a sample log message"
	result := formatter.FormatLog(sourceLog)

	// Check if the result contains the expected parts
	if !strings.Contains(result, "2023-03-15T14:30:45Z") ||
		!strings.Contains(result, "INFO") ||
		!strings.Contains(result, "This is a sample log message") {
		t.Errorf("FormatLog() = %v, does not contain expected content", result)
	}
}
