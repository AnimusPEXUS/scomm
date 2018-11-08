package main

import (
	"os/user"
	"path"
)

var SCOMM_CONFIG_DIR string

func init() {

	if _t, err := user.Current(); err == nil {
		SCOMM_CONFIG_DIR = path.Join(_t.HomeDir, ".config", "SComm")
	} else {
		panic(err.Error())
	}

}
