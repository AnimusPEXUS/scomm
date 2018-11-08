package main

import (
	"fmt"
	"sync"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type UIMainWindowTabApplications struct {
	main_window *UIMainWindow

	button_acc_plug_refresh     *gtk.Button
	button_acc_plug_enable      *gtk.Button
	button_acc_plug_disable     *gtk.Button
	button_acc_plug_start       *gtk.Button
	button_acc_plug_stop        *gtk.Button
	button_acc_plug_unregister  *gtk.Button
	button_acc_plug_rekey       *gtk.Button
	button_acc_plug_show_window *gtk.Button

	button_ava_plug_refresh *gtk.Button
	button_ava_plug_accept  *gtk.Button

	tw_acc *gtk.TreeView
	tw_ava *gtk.TreeView
	ls_acc *gtk.ListStore
	ls_ava *gtk.ListStore

	refresh_plug_preset_list_item_lock *sync.Mutex
}

func UIMainWindowTabApplicationsNew(
	builder *gtk.Builder,
	main_window *UIMainWindow,
) (*UIMainWindowTabApplications, error) {

	ret := new(UIMainWindowTabApplications)

	ret.refresh_plug_preset_list_item_lock = &sync.Mutex{}

	ret.main_window = main_window

	{
		t0, _ := builder.GetObject("button_acc_plug_refresh")
		t1, _ := t0.(*gtk.Button)
		ret.button_acc_plug_refresh = t1
	}

	{
		t0, _ := builder.GetObject("button_acc_plug_enable")
		t1, _ := t0.(*gtk.Button)
		ret.button_acc_plug_enable = t1
	}

	{
		t0, _ := builder.GetObject("button_acc_plug_disable")
		t1, _ := t0.(*gtk.Button)
		ret.button_acc_plug_disable = t1
	}

	{
		t0, _ := builder.GetObject("button_acc_plug_start")
		t1, _ := t0.(*gtk.Button)
		ret.button_acc_plug_start = t1
	}

	{
		t0, _ := builder.GetObject("button_acc_plug_stop")
		t1, _ := t0.(*gtk.Button)
		ret.button_acc_plug_stop = t1
	}

	{
		t0, _ := builder.GetObject("button_acc_plug_unregister")
		t1, _ := t0.(*gtk.Button)
		ret.button_acc_plug_unregister = t1
	}

	{
		t0, _ := builder.GetObject("button_acc_plug_rekey")
		t1, _ := t0.(*gtk.Button)
		ret.button_acc_plug_rekey = t1
	}

	{
		t0, _ := builder.GetObject("button_acc_plug_show_window")
		t1, _ := t0.(*gtk.Button)
		ret.button_acc_plug_show_window = t1
	}

	{
		t0, _ := builder.GetObject("button_ava_plug_refresh")
		t1, _ := t0.(*gtk.Button)
		ret.button_ava_plug_refresh = t1
	}

	{
		t0, _ := builder.GetObject("button_ava_plug_accept")
		t1, _ := t0.(*gtk.Button)
		ret.button_ava_plug_accept = t1
	}

	{
		t0, _ := builder.GetObject("tw_acc")
		t1, _ := t0.(*gtk.TreeView)
		ret.tw_acc = t1
	}

	{
		t0, _ := builder.GetObject("tw_ava")
		t1, _ := t0.(*gtk.TreeView)
		ret.tw_ava = t1
	}

	{
		ret.ls_acc, _ = gtk.ListStoreNew(
			glib.TYPE_STRING,  // Name
			glib.TYPE_BOOLEAN, // builtin?
			glib.TYPE_BOOLEAN, // enabled?
			glib.TYPE_STRING,  // Status
			glib.TYPE_STRING,  // Checksum
			glib.TYPE_STRING,  // Last ReKey time
		)

		ret.ls_ava, _ = gtk.ListStoreNew(
			glib.TYPE_STRING,  // Name
			glib.TYPE_BOOLEAN, // builtin?
			glib.TYPE_STRING,  // Checksum
			glib.TYPE_STRING,  // Path
			glib.TYPE_STRING,  // Descr
		)

		ret.tw_acc.SetModel(ret.ls_acc)
		ret.tw_ava.SetModel(ret.ls_ava)
	}

	{
		{
			// setup columns in tw_acc
			{
				rend, _ := gtk.CellRendererTextNew()
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Name",
					rend,
					"text",
					0,
				)
				ret.tw_acc.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererToggleNew()
				rend.SetActivatable(false)
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"BuiltIn?",
					rend,
					"active",
					1,
				)
				ret.tw_acc.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererToggleNew()
				rend.SetActivatable(false)
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Enabled?",
					rend,
					"active",
					2,
				)
				ret.tw_acc.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererTextNew()
				//rend.SetActivatable(false)
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Status",
					rend,
					"text",
					3,
				)
				ret.tw_acc.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererTextNew()
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Checksum",
					rend,
					"text",
					4,
				)
				ret.tw_acc.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererTextNew()
				//rend.SetActivatable(false)
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Last ReKey Time",
					rend,
					"text",
					5,
				)
				ret.tw_acc.AppendColumn(column)
			}
		}

		{
			// setup columns in tw_application_modules
			{
				rend, _ := gtk.CellRendererTextNew()
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Name",
					rend,
					"text",
					0,
				)
				ret.tw_ava.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererToggleNew()
				rend.SetActivatable(false)
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"BuiltIn?",
					rend,
					"active",
					1,
				)
				ret.tw_ava.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererTextNew()
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Checksum",
					rend,
					"text",
					2,
				)
				ret.tw_ava.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererTextNew()
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Path",
					rend,
					"text",
					3,
				)
				ret.tw_ava.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererTextNew()
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Description",
					rend,
					"text",
					4,
				)
				ret.tw_ava.AppendColumn(column)
			}
		}
	}

	//		ret.button_refresh_application_presets.Connect(
	//			"clicked",
	//			func() {
	//				ret.main_window.controller.application_controller.
	//					RefreshAllAcceptedApplicationListItems(false)
	//			},
	//		)

	ret.button_ava_plug_refresh.Connect(
		"clicked",
		func() {

			fmt.Println("clicked")

			mdl := ret.ls_ava

			{
				iter, ok := mdl.GetIterFirst()
				for {
					if !ok {
						break
					}
					ok = mdl.Remove(iter)
				}
			}

			plugins, err := ret.main_window.controller.plugin_searcher.FindAll()
			if err != nil {
				fmt.Println("error:", err.Error())
				return
			}

			fmt.Println("count", len(plugins))
			fmt.Println("plugins", plugins)

			for _, i := range plugins {
				if i.BuiltIn {
					iter := mdl.Append()

					//					mod, err := i.Mod()
					//					if err != nil {
					//						fmt.Printf("error executing .Mod(): %s\n", err.Error())
					//						mdl.Remove(iter)
					//						continue
					//					}

					mdl.Set(
						iter,
						[]int{0, 1, 2, 3, 4},
						[]interface{}{
							i.Name,
							true,
							"N/A",
							"N/A",
							i.Plugin.Description,
						},
					)
				}
			}
		},
	)

	ret.button_ava_plug_accept.Connect(
		"clicked",
		func() {

			var (
				builtin bool
				name    string
				sha512  string
			)

			sel, _ := ret.tw_ava.GetSelection()
			model, iter, ok := sel.GetSelected()

			if !ok {
				return
			}

			{
				val, _ := model.(*gtk.TreeModel).GetValue(iter, 0)
				name, _ = val.GetString()
			}

			{
				val, _ := model.(*gtk.TreeModel).GetValue(iter, 1)
				builtin_t, _ := val.GoValue()
				builtin, _ = builtin_t.(bool)
			}

			{
				val, _ := model.(*gtk.TreeModel).GetValue(iter, 2)
				sha512, _ = val.GetString()
			}

			//			fmt.Println("builtin", builtin)
			//			fmt.Println("name", name)
			//			fmt.Println("sha512", sha512)

			acc_window, err := UIPluginAcceptorWindowNew(
				ret.main_window,
				builtin,
				name,
				sha512,
			)
			if err != nil {
				fmt.Println("error", err.Error())
			}

			acc_window.Show()

		},
	)

	//	ret.button_enable_application_preset.Connect(
	//		"clicked",
	//		func() {
	//			value, ok := ret.GetSelectedPresetName()
	//			if ok {
	//				value_obj, err := dnet.ModuleNameNew(value)
	//				if err != nil {
	//					panic("programming error")
	//				}
	//				ret.main_window.controller.application_controller.SetModuleEnabled(
	//					value_obj,
	//					true,
	//					true,
	//				)
	//			}
	//		},
	//	)

	//	ret.button_disable_application_preset.Connect(
	//		"clicked",
	//		func() {
	//			value, ok := ret.GetSelectedPresetName()
	//			if ok {
	//				value_obj, err := dnet.ModuleNameNew(value)
	//				if err != nil {
	//					panic("programming error")
	//				}
	//				ret.main_window.controller.application_controller.SetModuleEnabled(
	//					value_obj,
	//					false,
	//					true,
	//				)
	//			}
	//		},
	//	)

	//	ret.button_open_window_application_preset.Connect(
	//		"clicked",
	//		func() {
	//			value, ok := ret.GetSelectedPresetName()
	//			if ok {

	//				value_obj, err := dnet.ModuleNameNew(value)
	//				if err != nil {
	//					panic("error")
	//				}

	//				err = ret.main_window.controller.
	//					application_controller.ModuleShowUI(value_obj)

	//				if err != nil {

	//					d := gtk.MessageDialogNew(
	//						ret.main_window.win,
	//						0,
	//						gtk.MESSAGE_ERROR,
	//						gtk.BUTTONS_OK,
	//						"Trying to show window resulted in error: "+err.Error(),
	//					)
	//					d.Run()
	//					d.Destroy()

	//				}
	//			}
	//		},
	//	)

	return ret, nil
}

