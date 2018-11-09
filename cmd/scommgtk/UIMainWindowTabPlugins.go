package main

import (
	"fmt"
	"time"

	"github.com/AnimusPEXUS/utils/worker"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type UIMainWindowTabPlugins struct {
	main_window *UIMainWindow

	//	button_acc_plug_refresh     *gtk.Button
	button_acc_plug_enable      *gtk.Button
	button_acc_plug_disable     *gtk.Button
	button_acc_plug_start       *gtk.Button
	button_acc_plug_stop        *gtk.Button
	button_acc_plug_unregister  *gtk.Button
	button_acc_plug_rekey       *gtk.Button
	button_acc_plug_show_window *gtk.Button

	button_ava_plug_refresh *gtk.Button
	button_ava_plug_accept  *gtk.Button

	tv_acc *gtk.TreeView
	tv_ava *gtk.TreeView
	ls_acc *gtk.ListStore
	ls_ava *gtk.ListStore

	status_refresher *worker.Worker
}

func UIMainWindowTabPluginsNew(
	builder *gtk.Builder,
	main_window *UIMainWindow,
) (*UIMainWindowTabPlugins, error) {

	ret := new(UIMainWindowTabPlugins)

	ret.status_refresher = worker.New(ret.statusRefresherTask)

	ret.main_window = main_window

	main_window.window.Connect("destroy", func() { ret.status_refresher.Stop() })

	ret.status_refresher.Start()

	//	{
	//		t0, _ := builder.GetObject("button_acc_plug_refresh")
	//		t1, _ := t0.(*gtk.Button)
	//		ret.button_acc_plug_refresh = t1
	//	}

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
		t0, _ := builder.GetObject("tv_acc")
		t1, _ := t0.(*gtk.TreeView)
		ret.tv_acc = t1
	}

	{
		t0, _ := builder.GetObject("tv_ava")
		t1, _ := t0.(*gtk.TreeView)
		ret.tv_ava = t1
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

		ret.tv_acc.SetModel(ret.ls_acc)
		ret.tv_ava.SetModel(ret.ls_ava)
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
				ret.tv_acc.AppendColumn(column)
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
				ret.tv_acc.AppendColumn(column)
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
				ret.tv_acc.AppendColumn(column)
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
				ret.tv_acc.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererTextNew()
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Checksum",
					rend,
					"text",
					4,
				)
				ret.tv_acc.AppendColumn(column)
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
				ret.tv_acc.AppendColumn(column)
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
				ret.tv_ava.AppendColumn(column)
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
				ret.tv_ava.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererTextNew()
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Checksum",
					rend,
					"text",
					2,
				)
				ret.tv_ava.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererTextNew()
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Path",
					rend,
					"text",
					3,
				)
				ret.tv_ava.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererTextNew()
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Description",
					rend,
					"text",
					4,
				)
				ret.tv_ava.AppendColumn(column)
			}
		}
	}

	ret.button_ava_plug_refresh.Connect(
		"clicked",
		func() {

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

			sel, _ := ret.tv_ava.GetSelection()
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

	return ret, nil
}

func (self *UIMainWindowTabPlugins) statusRefresherTask(

	set_starting func(),
	set_working func(),
	set_stopping func(),
	set_stopped func(),

	is_stop_flag func() bool,

) {
	set_working()
	for {
		time.Sleep(time.Second)
		if is_stop_flag() {
			break
		}

		c := make(chan bool)

		glib.IdleAdd(
			func() {

				defer func() { c <- true }()

				mdl := self.ls_acc

				table := self.main_window.controller.appplugsys.PluginInfoTable()

				found := make([]string, 0)

				{
					iter, ok := mdl.GetIterFirst()
					for {
						if !ok {
							break
						}

						val, _ := mdl.GetValue(iter, 0)
						val_str, _ := val.GetString()

						found = append(found, val_str)

						if _, ok2 := table[val_str]; !ok2 {
							ok = mdl.Remove(iter)
						} else {
							mdl.Set(
								iter,
								[]int{0, 1, 2, 3, 4, 5},
								[]interface{}{
									table[val_str].Name,
									table[val_str].BuiltIn,
									table[val_str].Enabled,
									table[val_str].WorkerStatus,
									table[val_str].Sha512,
									table[val_str].LastDBReKey,
								},
							)
							ok = mdl.IterNext(iter)
						}

					}
				}

				not_found := make([]string, 0)

			search:
				for k, _ := range table {
					for _, i := range found {
						if k == i {
							continue search
						}
					}
					not_found = append(not_found, k)
				}

				for _, i := range not_found {
					iter := mdl.Append()
					mdl.Set(
						iter,
						[]int{0, 1, 2, 3, 4, 5},
						[]interface{}{
							table[i].Name,
							table[i].BuiltIn,
							table[i].Enabled,
							table[i].WorkerStatus,
							table[i].Sha512,
							table[i].LastDBReKey,
						},
					)
				}

			},
		)

		<-c

	}
}
