package gowi

type TabControl struct {
	WindowBase
}

func (w *TabControl) WndProc(msg uint32, wparam, lparam uintptr) uintptr {
	return 0
}

func NewTabControl(window MainWindow, id uintptr) *TabControl {
	c := new(TabControl)
	c.InitWindow(window, id)

	RegMsgHandler(c)

	return c
}
