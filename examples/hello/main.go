package main

import (
	"fmt"
	"github.com/nvsoft/gowi"
	. "github.com/nvsoft/gowi/declarative"
)

type MyMainWindow struct {
	*gowi.MainWindow
	label1    *gowi.Label
	lineedit1 *gowi.LineEdit
	button1   *gowi.Button
}

func main() {
	app := gowi.New()

	mw := new(MyMainWindow)

	m := MainWindow{
		AssignTo: &mw.MainWindow,
		ID:       100,
		Controls: []Control{
			Label{
				AssignTo: &mw.label1,
				ID:       1002,
				Text:     "????",
			},
			LineEdit{
				AssignTo: &mw.lineedit1,
				ID:       40004,
				Text:     "abc",
			},
			Button{
				AssignTo: &mw.button1,
				ID:       1001,
				Text:     "Click Me!",
				OnClicked: func() {
					fmt.Println("Button clicked.")
					//ll := mw.lineedit1.GetWindowText()
					//fmt.Printf("ll=%v\n", ll)
				},
			},
		},
	}
	m.Create()
	m.Show()

	mw.SetWindowText("Hello!")
	app.Run()
}
