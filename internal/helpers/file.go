package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures that filepath points to a valid file
func IsFile(filePath string) error {
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

// Writes a LogEntry to a file
func WriteArrayToFile(logEntries []string, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

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

// Reads file and adds lines to array
func ReadFileToArray(filepath string) []string {
	err := IsFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
	}

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err = scanner.Err(); err != nil {
		fmt.Println(`error reading file`)
	}

	arr := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		arr = append(arr, line)
	}

	return arr
}
