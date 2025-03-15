package fileutil

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFileUtil_FileExists(t *testing.T) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "fileutil_test_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "fileutil_test_dir")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	fileUtil := NewFileUtil()

	tests := []struct {
		name     string
		filepath string
		want     bool
	}{
		{
			name:     "Existing file",
			filepath: tmpFile.Name(),
			want:     true,
		},
		{
			name:     "Non-existing file",
			filepath: filepath.Join(os.TempDir(), "non_existing_file.txt"),
			want:     false,
		},
		{
			name:     "Directory",
			filepath: tmpDir,
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileUtil.FileExists(tt.filepath); got != tt.want {
				t.Errorf("FileUtil.FileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileUtil_ReadLines(t *testing.T) {
	// Create a temporary file with some content
	tmpFile, err := os.CreateTemp("", "fileutil_test_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Write some lines to the file
	content := "Line 1\nLine 2\nLine 3\n"
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	fileUtil := NewFileUtil()

	// Test reading lines from the file
	lines, err := fileUtil.ReadLines(tmpFile.Name())
	if err != nil {
		t.Fatalf("FileUtil.ReadLines() error = %v", err)
	}

	// Check the number of lines
	expectedLines := []string{"Line 1", "Line 2", "Line 3"}
	if len(lines) != len(expectedLines) {
		t.Errorf("FileUtil.ReadLines() returned %d lines, want %d", len(lines), len(expectedLines))
	}

	// Check each line
	for i, line := range lines {
		if line != expectedLines[i] {
			t.Errorf("FileUtil.ReadLines()[%d] = %v, want %v", i, line, expectedLines[i])
		}
	}

	// Test reading from a non-existing file
	_, err = fileUtil.ReadLines(filepath.Join(os.TempDir(), "non_existing_file.txt"))
	if err == nil {
		t.Errorf("FileUtil.ReadLines() error = nil, want error for non-existing file")
	}
}

func TestFileUtil_WriteLines(t *testing.T) {
	// Create a temporary file path
	tmpFile, err := os.CreateTemp("", "fileutil_test_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	tmpFilePath := tmpFile.Name()
	tmpFile.Close() // Close the file so we can write to it with WriteLines
	defer os.Remove(tmpFilePath)

	fileUtil := NewFileUtil()

	// Test writing lines to the file
	lines := []string{"Line 1", "Line 2", "Line 3"}
	err = fileUtil.WriteLines(lines, tmpFilePath)
	if err != nil {
		t.Fatalf("FileUtil.WriteLines() error = %v", err)
	}

	// Read the file back to verify the content
	content, err := os.ReadFile(tmpFilePath)
	if err != nil {
		t.Fatalf("Failed to read temp file: %v", err)
	}

	// Check the content
	expectedContent := "Line 1" + strings.Repeat("-", 80) + "\n" +
		"Line 2" + strings.Repeat("-", 80) + "\n" +
		"Line 3" + strings.Repeat("-", 80) + "\n"
	if string(content) != expectedContent {
		t.Errorf("FileUtil.WriteLines() wrote %v, want %v", string(content), expectedContent)
	}
}
