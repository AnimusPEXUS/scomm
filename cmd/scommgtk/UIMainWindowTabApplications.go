package main

import (
	"fmt"
	"time"

	"github.com/AnimusPEXUS/appplugsys"
	"github.com/AnimusPEXUS/utils/worker"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type UIMainWindowTabApplications struct {
	main_window *UIMainWindow

	iv_app_icons       *gtk.IconView
	iv_app_icons_model *gtk.ListStore

	applications_tab_status_refresher *worker.Worker
}

func UIMainWindowTabApplicationsNew(
	builder *gtk.Builder,
	main_window *UIMainWindow,
) (*UIMainWindowTabApplications, error) {

	ret := new(UIMainWindowTabApplications)

	ret.applications_tab_status_refresher = worker.New(ret.statusRefresherTask)

	ret.main_window = main_window

	main_window.window.Connect("destroy", func() { ret.applications_tab_status_refresher.Stop() })

	ret.applications_tab_status_refresher.Start()

	{
		t0, _ := builder.GetObject("iv_app_icons")
		t1, _ := t0.(*gtk.IconView)
		ret.iv_app_icons = t1
	}

	if t, err := gtk.ListStoreNew(
		glib.TYPE_STRING,    // Title
		gdk.PixbufGetType(), // Icon
		glib.TYPE_STRING,    // Pluin Name
		glib.TYPE_STRING,    // Name
		glib.TYPE_STRING,    // Description

	); err != nil {
		return nil, err
	} else {
		ret.iv_app_icons_model = t
	}

	ret.iv_app_icons.Connect(
		"item-activated",
		func(
			iv *gtk.IconView,
			tp *gtk.TreePath,
		) {
			mdl := ret.iv_app_icons_model

			iter, err := mdl.GetIter(tp)
			if err != nil {
				fmt.Println("error", err)
				return
			}

			v, err := mdl.GetValue(iter, 2)
			if err != nil {
				fmt.Println("error", err)
				return
			}

			plug_name, err := v.GetString()
			if err != nil {
				fmt.Println("error", err)
				return
			}

			v, err = mdl.GetValue(iter, 3)
			if err != nil {
				fmt.Println("error", err)
				return
			}

			name, err := v.GetString()
			if err != nil {
				fmt.Println("error", err)
				return
			}

			err = ret.main_window.controller.appplugsys.DisplayApplication(
				plug_name,
				name,
			)
			if err != nil {

				d := gtk.MessageDialogNew(
					ret.main_window.window,
					0,
					gtk.MESSAGE_ERROR,
					gtk.BUTTONS_OK,
					"Application Display Error:\n"+err.Error(),
				)
				d.Run()
				d.Close()
			}
			return

		},
	)

	ret.iv_app_icons.SetTextColumn(0)
	ret.iv_app_icons.SetPixbufColumn(1)
	ret.iv_app_icons.SetModel(ret.iv_app_icons_model)

	return ret, nil
}

func (self *UIMainWindowTabApplications) statusRefresherTask(

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

				mdl := self.iv_app_icons_model

				table := self.main_window.controller.appplugsys.ApplicationInfoTable()

				table_have_this := func(plugin_name, name string) bool {
					for _, i := range table {
						if i.PluginName == plugin_name && i.Name == name {
							return true
						}
					}
					return false
				}

				table_get_item := func(plugin_name, name string) *appplugsys.AppPlugSysApplicationDisplayItem {
					for _, i := range table {
						if i.PluginName == plugin_name && i.Name == name {
							return i
						}
					}
					return nil
				}

				found := make([][2]string, 0)

				{
					iter, ok := mdl.GetIterFirst()
					for ok {

						val, err := mdl.GetValue(iter, 2)
						if err != nil {
							panic(err)
						}

						plugin_name, err := val.GetString()
						if err != nil {
							panic(err)
						}

						val, err = mdl.GetValue(iter, 3)
						if err != nil {
							panic(err)
						}

						name, err := val.GetString()
						if err != nil {
							panic(err)
						}

						found = append(found, [2]string{plugin_name, name})

						if !table_have_this(plugin_name, name) {
							ok = mdl.Remove(iter)
						} else {
							it := table_get_item(plugin_name, name)
							if it == nil {
								panic("programming error")
							}

							pbfl, err := gdk.PixbufLoaderNew()
							if err != nil {
								panic(err)
							}
							pbf, err := pbfl.WriteAndReturnPixbuf(it.Icon)
							if err != nil {
								panic(err)
							}

							mdl.Set(
								iter,
								[]int{0, 1, 2, 3, 4, 5},
								[]interface{}{
									it.Title,
									pbf,
									it.PluginName,
									it.Name,
									it.Description,
								},
							)
							ok = mdl.IterNext(iter)
						}

					}
				}

				not_found := make([][2]string, 0)

			search:
				for _, v := range table {
					for _, i := range found {
						if v.PluginName == i[0] && v.Name == i[1] {
							continue search
						}
					}
					not_found = append(not_found, [2]string{v.PluginName, v.Name})
				}

				for _, i := range not_found {
					it := table_get_item(i[0], i[1])
					if it == nil {
						panic("programming error")
					}
					iter := mdl.Append()

					pbfl, err := gdk.PixbufLoaderNew()
					if err != nil {
						panic(err)
					}
					pbf, err := pbfl.WriteAndReturnPixbuf(it.Icon)
					if err != nil {
						panic(err)
					}

					err = mdl.Set(
						iter,
						[]int{0, 1, 2, 3, 4},
						[]interface{}{
							it.Title,
							pbf,
							it.PluginName,
							it.Name,
							it.Description,
						},
					)
					if err != nil {
						panic(err)
					}
				}

			},
		)

		<-c

	}
}
