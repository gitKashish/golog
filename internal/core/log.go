package core

import (
	"fmt"
	"os"
)

// Parses a single log line
func ParseLogLine(line string) string {
	template, err := GetTemplate()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	template.Parse(line)

	return template.Execute()
}
