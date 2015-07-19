package gowi

import (
	"fmt"
	"github.com/CaryLorrk/gosig"
)

type Button struct {
	WindowBase
	OnClicked func()
	Clicked   *gosig.Signal
}

func (w *Button) WndProc(msg uint32, wparam, lparam uintptr) uintptr {
	fmt.Println("Button.WndProc")
	w.Clicked.Emit()
	return 0
}

func NewButton(window MainWindow, id uintptr) *Button {
	c := new(Button)
	c.InitWindow(window, id)

	//c.Clicked, _ = gosig.NewSignal(func() {})
	c.Clicked, _ = gosig.NewSignal(func() {})
	//window.Connect(b1.Clicked, w.OnButton1Clicked)
	c.Clicked.Connect(func() {
		if c.OnClicked != nil {
			c.OnClicked()
		}
	})

	fmt.Printf("RegMsg h=%v\n", c.Handle())
	RegMsgHandler(c)

	return c
}
