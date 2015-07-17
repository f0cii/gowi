package main

import (
	"github.com/nvsoft/gowi"
	. "github.com/nvsoft/gowi/declarative"
	"fmt"
)

type MyMainWindow struct {
	*gowi.MainWindow
	label1 *gowi.Label
	button1 *gowi.Button
}

func main() {
	engine := gowi.New()

	mw := new(MyMainWindow)

	m := MainWindow{
		AssignTo:&mw.MainWindow,
		ID:1000,
		Controls:[]Widget{
			Button{
				AssignTo:&mw.button1,
				ID:1001,
				OnClicked:func() {
					fmt.Println("Button clicked.")
					mw.SetWindowText("Clicked")
				},
			},
		},
	}
	m.Create()

	mw.SetWindowText("aaaaa")
	engine.Run()
}