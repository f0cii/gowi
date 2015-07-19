package declarative

import (
	"github.com/nvsoft/gowi"
)

type ListViewHeader struct {

}

type ListView struct {
	AssignTo **gowi.ListView
	ID       uintptr
}

func (l ListView) Create(window *MainWindow) {
	win := *window.AssignTo
	c := gowi.NewListView(*win, l.ID)
	if l.AssignTo != nil {
		*l.AssignTo = c
	}
}