package gowi

import (
	"fmt"
	"github.com/nvsoft/win"
)

// GWidget
type Widget interface {

	AsWidgetBase() *WidgetBase

	InitWidget(window Window, id uintptr)

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

	SetEnabled(b bool)

	SetDisabled(disable bool)

	Close()

	WndProc(msg uint32, wparam, lparam uintptr) uintptr
}

// WidgetBase
type WidgetBase struct {
	hwnd                      win.HWND
	parent                    win.HWND
	idd                       uintptr
}

func (w *WidgetBase) AsWidgetBase() *WidgetBase {
	return w
}

// 设置窗口对应资源Idd
func (w *WidgetBase) InitWidget(window Window, id uintptr) {
	w.idd = id
	h := window.AsWindowBase().Handle()
	fmt.Printf("h=%v\n", h)
	w.hwnd = win.GetDlgItem(h, id)
	fmt.Printf("w.hwnd=%v\n", w.hwnd)
}

func (w *WidgetBase) Handle() win.HWND {
	return w.hwnd
}

func (w *WidgetBase) GetWindowText() string {
	return win.GetWindowText(w.hwnd)
}

func (w *WidgetBase) SetWindowText(title string) {
	fmt.Printf("SetCaption hwnd: %v, %s\n", w.hwnd, title)
	win.SetWindowText(w.hwnd, title)
}

// IconType: 1 - ICON_BIG; 0 - ICON_SMALL
func (w *WidgetBase) SetIcon(iconType int, icon *Icon) {
	if iconType > 1 {
		panic("IconType is invalid")
	}

	win.SendMessage(w.hwnd, win.WM_SETICON, uintptr(iconType), uintptr(icon.Handle()))
}

func (w *WidgetBase) Show() {
	win.ShowWindow(w.hwnd, win.SW_SHOW)
}

func (w *WidgetBase) Hide() {
	win.ShowWindow(w.hwnd, win.SW_HIDE)
}

func (w *WidgetBase) ShowMinimized() {
	win.ShowWindow(w.hwnd, win.SW_MINIMIZE)
}

func (w *WidgetBase) ShowMaximized() {
	win.ShowWindow(w.hwnd, win.SW_MAXIMIZE)
}

func (w *WidgetBase) ShowFullScreen() {
	win.ShowWindow(w.hwnd, win.SW_SHOWMAXIMIZED)
}

func (w *WidgetBase) ShowNormal() {
	win.ShowWindow(w.hwnd, win.SW_SHOWNORMAL)
}

func (w *WidgetBase) IsEnabled() bool {
	return win.IsWindowEnabled(w.hwnd)
}

func (w *WidgetBase) IsVisible() bool {
	return win.IsWindowVisible(w.hwnd)
}

func (w *WidgetBase) SetEnabled(b bool) {
	win.EnableWindow(w.hwnd, b)
}

func (w *WidgetBase) SetDisabled(disable bool) {
	win.EnableWindow(w.hwnd, !disable)
}

func (w *WidgetBase) Close() {
	win.SendMessage(w.hwnd, win.WM_CLOSE, 0, 0)
}

func (w *WidgetBase) WndProc(msg uint32, wparam, lparam uintptr) uintptr {
	fmt.Println("WidgetBase.WndProc")
	return uintptr(0)
}
