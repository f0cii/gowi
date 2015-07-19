package gowi

import (
	"github.com/nvsoft/win"
)

func init() {
	hInst := win.GetModuleHandle(nil)
	if hInst == 0 {
		panic("GetModuleHandle")
	}
	gWindowRegistry = make(map[win.HWND]Window)
}
