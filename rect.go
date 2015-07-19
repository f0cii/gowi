package gowi

import (
	"github.com/nvsoft/win"
)

type Rect struct {
	X, Y, Width, Height int
}

func rectFromRECT(r win.RECT) Rect {
	return Rect{
		X:      int(r.Left),
		Y:      int(r.Top),
		Width:  int(r.Right - r.Left),
		Height: int(r.Bottom - r.Top),
	}
}

func (r Rect) ToRECT() win.RECT {
	return win.RECT{
		int32(r.X),
		int32(r.Y),
		int32(r.X + r.Width),
		int32(r.Y + r.Height),
	}
}
