package windows

const (
	VK_BACK       uintptr = 0x08
	VK_RETURN     uintptr = 0x0D
	VK_SPACE      uintptr = 0x20
	VK_TAB        uintptr = 0x09
	VK_SHIFT      uintptr = 0x10
	VK_LSHIFT     uintptr = 0xA0
	VK_RSHIFT     uintptr = 0xA1
	VK_CONTROL    uintptr = 0x11
	VK_LCONTROL   uintptr = 0xA2
	VK_RCONTROL   uintptr = 0xA3
	VK_MENU       uintptr = 0x12
	VK_LWIN       uintptr = 0x5B
	VK_RWIN       uintptr = 0x5C
	VK_ESCAPE     uintptr = 0x1B
	VK_END        uintptr = 0x23
	VK_HOME       uintptr = 0x24
	VK_LEFT       uintptr = 0x25
	VK_RIGHT      uintptr = 0x27
	VK_UP         uintptr = 0x26
	VK_DOWN       uintptr = 0x28
	VK_PRIOR      uintptr = 0x21
	VK_NEXT       uintptr = 0x22
	VK_OEM_PERIOD uintptr = 0xBE
	VK_DECIMAL    uintptr = 0x6E
	VK_OEM_PLUS   uintptr = 0xBB
	VK_OEM_MINUS  uintptr = 0xBD
	VK_ADD        uintptr = 0x6B
	VK_SUBTRACT   uintptr = 0x6D
	VK_CAPITAL    uintptr = 0x14
)

var keyName = map[uintptr]string{
	VK_BACK:       "[BACKSPACE]",
	VK_RETURN:     "\n",
	VK_SPACE:      "_",
	VK_TAB:        "[TAB]",
	VK_SHIFT:      "[SHIFT]",
	VK_LSHIFT:     "[LSHIFT]",
	VK_RSHIFT:     "[RSHIFT]",
	VK_CONTROL:    "[CONTROL]",
	VK_LCONTROL:   "[LCONTROL]",
	VK_RCONTROL:   "[RCONTROL]",
	VK_MENU:       "[ALT]",
	VK_LWIN:       "[LWIN]",
	VK_RWIN:       "[RWIN]",
	VK_ESCAPE:     "[ESCAPE]",
	VK_END:        "[END]",
	VK_HOME:       "[HOME]",
	VK_LEFT:       "[LEFT]",
	VK_RIGHT:      "[RIGHT]",
	VK_UP:         "[UP]",
	VK_DOWN:       "[DOWN]",
	VK_PRIOR:      "[PG_UP]",
	VK_NEXT:       "[PG_DOWN]",
	VK_OEM_PERIOD: ".",
	VK_DECIMAL:    ".",
	VK_OEM_PLUS:   "+",
	VK_OEM_MINUS:  "-",
	VK_ADD:        "+",
	VK_SUBTRACT:   "-",
	VK_CAPITAL:    "[CAPSLOCK]",
}
