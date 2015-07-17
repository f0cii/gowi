package main

import (
	"fmt"
	"github.com/nvsoft/gowi"
	//. "github.com/nvsoft/gowi/declarative"
)

// 主窗口
type MainWindow struct {
	gowi.MainWindow
	button1 *gowi.Button
	label1 *gowi.Label
	//Controls []Widget
}

// 创建主窗口实例
func NewMainWindow() *MainWindow {
	w := &MainWindow{
		//Controls:[]gowi.Widget{
		//	Button{ID:2,OnClicked:func(){
		//		fmt.Printf("okk\n");
		//	}},
		//},
	}

	w.Create(IDD_MAIN)

	b1 := gowi.NewButton(w, IDD_BUTTON1)
	l1 := gowi.NewLabel(w, IDD_LABEL1)

	w.Connect(b1.Clicked, w.OnButton1Clicked)

	l1.SetWindowText("亲，你好！")

	w.button1 = b1
	w.label1 = l1

	return w
}

// 按钮点击事件
func (w *MainWindow) OnButton1Clicked() {
	fmt.Println("OnButton1Clicked")
	text := w.button1.GetWindowText()
	if text == "开始" {
		w.button1.SetWindowText("停止")
	} else {
		w.button1.SetWindowText("开始")
	}
}

