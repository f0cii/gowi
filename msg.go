package gowi

import (
	"github.com/nvsoft/win"
)

var gWidgetRegistry map[win.HWND]Widget

func RegMsgHandler(widget Widget) {
	gWidgetRegistry[widget.Handle()] = widget
}

func UnRegMsgHandler(hwnd win.HWND) {
	delete(gWidgetRegistry, hwnd)
}

func GetMsgHandler(hwnd win.HWND) Widget {
	if widget, isExists := gWidgetRegistry[hwnd]; isExists {
		return widget
	}

	return nil
}
