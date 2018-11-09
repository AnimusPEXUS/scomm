package main

import (
	"github.com/gotk3/gotk3/gtk"
)

type UIMainWindow struct {
	controller *Controller

	tab_applications *UIMainWindowTabApplications
	tab_plugins      *UIMainWindowTabPlugins

	window *gtk.Window

	/*
		button_dnet     *gtk.Button
		button_storage  *gtk.Button
		button_keys     *gtk.Button
		button_certs    *gtk.Button
		button_networks *gtk.Button
		button_services *gtk.Button
	*/
	//	button_online   *gtk.ToggleButton
	button_home_sep *gtk.Separator
	button_home     *gtk.Button
	mi_storage      *gtk.MenuItem
	mi_about        *gtk.MenuItem
	box_for_log     *gtk.Box

	notebook_main *gtk.Notebook

	//logger *gologger.Logger
}

func UIMainWindowNew(controller *Controller) (*UIMainWindow, error) {

	ret := new(UIMainWindow)

	//ret.logger = gologger.New()

	ret.controller = controller

	builder, err := gtk.BuilderNew()
	if err != nil {
		return nil, err
	}

	data, err := uiMainGladeBytes()
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

	if res, err := UIMainWindowTabApplicationsNew(
		builder,
		ret,
	); err != nil {
		return nil, err
	} else {
		ret.tab_applications = res
	}

	if res, err := UIMainWindowTabPluginsNew(
		builder,
		ret,
	); err != nil {
		return nil, err
	} else {
		ret.tab_plugins = res
	}

	//	{
	//		t0, _ := builder.GetObject("button_online")
	//		t1, _ := t0.(*gtk.ToggleButton)
	//		ret.button_online = t1
	//	}

	{
		t0, _ := builder.GetObject("box_for_log")
		t1, _ := t0.(*gtk.Box)
		ret.box_for_log = t1
	}

	//	ret.box_for_log.PackEnd(
	//		log_viewer.UILogViewerNew(ret.controller.logger).GetRoot(),
	//		true,
	//		true,
	//		0,
	//	)

	//	ret.button_online.Connect(
	//		"toggled",
	//		func() {
	//			fmt.Println("toggled")
	//			ret.controller.logger.Info("testing logger")
	//		},
	//	)

	return ret, nil
}

func (self *UIMainWindow) Show() {
	self.window.ShowAll()
	return
}
