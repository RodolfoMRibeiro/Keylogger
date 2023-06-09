package util

import (
	"fmt"
	"os"
)

func CreateFile(fileName string) {
	filePath, err := os.Getwd()
	if err != nil {
		return
	}

	filePath = filePath + fileName

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Create the file
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		fmt.Println("File created:", filePath)
	} else {
		fmt.Println("File already exists:", filePath)
	}
}
