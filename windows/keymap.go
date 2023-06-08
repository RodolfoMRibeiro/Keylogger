package windows

const (
	VK_BACK       = 0x08
	VK_RETURN     = 0x0D
	VK_SPACE      = 0x20
	VK_TAB        = 0x09
	VK_SHIFT      = 0x10
	VK_LSHIFT     = 0xA0
	VK_RSHIFT     = 0xA1
	VK_CONTROL    = 0x11
	VK_LCONTROL   = 0xA2
	VK_RCONTROL   = 0xA3
	VK_MENU       = 0x12
	VK_LWIN       = 0x5B
	VK_RWIN       = 0x5C
	VK_ESCAPE     = 0x1B
	VK_END        = 0x23
	VK_HOME       = 0x24
	VK_LEFT       = 0x25
	VK_RIGHT      = 0x27
	VK_UP         = 0x26
	VK_DOWN       = 0x28
	VK_PRIOR      = 0x21
	VK_NEXT       = 0x22
	VK_OEM_PERIOD = 0xBE
	VK_DECIMAL    = 0x6E
	VK_OEM_PLUS   = 0xBB
	VK_OEM_MINUS  = 0xBD
	VK_ADD        = 0x6B
	VK_SUBTRACT   = 0x6D
	VK_CAPITAL    = 0x14
)

var keyName = map[int]string{
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
