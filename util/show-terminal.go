package util

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	user32     = syscall.NewLazyDLL("user32.dll")
	showWindow = user32.NewProc("ShowWindow")
	findWindow = user32.NewProc("FindWindowA")
)

const (
	consoleWindowClass = "ConsoleWindowClass"
	hide               = 0
	show               = 5
)

// SetConsoleVisibility sets the visibility of the console window.
func SetConsoleVisibility(isVisible bool) error {
	windowClass, err := syscall.UTF16PtrFromString(consoleWindowClass)
	if err != nil {
		return fmt.Errorf("failed to get window class: %w", err)
	}

	var mode int
	if isVisible {
		mode = show
	} else {
		mode = hide
	}

	handle, _, err := findWindow.Call(uintptr(unsafe.Pointer(windowClass)), 0)
	if handle == 0 {
		return fmt.Errorf("failed to find console window: %w", err)
	}

	_, _, err = showWindow.Call(handle, uintptr(mode))
	if err != nil {
		return fmt.Errorf("failed to set console visibility: %w", err)
	}

	return nil
}
