package main

import (
	"time"

	"github.com/AnimusPEXUS/utils/worker"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type UIMainWindowTabApplications struct {
	main_window *UIMainWindow

	iv_app_icons *gtk.IconView

	status_refresher *worker.Worker
}

func UIMainWindowTabApplicationsNew(
	builder *gtk.Builder,
	main_window *UIMainWindow,
) (*UIMainWindowTabApplications, error) {

	ret := new(UIMainWindowTabApplications)

	ret.status_refresher = worker.New(ret.statusRefresherTask)

	ret.main_window = main_window

	main_window.window.Connect("destroy", func() { ret.status_refresher.Stop() })

	ret.status_refresher.Start()

	//	{
	//		ret.ls_acc, _ = gtk.ListStoreNew(
	//			glib.TYPE_STRING,  // Name
	//			glib.TYPE_BOOLEAN, // builtin?
	//			glib.TYPE_BOOLEAN, // enabled?
	//			glib.TYPE_STRING,  // Status
	//			glib.TYPE_STRING,  // Checksum
	//			glib.TYPE_STRING,  // Last ReKey time
	//		)

	//		ret.ls_ava, _ = gtk.ListStoreNew(
	//			glib.TYPE_STRING,  // Name
	//			glib.TYPE_BOOLEAN, // builtin?
	//			glib.TYPE_STRING,  // Checksum
	//			glib.TYPE_STRING,  // Path
	//			glib.TYPE_STRING,  // Descr
	//		)

	//		ret.tw_acc.SetModel(ret.ls_acc)
	//		ret.tw_ava.SetModel(ret.ls_ava)
	//	}

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

			},
		)

		<-c

	}
}
