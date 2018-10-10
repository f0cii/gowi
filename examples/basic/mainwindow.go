package main

import (
	"fmt"

	"github.com/sumorf/gowi"
)

type MainWindow struct {
	gowi.MainWindow
	button1 *gowi.Button
	label1  *gowi.Label
}

func NewMainWindow() *MainWindow {
	w := &MainWindow{}

	w.Create(IDD_MAIN)

	b1 := gowi.NewButton(w.MainWindow, IDD_BUTTON1)
	l1 := gowi.NewLabel(w.MainWindow, IDD_LABEL1)

	b1.OnClicked = w.OnButton1Clicked

	l1.SetWindowText("Hello！")

	w.button1 = b1
	w.label1 = l1

	return w
}

func (w *MainWindow) OnButton1Clicked() {
	fmt.Println("OnButton1Clicked")
	text := w.button1.GetWindowText()
	if text == "Start" {
		w.button1.SetWindowText("Stop")
	} else {
		w.button1.SetWindowText("Start")
	}
}
