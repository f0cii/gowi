package declarative

import (
	"github.com/nvsoft/gowi"
)

type Button struct {
	AssignTo            **gowi.Button
	ID 					uintptr
	Name                string
	//Enabled             Property
	//Visible             Property
	//Font                Font
	//ToolTipText         Property
	//MinSize             Size
	//MaxSize             Size
	//StretchFactor       int
	//Row                 int
	//RowSpan             int
	//Column              int
	//ColumnSpan          int
	//AlwaysConsumeSpace  bool
	//ContextMenuItems    []MenuItem
	//OnKeyDown           walk.KeyEventHandler
	//OnKeyPress          walk.KeyEventHandler
	//OnKeyUp             walk.KeyEventHandler
	//OnMouseDown         walk.MouseEventHandler
	//OnMouseMove         walk.MouseEventHandler
	//OnMouseUp           walk.MouseEventHandler
	//OnSizeChanged       walk.EventHandler
	//Text                Property
	//Checked             Property
	//CheckState          Property
	//Tristate            bool
	OnClicked           gowi.EventHandler
	//OnCheckedChanged    walk.EventHandler
	//OnCheckStateChanged walk.EventHandler
}