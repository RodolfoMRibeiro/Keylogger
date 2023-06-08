package util

import (
	"syscall"
	"unsafe"
)

var (
	user_32    = syscall.NewLazyDLL("user32.dll")
	showWindow = user_32.NewProc("ShowWindow")
	findWindow = user_32.NewProc("FindWindowA")
)

const (
	CONSOLE_WINDOW_CLASS = "ConsoleWindowClass"
	HIDE                 = 0
	SHOW                 = 5
)

// Set the visibility of the console window.
func Stealth(isVisible bool) {
	windowClass, _ := syscall.UTF16PtrFromString(CONSOLE_WINDOW_CLASS)
	var mode int

	if isVisible {
		mode = SHOW
	} else {
		mode = HIDE
	}

	handle, _, _ := findWindow.Call(uintptr(unsafe.Pointer(windowClass)), 0)

	showWindow.Call(handle, uintptr(mode))
}
