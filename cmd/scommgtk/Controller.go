package main

import (
	"path"
	"time"

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
	appplugsys      *appplugsys.AppPlugSys
}

func NewController(
	storage_name string,
	storage_key string,
) (*Controller, error) {

	self := new(Controller)

	if t, err := UIMainWindowNew(self); err != nil {
		return nil, err
	} else {
		self.main_window = t
	}

	db, err := OpenMainStorage(storage_name)
	if err != nil {
		return nil, err
	}

	p := "PRAGMA key = '" + storage_key + "';" // TODO: escape storage_key
	err = db.Exec(p).Error
	if err != nil {
		return nil, err
	}

	self.db = db

	if t, err := NewPluginSearcher(
		[]*appplugsys.Plugin{
			builtin_ownkeypair.GetPlugin(),
		},
		[]string{path.Join(SCOMM_CONFIG_DIR, "plugins")},
	); err != nil {
		return nil, err
	} else {
		self.plugin_searcher = t
	}

	if t, err := appplugsys.NewAppPlugSys(
		self.db,
		func(info *appplugsys.DBPluginInfo) (*gorm.DB, bool, bool, error) {
			db, err := OpenApplicationStorage(
				storage_name,
				info.Name,
			)
			if err != nil {
				return nil, false, false, err
			}

			if info.Key != "" {
				p := "PRAGMA key = '" + info.Key + "';"
				err = db.Exec(p).Error
				if err != nil {
					return nil, false, false, err
				}
			}

			need_key := false
			need_rekey := false

			need_key = info.Key == ""

			if !need_key {
				tt := time.Now().UTC()
				dur := tt.Sub(info.LastDBReKey)

				if dur.Hours() > time.Duration(30*(24*time.Hour)).Hours() {
					need_rekey = true
				}
			}

			return db, need_key, need_rekey, nil

			//			if info.Key == "" || -
		},
		self.plugin_searcher,
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
