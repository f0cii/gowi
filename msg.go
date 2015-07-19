package gowi

import (
	"github.com/nvsoft/win"
)

var gWindowRegistry map[win.HWND]Window

func RegMsgHandler(window Window) {
	gWindowRegistry[window.Handle()] = window
}

func UnRegMsgHandler(hwnd win.HWND) {
	delete(gWindowRegistry, hwnd)
}

func GetMsgHandler(hwnd win.HWND) Window {
	if widget, isExists := gWindowRegistry[hwnd]; isExists {
		return widget
	}

	return nil
}