//func (self *UIMainWindowTabApplications) GetSelectedPresetName() (
//	string,
//	bool,
//) {
//	sel, _ := self.tw_application_presets.GetSelection()
//	model, iter, ok := sel.GetSelected()
//	if ok {
//		val, _ := model.(*gtk.TreeModel).GetValue(iter, 0)
//		val_str, _ := val.GetString()
//		return val_str, true
//	}
//	return "", false
//}

/*
func (self *UIMainWindowTabApplications) GetSelectedModuleName() (
	string,
	bool,
) {
	sel, _ := self.tw_application_modules.GetSelection()
	model, iter, ok := sel.GetSelected()
	if ok {
		val, _ := model.(*gtk.TreeModel).GetValue(iter, 0)
		val_str, _ := val.GetString()
		return val_str, true
	}
	return "", false
}
*/

//func (self *UIMainWindowTabApplications) RefreshAppPresetListItem(
//	item_name string,
//	no_db bool,
//) error {

//	go func() {
//		// chann := make(chan bool, 1)

//		self.refresh_plug_preset_list_item_lock.Lock()
//		defer self.refresh_plug_preset_list_item_lock.Unlock()

//		glib.IdleAdd(
//			func() {
//				// defer func() { chann <- true }()
//				item_name_obj := dnet.ModuleNameNewF(item_name)

