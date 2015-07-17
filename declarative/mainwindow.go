package declarative

import (
	"github.com/nvsoft/gowi"
)

type MainWindow struct {
	AssignTo            **gowi.MainWindow
	ID uintptr
	Icon string
	Controls []Widget
}

func (this *MainWindow) Create() {
	w := &gowi.MainWindow{}
	w.Create(this.ID)
	*this.AssignTo = w

	//for _, c := range this.Controls {

	//}
}