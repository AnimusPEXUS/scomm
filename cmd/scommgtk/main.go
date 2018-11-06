package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	win, err := UIMainWindowNew()
	if err != nil {
		log.Fatal(err)
	}
	win.Show()

	gtk.Main()
}
