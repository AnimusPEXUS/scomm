package main

import (
	"github.com/gotk3/gotk3/gtk"
)

type UISysTrayMenu struct {
	menu *gtk.Menu
}

func NewUISysTrayMenu() (*UISysTrayMenu, error) {
	self := new(UISysTrayMenu)
	if _, err := gtk.MenuNew(); err != nil {
		return nil, err
	}
	return self, nil
}
