package gowi

import (
	"github.com/nvsoft/win"
)

var hInst win.HINSTANCE

// 获取应用实例句柄
func GetAppInstance() win.HINSTANCE {
	return hInst
}

type Engine struct {
}

// 运行主消息循环
func (e *Engine) Run() int {
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

func New() *Engine {
	engine := new(Engine)
	return engine
}
