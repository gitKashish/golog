package core

import (
	"fmt"
	"os"
)

// Parses a single log line into a LogEntry
func ParseLogLine(line string) string {
	template, err := GetTemplate()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = template.Parse(line)
	if err != nil {
		fmt.Println(err.Error())
	}

	return template.Execute()
}
