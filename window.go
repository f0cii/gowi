package gowi

import (
	"fmt"
	"github.com/nvsoft/win"
	"syscall"
	"unsafe"
)

// Window
type Window interface {
	AsWindowBase() *WindowBase

	InitWindow(window MainWindow, id uintptr)

	Handle() win.HWND

	GetWindowText() string

	SetWindowText(title string)

	SetIcon(iconType int, icon *Icon)

	Show()

	Hide()

	ShowMinimized()

	ShowMaximized()

	ShowFullScreen()

	ShowNormal()

	IsEnabled() bool

	IsVisible() bool

	SetVisible(value bool)

	SetEnabled(b bool)

	SetDisabled(disable bool)

	SetFocus()

	Bounds() Rect

	SetBounds(value Rect)

	Close()

	WndProc(msg uint32, wparam, lparam uintptr) uintptr
}

// WindowBase
type WindowBase struct {
	hwnd   win.HWND
	parent win.HWND
	idd    uintptr
}

func (w *WindowBase) AsWindowBase() *WindowBase {
	return w
}

// 设置窗口对应资源Idd
func (w *WindowBase) InitWindow(window MainWindow, id uintptr) {
	w.idd = id
	h := window.AsWindowBase().Handle()
	fmt.Printf("h=%v\n", h)
	w.hwnd = win.GetDlgItem(h, id)
	fmt.Printf("w.hwnd=%v\n", w.hwnd)
}

func (w *WindowBase) Handle() win.HWND {
	return w.hwnd
}

//func (w *WidgetBase) GetWindowText() string {
//	return win.GetWindowText(w.hwnd)
//}

func (w *WindowBase) SetWindowText(title string) {
	fmt.Printf("SetCaption hwnd: %v, %s\n", w.hwnd, title)
	win.SetWindowText(w.hwnd, title)
}

func (w *WindowBase) GetWindowText() string {
	textLength := win.SendMessage(w.hwnd, win.WM_GETTEXTLENGTH, 0, 0)
	buf := make([]uint16, textLength+1)
	win.SendMessage(w.hwnd, win.WM_GETTEXT, uintptr(textLength+1), uintptr(unsafe.Pointer(&buf[0])))
	return syscall.UTF16ToString(buf)
}

//func (w *WidgetBase) SetWindowText(text string) {
//	if win.TRUE != win.SendMessage(w.hwnd, win.WM_SETTEXT, 0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text)))) {
//return errors.New("WM_SETTEXT failed")
//	}

//return nil
//}

// IconType: 1 - ICON_BIG; 0 - ICON_SMALL
func (w *WindowBase) SetIcon(iconType int, icon *Icon) {
	if iconType > 1 {
		panic("IconType is invalid")
	}

	win.SendMessage(w.hwnd, win.WM_SETICON, uintptr(iconType), uintptr(icon.Handle()))
}

func (w *WindowBase) Show() {
	win.ShowWindow(w.hwnd, win.SW_SHOW)
}

func (w *WindowBase) Hide() {
	win.ShowWindow(w.hwnd, win.SW_HIDE)
}

func (w *WindowBase) ShowMinimized() {
	win.ShowWindow(w.hwnd, win.SW_MINIMIZE)
}

func (w *WindowBase) ShowMaximized() {
	win.ShowWindow(w.hwnd, win.SW_MAXIMIZE)
}

func (w *WindowBase) ShowFullScreen() {
	win.ShowWindow(w.hwnd, win.SW_SHOWMAXIMIZED)
}

func (w *WindowBase) ShowNormal() {
	win.ShowWindow(w.hwnd, win.SW_SHOWNORMAL)
}

func (w *WindowBase) IsEnabled() bool {
	return win.IsWindowEnabled(w.hwnd)
}

func (w *WindowBase) IsVisible() bool {
	return win.IsWindowVisible(w.hwnd)
}

func (w *WindowBase) SetVisible(value bool) {
	var cmd int32
	if value {
		cmd = win.SW_SHOW
	} else {
		cmd = win.SW_HIDE
	}
	win.ShowWindow(w.hwnd, cmd)
}

func (w *WindowBase) SetEnabled(b bool) {
	win.EnableWindow(w.hwnd, b)
}

func (w *WindowBase) SetDisabled(disable bool) {
	win.EnableWindow(w.hwnd, !disable)
}

func (w *WindowBase) Close() {
	win.SendMessage(w.hwnd, win.WM_CLOSE, 0, 0)
}

func (w *WindowBase) SetFocus() {
	win.SetFocus(w.hwnd)
}

func (w *WindowBase) Bounds() Rect {
	var rect win.RECT
	win.GetWindowRect(w.hwnd, &rect)
	return rectFromRECT(rect)
}

func (w *WindowBase) SetBounds(value Rect) {
	win.MoveWindow(w.hwnd, int32(value.X), int32(value.Y), int32(value.Width), int32(value.Height), true)
}

func (w *WindowBase) WndProc(msg uint32, wparam, lparam uintptr) uintptr {
	fmt.Println("WidgetBase.WndProc")
	return uintptr(0)
}
