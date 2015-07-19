package declarative

import (
	"github.com/nvsoft/gowi"
)

type LineEdit struct {
	AssignTo **gowi.LineEdit
	ID       uintptr
	Text     Property
}

func (l LineEdit) Create(window *MainWindow) {
	win := *window.AssignTo
	c := gowi.NewLineEdit(*win, l.ID)
	if l.Text != nil {
		c.SetWindowText(l.Text.(string))
	}
	if l.AssignTo != nil {
		*l.AssignTo = c
	}
}
