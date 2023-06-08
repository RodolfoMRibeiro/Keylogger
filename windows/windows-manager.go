package windows

import (
	"syscall"
	"unsafe"
)

func NewWindowManager() *WindowsManager {
	user32DLL := syscall.NewLazyDLL("user32.dll")

	return &WindowsManager{
		getForegroundWindow:      user32DLL.NewProc("GetForegroundWindow"),
		getWindowThreadProcessId: user32DLL.NewProc("GetWindowThreadProcessId"),
		getKeyboardLayout:        user32DLL.NewProc("GetKeyboardLayout"),
		getWindowText:            user32DLL.NewProc("GetWindowTextW"),
		mapVirtualKeyEx:          user32DLL.NewProc("MapVirtualKeyExW"),
		getKeyState:              user32DLL.NewProc("GetKeyState"),
		setWindowsHook:           user32DLL.NewProc("SetWindowsHookExW"),
		callNextHook:             user32DLL.NewProc("CallNextHookEx"),
		getMessage:               user32DLL.NewProc("GetMessageW"),
		unhookWindowsHookEx:      user32DLL.NewProc("UnhookWindowsHookEx"),
	}
}

type WindowsManager struct {
	getForegroundWindow      *syscall.LazyProc
	getWindowThreadProcessId *syscall.LazyProc
	getKeyboardLayout        *syscall.LazyProc
	getWindowText            *syscall.LazyProc
	mapVirtualKeyEx          *syscall.LazyProc
	getKeyState              *syscall.LazyProc
	setWindowsHook           *syscall.LazyProc
	callNextHook             *syscall.LazyProc
	getMessage               *syscall.LazyProc
	unhookWindowsHookEx      *syscall.LazyProc
}

func (w *WindowsManager) GetForegroundWindow() uintptr {
	foreground, _, _ := w.getForegroundWindow.Call()
	return foreground
}

func (w *WindowsManager) GetWindowThreadProcessId(hwnd uintptr) (uint32, error) {
	var threadID uint32
	w.getWindowThreadProcessId.Call(hwnd, uintptr(unsafe.Pointer(&threadID)))
	return threadID, nil
}

func (w *WindowsManager) GetKeyboardLayout(threadID uint32) (uintptr, error) {
	layout, _, _ := w.getKeyboardLayout.Call(uintptr(threadID))
	return layout, nil
}

func (w *WindowsManager) GetWindowText(hwnd uintptr) (string, error) {
	windowTitle := make([]uint16, 256)
	w.getWindowText.Call(hwnd, uintptr(unsafe.Pointer(&windowTitle[0])), 256)
	windowTitleStr := syscall.UTF16ToString(windowTitle)
	return windowTitleStr, nil
}

func (w *WindowsManager) MapVirtualKeyEx(keyCode int, flags uintptr, layout uintptr) (uintptr, error) {
	virtualKey, _, _ := w.mapVirtualKeyEx.Call(uintptr(keyCode), flags, layout)
	return virtualKey, nil
}

func (w *WindowsManager) GetKeyState(keyCode int) (int16, error) {
	r, _, _ := w.getKeyState.Call(uintptr(keyCode))
	return int16(r), nil
}
