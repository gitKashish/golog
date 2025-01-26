package core

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

// Represents a structured log entry
type LogEntry struct {
	Timestamp string                 `json:"timestamp,omitempty"`
	PmId      string                 `json:"pmid,omitempty"`
	Server    string                 `json:"server,omitempty"`
	Module    string                 `json:"module,omitempty"`
	Api       string                 `json:"api,omitempty"`
	Details   map[string]interface{} `json:"details,omitempty"`
	Raw       string                 `json:"-"`
}

// Parses a single log line into a LogEntry
func ParseLogLine(line string) *LogEntry {
	if strings.Trim(line, " ") == "" {
		return nil
	}

	timestampPattern := regexp.MustCompile(`-->(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3})`)
	pmIdPattern := regexp.MustCompile(`(^\d+)\|[^\|]+?\s\|`)
	serverPattern := regexp.MustCompile(`\|([^\|]+?)\s\|`)
	modulePattern := regexp.MustCompile(`:----:\s(.*)\s:=:\s.*\s:=:`)
	apiPattern := regexp.MustCompile(`:=:\s(.*)\s:=:`)
	detailsPattern := regexp.MustCompile(`:----:\s.*\s:=:\s.*\s:=:\s(.+)`)

	entry := &LogEntry{Raw: line}

	// Extract timestamp
	if matches := timestampPattern.FindStringSubmatch(line); len(matches) > 1 {
		entry.Timestamp = matches[1]
	}

	// Extract pmId
	if matches := pmIdPattern.FindStringSubmatch(line); len(matches) > 1 {
		entry.PmId = matches[1]
	}

	// Extract server
	if matches := serverPattern.FindStringSubmatch(line); len(matches) > 1 {
		entry.Server = matches[1]
	}

	// Extract module
	if matches := modulePattern.FindStringSubmatch(line); len(matches) > 1 {
		entry.Module = matches[1]
	}

	// Extract api
	if matches := apiPattern.FindStringSubmatch(line); len(matches) > 1 {
		entry.Api = matches[1]
	}

	// Extract details
	if matches := detailsPattern.FindStringSubmatch(line); len(matches) > 1 {
		var details map[string]interface{}
		if err := json.Unmarshal([]byte(matches[1]), &details); err == nil {
			entry.Details = details
		}
	}

	return entry
}

// Formats a LogEntry into a readable string
func FormatLogEntry(entry *LogEntry) string {
	if entry.Details != nil {
		prettyDetails, _ := json.MarshalIndent(entry.Details, "", "  ")
		return fmt.Sprintf("Timestamp: %s\nPM ID: %s\nServer: %s\nModule: %s\nAPI: %s\nDetails:\n%s\n",
			entry.Timestamp,
			entry.PmId,
			entry.Server,
			entry.Module,
			entry.Api,
			string(prettyDetails))
	}
	return entry.Raw + "\n"
}
