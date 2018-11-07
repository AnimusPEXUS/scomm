package main

import (
	"time"

	"github.com/AnimusPEXUS/utils/worker"
)

type UISysTrayIcon struct {
	w *worker.Worker
}

func NewUISysTrayIcon() *UISysTrayIcon {
	self := new(UISysTrayIcon)
	self.w = worker.New(self.workerFunction)
	return self
}

func (self *UISysTrayIcon) workerFunction(
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
	}
	return
}
