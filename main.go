package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
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
func parseLogLine(line string) *LogEntry {
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
func formatLogEntry(entry *LogEntry) string {
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

// Writes a LogEntry to a file
func writeLogEntriesToFile(logEntries []string, file *os.File) error {
	for _, logEntry := range logEntries {
		entry := logEntry + strings.Repeat("-", 80) + "\n"
		_, err := file.WriteString(entry)
		if err != nil {
			err = fmt.Errorf("error writing to file %s", file.Name())
			return err
		}
	}
	return nil
}

// Checks if a file is directory
func isFile(filePath string) error {
	info, err := os.Stat(filePath)
	if err != nil {
		if err == os.ErrNotExist {
			err = fmt.Errorf("file does not exist: %s", filePath)
			return err
		}
		err = fmt.Errorf("cannot check file stats: %s", err.Error())
		return err
	}

	if info.IsDir() {
		err = fmt.Errorf("file %s is a directory", filePath)
		return err
	}

	return nil
}

// Usage text
func usage() string {
	usageElements := []string{
		"Usage: golog -source=<source-file-path> [-target=<target-file-path>] [-show]",
		"Options:",
		"\tsource (required): File path of source file with unformatted logs.",
		"\ttarget: File path where you want to store the formatted logs. File is created if it does not exist.",
		"\tshow: Set flag to print formatted logs on console even if -target flag is set",
	}
	return strings.Join(usageElements, "\n")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Too few arguments!\n%s", usage())
		return
	}

	// flag for setting an input file path to read unformatted logs from
	sourceFileUsageText := "File path of source file with unformatted logs"
	sourceFilePath := flag.String("source", "", sourceFileUsageText)

	// flag for setting an output file path to store the formatted logs into
	outFileUsageText := "File path where you want to store the formatted logs. File is created if it does not exist"
	outFilePath := flag.String("target", "", outFileUsageText)

	// flag to show logs even if target file is given
	showLogsUsageText := "Set flag to print formatted logs on console even if -target flag is set"
	showLogs := flag.Bool("show", false, showLogsUsageText)

	flag.Parse()
	if *sourceFilePath == "" {
		fmt.Printf("No source file specified")
		return
	}

	err := isFile(*sourceFilePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Open source file
	sourceFile, err := os.Open(*sourceFilePath)
	if err != nil {
		fmt.Printf("error opening file %s", *sourceFilePath)
		return
	}
	defer sourceFile.Close()

	scanner := bufio.NewScanner(sourceFile)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Scanning source log lines
	entries := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		entry := parseLogLine(line)
		if entry != nil {
			entries = append(entries, formatLogEntry(entry))
		}
	}

	if *outFilePath != "" {
		// Writing formatted logs to file
		outFile, err := os.Create(*outFilePath)
		if err != nil {
			fmt.Printf("error opening/creating output file %s", *outFilePath)
			return
		}
		defer outFile.Close()
		err = writeLogEntriesToFile(entries, outFile)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("formatted logs written to file %s", *outFilePath)
	}

	if *outFilePath == "" || *showLogs {
		// Writing logs to console
		for _, logEntry := range entries {
			entry := logEntry + strings.Repeat("-", 80) + "\n"
			fmt.Println(entry)
		}
	}
}
