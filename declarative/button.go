package declarative

import (
	"github.com/nvsoft/gowi"
)

type Button struct {
	AssignTo  **gowi.Button
	ID        uintptr
	Text      Property
	Enabled   Property
	Visible   Property
	OnClicked gowi.EventHandler
}

func (b Button) Create(window *MainWindow) {
	win := *window.AssignTo
	c := gowi.NewButton(*win, b.ID)
	if b.Text != nil {
		c.SetWindowText(b.Text.(string))
	}
	if b.Enabled != nil {
		c.SetEnabled(b.Enabled.(bool))
	}
	if b.Visible != nil {
		c.SetVisible(b.Visible.(bool))
	}
	c.OnClicked = b.OnClicked
	if b.AssignTo != nil {
		*b.AssignTo = c
	}
}
