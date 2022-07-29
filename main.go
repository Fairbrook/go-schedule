package main

import (
	"fmt"
	"log"

	"github.com/Fairbrook/go-schedule/ui"
	"github.com/jroimartin/gocui"
)

func main() {
	fmt.Print("hola")
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	g.Cursor = true
	g.SetManagerFunc(ui.Layout)
	if err := ui.BindGlobalKeys(g); err != nil {
		log.Panicln(err)
	}
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
