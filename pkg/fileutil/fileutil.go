package fileutil

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// FileReader defines the interface for reading files
type FileReader interface {
	// ReadLines reads all lines from a file
	ReadLines(filepath string) ([]string, error)
	// FileExists checks if a file exists
	FileExists(filepath string) bool
}

// FileWriter defines the interface for writing files
type FileWriter interface {
	// WriteLines writes lines to a file
	WriteLines(lines []string, filepath string) error
}

// FileUtil implements both FileReader and FileWriter interfaces
type FileUtil struct{}

// NewFileUtil creates a new FileUtil
func NewFileUtil() *FileUtil {
	return &FileUtil{}
}

// ReadLines reads all lines from a file
func (f *FileUtil) ReadLines(filepath string) ([]string, error) {
	if !f.FileExists(filepath) {
		return nil, fmt.Errorf("file does not exist: %s", filepath)
	}

	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines, nil
}

// FileExists checks if a file exists and is not a directory
func (f *FileUtil) FileExists(filepath string) bool {
	info, err := os.Stat(filepath)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// WriteLines writes lines to a file
func (f *FileUtil) WriteLines(lines []string, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	for _, line := range lines {
		entry := line + strings.Repeat("-", 80) + "\n"
		_, err := file.WriteString(entry)
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}
	return nil
}
