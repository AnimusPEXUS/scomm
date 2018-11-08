package main

import (
	"fmt"

	"github.com/AnimusPEXUS/appplugsys"
	"github.com/gotk3/gotk3/gtk"
)

type UIPluginAcceptorWindow struct {
	builtin bool
	name    string // for builtin
	sha512  string // for external

	last_search_result *appplugsys.PluginSearcherSearchResultItem

	main_window *UIMainWindow

	window *gtk.Window

	label_builtin        *gtk.Label
	label_name           *gtk.Label
	label_title          *gtk.Label
	label_description    *gtk.Label
	label_filepath       *gtk.Label
	label_checksum       *gtk.Label
	button_research      *gtk.Button
	button_open_external *gtk.Button
	button_accept_module *gtk.Button
	button_cancel        *gtk.Button
}

func UIPluginAcceptorWindowNew(
	main_window *UIMainWindow,
	builtin bool,
	name string, // for builtin
	sha512 string, // for external
) (*UIPluginAcceptorWindow, error) {

	ret := new(UIPluginAcceptorWindow)

	ret.builtin = builtin
	ret.name = name
	ret.sha512 = sha512

	ret.main_window = main_window

	builder, err := gtk.BuilderNew()
	if err != nil {
		return nil, err
	}

	data, err := uiPluginAcceptorGladeBytes()
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
		t0, _ := builder.GetObject("label_builtin")
		t1, _ := t0.(*gtk.Label)
		ret.label_builtin = t1
	}

	{
		t0, _ := builder.GetObject("label_name")
		t1, _ := t0.(*gtk.Label)
		ret.label_name = t1
	}

	{
		t0, _ := builder.GetObject("label_title")
		t1, _ := t0.(*gtk.Label)
		ret.label_title = t1
	}

	{
		t0, _ := builder.GetObject("label_description")
		t1, _ := t0.(*gtk.Label)
		ret.label_description = t1
	}

	{
		t0, _ := builder.GetObject("label_filepath")
		t1, _ := t0.(*gtk.Label)
		ret.label_filepath = t1
	}

	{
		t0, _ := builder.GetObject("label_checksum")
		t1, _ := t0.(*gtk.Label)
		ret.label_checksum = t1
	}

	{
		t0, _ := builder.GetObject("button_research")
		t1, _ := t0.(*gtk.Button)
		ret.button_research = t1
	}

	{
		t0, _ := builder.GetObject("button_open_external")
		t1, _ := t0.(*gtk.Button)
		ret.button_open_external = t1
	}

	{
		t0, _ := builder.GetObject("button_accept_module")
		t1, _ := t0.(*gtk.Button)
		ret.button_accept_module = t1
	}

	{
		t0, _ := builder.GetObject("button_cancel")
		t1, _ := t0.(*gtk.Button)
		ret.button_cancel = t1
	}

	ret.button_cancel.Connect(
		"clicked",
		func() {
			ret.window.Close()
		},
	)

	ret.button_research.Connect(
		"clicked",
		func() {
			ret.ReSearchModuleFile()
		},
	)

	ret.button_accept_module.Connect(
		"clicked",
		func() {

			if ret.last_search_result == nil {
				d := gtk.MessageDialogNew(
					ret.window,
					0,
					gtk.MESSAGE_ERROR,
					gtk.BUTTONS_OK,
					"Search result is not acceptable",
				)
				d.Run()
				d.Close()
				return
			}

			plug_wrap, err := appplugsys.NewPluginWrapFromSearchItem(
				ret.last_search_result,
				ret.main_window.controller.appplugsys,
			)
			if err != nil {
				d := gtk.MessageDialogNew(
					ret.window,
					0,
					gtk.MESSAGE_ERROR,
					gtk.BUTTONS_OK,
					"Error wrapping plugin: "+err.Error(),
				)
				d.Run()
				d.Close()
				return
			}

			err = ret.main_window.controller.appplugsys.AcceptPlugin(plug_wrap)

			defer ret.window.Close()

			if err != nil {
				d := gtk.MessageDialogNew(
					ret.window,
					0,
					gtk.MESSAGE_ERROR,
					gtk.BUTTONS_OK,
					"Error accepting plugin: "+err.Error(),
				)
				d.Run()
				d.Close()
				return
			}
		},
	)

	{

		if !builtin {
			ret.ReSearchModuleFile()
		}

		ret.RefreshStates()

	}

	return ret, nil

}

func (self *UIPluginAcceptorWindow) Show() {
	self.window.ShowAll()
	return
}

func (self *UIPluginAcceptorWindow) ReSearchModuleFile() {

	var search_res *appplugsys.PluginSearcherSearchResultItem
	var err error

	if self.builtin {
		search_res, err = self.main_window.controller.plugin_searcher.
			FindBuiltIn(self.name)
	} else {
		search_res, err = self.main_window.controller.plugin_searcher.
			FindBySha512(self.sha512)
	}
	if err != nil {
		self.last_search_result = nil
		// TODO: display search error
	} else {
		self.last_search_result = search_res
	}

	self.RefreshStates()
}

func (self *UIPluginAcceptorWindow) RefreshLabels() {
	self.label_builtin.SetText(fmt.Sprintf("%v", self.builtin))

	if self.builtin {
		self.label_filepath.SetText("N/A")
	} else {
		if self.last_search_result == nil {
			self.label_filepath.SetText("not found")
		} else {
			self.label_filepath.SetText(self.last_search_result.Path)
		}
	}
	if self.builtin {
		self.label_checksum.SetText("N/A")
	} else {
		self.label_checksum.SetText(self.sha512)
	}
	return
}

func (self *UIPluginAcceptorWindow) RefreshEnableds() {
}

func (self *UIPluginAcceptorWindow) RefreshStates() {
	self.RefreshLabels()
	self.RefreshEnableds()
}
