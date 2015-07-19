package gowi

import (
	"fmt"
)

type Label struct {
	WindowBase
}

func (w *Label) WndProc(msg uint32, wparam, lparam uintptr) uintptr {
	fmt.Println("Label.WndProc")
	return 0
}

func NewLabel(window MainWindow, id uintptr) *Label {
	c := new(Label)
	c.InitWindow(window, id)

	fmt.Printf("RegMsg h=%v\n", c.Handle())
	RegMsgHandler(c)

	return c
}
