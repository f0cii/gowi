package declarative

import (
	"github.com/nvsoft/gowi"
)

type TabPage struct {
}

type TabControl struct {
	AssignTo **gowi.TabControl
	ID       uintptr
	Pages    []TabPage
}

func (l TabControl) Create(window *MainWindow) {
	win := *window.AssignTo
	c := gowi.NewTabControl(*win, l.ID)
	if l.AssignTo != nil {
		*l.AssignTo = c
	}
}
