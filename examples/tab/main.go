package main

import (
	"github.com/nvsoft/gowi"
	. "github.com/nvsoft/gowi/declarative"
)

type MyMainWindow struct {
	*gowi.MainWindow
	tab1 *gowi.TabControl
}

func main() {
	app := gowi.New()

	mw := &MyMainWindow{}

	m := MainWindow{
		AssignTo: &mw.MainWindow,
		ID:       108,
		Controls: []Control{
			TabControl{
				AssignTo: &mw.tab1,
				ID:       40000,
			},
		},
	}

	m.Create()
	m.Show()

	app.Run()
}
