package main

import (
	"github.com/gotk3/gotk3/gtk"
)

type UIWindowStorageLoad struct {
	window *gtk.Window

	button_load  *gtk.Button
	button_close *gtk.Button

	entry_name *gtk.Entry
	entry_key  *gtk.Entry
}

func UIWindowStorageLoadNew(preset_entry_name string) (*UIWindowStorageLoad, error) {

	ret := new(UIWindowStorageLoad)

	builder, err := gtk.BuilderNew()
	if err != nil {
		return nil, err
	}

	data, err := uiStorageLoadGladeBytes()
	if err != nil {
		return nil, err
	}

	err = builder.AddFromString(string(data))
	if err != nil {
		return nil, err
	}

	{
		t0, _ := builder.GetObject("window")
		t1, _ := t0.(*gtk.Window)
		ret.window = t1
	}

	{
		t0, _ := builder.GetObject("button_load")
		t1, _ := t0.(*gtk.Button)
		ret.button_load = t1
	}

	{
		t0, _ := builder.GetObject("entry_name")
		t1, _ := t0.(*gtk.Entry)
		ret.entry_name = t1
	}

	{
		t0, _ := builder.GetObject("entry_key")
		t1, _ := t0.(*gtk.Entry)
		ret.entry_key = t1
	}

	ret.entry_name.SetText(preset_entry_name)

	ret.button_load.Connect(
		"clicked",
		func(
			button *gtk.Button,
			win *UIWindowStorageLoad,
		) {
			name, err := win.entry_name.GetText()
			if err != nil {
				panic(err.Error)
			}
			key, err := win.entry_key.GetText()
			if err != nil {
				panic(err.Error)
			}
			controller, err := NewController(
				name,
				key,
			)
			if err != nil {
				d := gtk.MessageDialogNew(
					ret.window,
					0,
					gtk.MESSAGE_ERROR,
					gtk.BUTTONS_OK,
					"Error creating controller: "+err.Error(),
				)
				d.Run()
				d.Destroy()
				return
			}
			controller.ShowMainWindow()
			win.window.Destroy()
		},
		ret,
	)

	return ret, nil
}

func (self *UIWindowStorageLoad) Show() {
	self.window.ShowAll()
}
