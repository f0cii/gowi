# gowi
Gui for golang, UI created by ResEdit

# Code

import (
    "github.com/nvsoft/gowi"
    . "github.com/nvsoft/gowi/declarative"
    "fmt"
)

type MyMainWindow struct {
    *gowi.MainWindow
    label1 *gowi.Label
    button1 *gowi.Button
}

func main() {
    app := gowi.New()

    mw := new(MyMainWindow)

    m := MainWindow{
        AssignTo:&mw.MainWindow,
        ID:1000,
        Controls:[]Control{
            Button{
                AssignTo:&mw.button1,
                ID:1001,
                OnClicked:func() {
                    fmt.Println("Button clicked.")
                    mw.SetWindowText("Clicked")
                },
            },
        },
    }
    m.Create()
    m.Show()

    mw.SetWindowText("Hello")
    app.Run()
}

# Screenshots

设计：

![ResEdit](https://raw.githubusercontent.com/nvsoft/gowi/master/doc/resedit.png)

运行：

![运行截图](https://raw.githubusercontent.com/nvsoft/gowi/master/doc/screenshot1.png)
