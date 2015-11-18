package gowi

import (
	"github.com/lroc/win"
)

var hInst win.HINSTANCE

// Get application instance handle
func GetAppInstance() win.HINSTANCE {
	return hInst
}

type Application struct {
}

// Run the main message loop
func (e *Application) Run() int {
	var msg win.MSG

	for {
		switch win.GetMessage(&msg, 0, 0, 0) {
		case 0:
			return int(msg.WParam)
		case -1:
			return -1
		}

		win.TranslateMessage(&msg)
		win.DispatchMessage(&msg)
	}

	return 0
}

func New() *Application {
	app := new(Application)
	return app
}
