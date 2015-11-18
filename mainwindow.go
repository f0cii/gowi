package gowi

import (
	"fmt"
	"github.com/CaryLorrk/gosig"
	"github.com/lroc/win"
	"syscall"
)

var defaultDialogProcPtr = syscall.NewCallback(defaultDialogProc)

// Dialog class
type MainWindow struct {
	WindowBase
}

func (w *MainWindow) AsWindowBase() *MainWindow {
	return w
}

func (w *MainWindow) Create(idd uintptr) {
	if idd == 0 {
		panic("Create MainWindow FAIL.")
	}
	w.idd = idd

	w.hwnd = win.CreateDialogParam(hInst, win.MAKEINTRESOURCE(w.idd), 0, defaultDialogProcPtr, 0)
	fmt.Printf("Create window hwnd=%v\n", w.hwnd)
	if w.hwnd == win.HWND(0) {
		fmt.Printf("Create window fail.[%v]\n", idd)
	}
}

func (w *MainWindow) Connect(sig *gosig.Signal, slot interface{}) {
	sig.Connect(slot)
}

func defaultDialogProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_COMMAND:
		fmt.Printf("WM_COMMAND msg=%v\n", msg)
		if lParam != 0 { //Reflect message to control
			h := win.HWND(lParam)
			fmt.Printf("WM_COMMAND h=%v\n", h)
			if handler := GetMsgHandler(h); handler != nil {
				fmt.Println("WM_COMMAND handler.WndProc")
				ret := handler.WndProc(msg, wParam, lParam)
				if ret != 0 {
					//win.SetWindowLong(hwnd, win.DWL_MSGRESULT, int32(ret))
					fmt.Println("WM_COMMAND TRUE")
					return win.TRUE
				}
			}
		}
		fmt.Println("WM_COMMAND DONE")
		return 0
	case win.WM_CLOSE:
		win.DestroyWindow(hwnd)
		return 0
	case win.WM_DESTROY:
		win.PostQuitMessage(0)
		return 1
	}

	return 0
}
