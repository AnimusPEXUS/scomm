package main

import "github.com/jinzhu/gorm"

type Controller struct {
	db *gorm.DB
}

func NewController(
	name string,
	key string,
) (*Controller, error) {
	self := new(Controller)
	return self, nil
}

func (self *Controller) ShowMainWindow() {

}
