package twenty_twenty_three

import (
	"os"
)

func ReadFile(file_name string) string {
	fileContent, err := os.ReadFile("resources" + string(os.PathSeparator) + file_name)
	if err != nil {
		panic(err)
	}

	// Convert []byte to string
	return string(fileContent)
}
