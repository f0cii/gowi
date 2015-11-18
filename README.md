# gowi
Gui for golang, UI created by ResEdit

# Code

import (
    "github.com/lroc/gowi"
    . "github.com/lroc/gowi/declarative"
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

How to design：

![ResEdit](https://raw.githubusercontent.com/lroc/gowi/master/doc/resedit.png)

Run：

![Running screenshot](https://raw.githubusercontent.com/lroc/gowi/master/doc/screenshot1.png)
