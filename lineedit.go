package gowi

import (
	"fmt"
)

type LineEdit struct {
	WindowBase
}

func (e *LineEdit) WndProc(msg uint32, wparam, lparam uintptr) uintptr {
	fmt.Println("LineEdit.WndProc")
	return 0
}

func NewLineEdit(window MainWindow, id uintptr) *LineEdit {
	c := new(LineEdit)
	c.InitWindow(window, id)

	fmt.Printf("RegMsg h=%v\n", c.Handle())
	RegMsgHandler(c)

	return c
}
