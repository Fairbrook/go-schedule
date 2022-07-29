package ui

import (

	"github.com/jroimartin/gocui"
)

var names = [7]string{"lunes", "martes", "miercoles", "jueves", "viernes", "sabado", "domingo"}

func newRepetiveTask(g *gocui.Gui, v *gocui.View) error {
	posX, posY, maxX, maxY, _ := g.ViewPosition(v.Name())
    if v, err := g.SetView("test", posX+1, posY+2, maxX-1, maxY-1); err!=nil{
		if err != gocui.ErrUnknownView {
			return err
		}
        v.Title = "test"
        return nil
    }
	return nil
}

func drawDay(name string, g *gocui.Gui, offset int, width int, height int) error {
	if v, err := g.SetView(name, int(offset), 0, int(offset+width), int(height)); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		g.SetKeybinding(name, 'n', gocui.ModNone, newRepetiveTask)
		v.Title = name
		g.SetCurrentView(name)
	}
	return nil
}

func drawDays(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	dayWidth := maxX / 7
	for i := 0; i < 7; i++ {
		if err := drawDay(names[i], g, dayWidth*i+1, dayWidth, maxY-1); err != nil {
			return err
		}
	}
	return nil
}

func Layout(g *gocui.Gui) error {
	return drawDays(g)
}
