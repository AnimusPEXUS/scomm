package main

import (
	"path"

	"github.com/AnimusPEXUS/appplugsys"
	"github.com/AnimusPEXUS/scomm/cmd/scommgtk/builtin_plugins/builtin_ownkeypair"
	"github.com/jinzhu/gorm"

	_ "github.com/xeodou/go-sqlcipher"
)

type Controller struct {
	db *gorm.DB

	storage_name string
	storage_key  string

	main_window *UIMainWindow

	plugin_searcher appplugsys.PluginSearcherI
	plugin_opener   appplugsys.PluginOpenerI
	plugin_acceptor appplugsys.PluginAcceptorI
	appplugsys      *appplugsys.AppPlugSys
}

func NewController(
	storage_name string,
	storage_key string,
) (*Controller, error) {
	self := new(Controller)
	self.main_window = UIMainWindowNew(self)

	db, err := OpenMainStorage(storage_name)
	if err != nil {
		return nil, err
	}

	p := "PRAGMA key = '" + storage_key + "';"
	err = db.Exec(p).Error
	if err != nil {
		return nil, err
	}

	self.db = db

	if t, err := NewPluginSearcher(
		[]*appplugsys.Plugin{
			builtin_ownkeypair.Plugin,
		},
		[]string{path.Join(SCOMM_CONFIG_DIR, "plugins")},
	); err != nil {
		return nil, err
	} else {
		self.plugin_searcher = t
	}

	if t, err := appplugsys.NewAppPlugSys(
		self.db,
		self.plugin_searcher,
		self.plugin_opener,
		self.plugin_acceptor,
		nil,
	); err != nil {
		return nil, err
	} else {
		self.appplugsys = t
	}

	return self, nil
}

func (self *Controller) ShowMainWindow() {
	self.main_window.Show()
}
