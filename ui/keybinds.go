package ui

import "github.com/jroimartin/gocui"

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func BindGlobalKeys(g *gocui.Gui) error {
	if err := g.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
		return err
	}
	return nil
}