//				presets_mdl := self.application_presets
//				wrappers := self.main_window.controller.
//					application_controller.application_wrappers

//				delete_preset := true
//				var preset_iter *gtk.TreeIter

//				for i, _ := range wrappers {
//					if i == item_name {
//						delete_preset = false
//						break
//					}
//				}

//				iter, ok := presets_mdl.GetIterFirst()
//				for ok {
//					val, _ := presets_mdl.GetValue(iter, 0)
//					val_str, _ := val.GetString()
//					if val_str == item_name {
//						preset_iter = iter
//						break
//					}
//					ok = presets_mdl.IterNext(iter)
//				}

//				if preset_iter == nil {
//					preset_iter = presets_mdl.Append()
//				}

//				if delete_preset {
//					if preset_iter != nil {
//						presets_mdl.Remove(preset_iter)
//					}
//				} else {

//					var stat *ApplicationStatus = nil
//					cs := "N/A"
//					if !no_db {
//						var err error
//						stat, err = self.main_window.controller.
//							application_controller.GetModuleStatus(item_name_obj)
//						if err != nil {
//							return // err
//						}
//						if !stat.Builtin {
//							cs = stat.Checksum
//						}
//					}

//					module_running_status := "N/A"
//					if self.main_window != nil &&
//						self.main_window.controller != nil &&
//						self.main_window.controller.application_controller != nil &&
//						self.main_window.controller.
//							application_controller.application_wrappers != nil {
//						if t, ok := self.main_window.controller.
//							application_controller.application_wrappers[item_name]; ok {
//							if t.Module.IsWorker() == false {
//								module_running_status = "Not a Worker"
//							} else {
//								if t.Instance != nil {
//									module_running_status = t.Instance.Status().StringT()
//								}
//							}
//						}
//					}

//					if !no_db {
//						presets_mdl.Set(
//							preset_iter,
//							[]int{0, 1, 2, 3, 4, 5},
//							[]interface{}{
//								item_name,
//								stat.Builtin,
//								stat.Enabled,
//								module_running_status,
//								cs,
//								stat.LastDBReKey.String(),
//							},
//						)
//					} else {
//						presets_mdl.Set(
//							preset_iter,
//							[]int{0, 3, 4},
//							[]interface{}{
//								item_name,
//								module_running_status,
//								cs,
//							},
//						)
//					}
//				}

//			},
//		)
//		// <-chann
//	}()
//	return nil
//}
