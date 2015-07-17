package gowi

import (
	"fmt"
	//"github.com/CaryLorrk/gosig"
)

type Label struct {
	WidgetBase
	//Clicked *gosig.Signal
}

func (w *Label) WndProc(msg uint32, wparam, lparam uintptr) uintptr {
	fmt.Println("Label.WndProc")
	//w.Clicked.Emit()
	return 0
}

func NewLabel(window Window, id uintptr) *Label {
	c := new(Label)
	c.InitWidget(window, id)

	//c.Clicked, _ = gosig.NewSignal(func() {})

	fmt.Printf("RegMsg h=%v\n", c.Handle())
	RegMsgHandler(c)

	return c
}
