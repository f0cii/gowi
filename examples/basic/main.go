package main

import (
	"fmt"
	"github.com/nvsoft/gowi"
)

func main() {
	fmt.Println("init")

	engine := gowi.New()

	w := NewMainWindow()
	w.SetWindowText("Go程序示例")
	w.Show()

	engine.Run()
}
