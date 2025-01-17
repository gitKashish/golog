package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// LogEntry represents a structured log entry
type LogEntry struct {
	Timestamp string                 `json:"timestamp,omitempty"`
	PmId      string                 `json:"pmid,omitempty`
	Server    string                 `json:"server,omitempty"`
	Module    string                 `json:"module,omitempty"`
	Api       string                 `json:"api,omitempty"`
	Event     string                 `json:"event,omitempty"`
	Details   map[string]interface{} `json:"details,omitempty"`
	Raw       string                 `json:"-"`
}

// parseLogLine parses a single log line into a LogEntry
func parseLogLine(line string) *LogEntry {
	timestampPattern := regexp.MustCompile(`-->(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3})`)
	pmIdPattern := regexp.MustCompile(`(^\d+)\|[^\|]+?\s\|`)
	serverPattern := regexp.MustCompile(`\|([^\|]+?)\s\|`)
	modulePattern := regexp.MustCompile(`:----:\s(.*)\s:=:`)
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
			if event, exists := details["EVENT"]; exists {
				entry.Event = fmt.Sprintf("%v", event)
			}
		}
	}

	return entry
}

// formatLogEntry formats a LogEntry into a readable string
func formatLogEntry(entry *LogEntry) string {
	if entry.Details != nil {
		prettyDetails, _ := json.MarshalIndent(entry.Details, "", "  ")
		return fmt.Sprintf("Timestamp: %s\nPM ID: %s\nServer: %s\nModule: %s\nAPI: %s\nEvent: %s\nDetails:\n%s\n",
			entry.Timestamp, entry.PmId, entry.Server, entry.Module, entry.Api, entry.Event, string(prettyDetails))
	}
	return entry.Raw
}

func main() {

	// Getting file path from args
	filePath := os.Args[1]

	info, err := os.Stat(filePath)

	// Ensuring file exists
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		}
		fmt.Println("Could not get file stats", err)
		return
	}

	// Ensuring it is not a directory
	if info.IsDir() {
		fmt.Println("Given path leads to a directory")
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		entry := parseLogLine(line)
		fmt.Println(formatLogEntry(entry))
		fmt.Println(strings.Repeat("-", 80))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
