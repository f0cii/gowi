package declarative

import (
	"github.com/nvsoft/gowi"
	"github.com/nvsoft/win"
)

type MainWindow struct {
	AssignTo **gowi.MainWindow
	ID       uintptr
	Icon     string
	Controls []Control
}

func (this *MainWindow) Create() {
	w := &gowi.MainWindow{}
	w.Create(this.ID)

	if this.AssignTo != nil {
		*this.AssignTo = w
	}

	for _, c := range this.Controls {
		c.Create(this)
	}
}

func (this *MainWindow) Show() {
	// Center in the owner window
	w := *this.AssignTo

	//fmt.Printf("aaaa\n")
	sWidth := win.GetSystemMetrics(win.SM_CXFULLSCREEN)
	sHeight := win.GetSystemMetrics(win.SM_CYFULLSCREEN)
	if sWidth != 0 && sHeight != 0 {
		rect := w.Bounds()
		rect.X = (int(sWidth) / 2) - (rect.Width / 2)
		rect.Y = (int(sHeight) / 2) - (rect.Height / 2)
		w.SetBounds(rect)
	}

	w.Show()
}
