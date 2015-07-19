package gowi

type ListView struct {
	WindowBase
}

func (e *ListView) WndProc(msg uint32, wparam, lparam uintptr) uintptr {
	return 0
}

func NewListView(window MainWindow, id uintptr) *ListView {
	c := new(ListView)
	c.InitWindow(window, id)

	RegMsgHandler(c)

	return c
}