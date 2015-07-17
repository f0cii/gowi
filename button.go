package gowi

import (
	"fmt"
	"github.com/CaryLorrk/gosig"
)

type Button struct {
	WidgetBase
	Clicked *gosig.Signal
}

func (w *Button) WndProc(msg uint32, wparam, lparam uintptr) uintptr {
	fmt.Println("Button.WndProc")
	w.Clicked.Emit()
	return 0
}

func NewButton(window Window, id uintptr) *Button {
	c := new(Button)
	c.InitWidget(window, id)

	c.Clicked, _ = gosig.NewSignal(func() {})

	fmt.Printf("RegMsg h=%v\n", c.Handle())
	RegMsgHandler(c)

	return c
}
