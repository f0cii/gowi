package declarative

import (
	"github.com/nvsoft/gowi"
)

type Label struct {
	AssignTo **gowi.Label
	ID       uintptr
	Text     Property
}

func (l Label) Create(window *MainWindow) {
	win := *window.AssignTo
	c := gowi.NewLabel(*win, l.ID)
	if l.Text != nil {
		c.SetWindowText(l.Text.(string))
	}
	if l.AssignTo != nil {
		*l.AssignTo = c
	}
}
