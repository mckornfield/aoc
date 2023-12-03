package resources

import (
	"os"
	"strings"
)

func ReadFile(file_name string) string {
	fileContent, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}

	// Convert []byte to string
	return string(fileContent)
}

func ReadLines(file_name string) []string {
	return strings.Split(ReadFile(file_name), "\n")
}
