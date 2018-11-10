package builtin_ownkeypair

import (
	"github.com/AnimusPEXUS/appplugsys"
	"github.com/jinzhu/gorm"
)

func GetPlugin() *appplugsys.Plugin {

	ret := &appplugsys.Plugin{
		Name:        "builtin_ownkeypair",
		Description: "manipulate own keypair",
		Requires:    []string{},
		Init: func(
			iface *appplugsys.AppPlugSysIface,
			db *gorm.DB,
		) error {
			if controller == nil {
				c, err := NewController(iface, db)
				if err != nil {
					return err
				}
				controller = c
			}

			return nil
		},
		Destroy: func() error {
			if controller != nil {
				controller.Destroy()
			}
			return nil
		},

		GetController: func() interface{} {
			return controller
		},
		Applications: map[string]*appplugsys.PluginApplication{},
	}
	ret.Applications["ownkeypair"] = &appplugsys.PluginApplication{
		Name:  "ownkeypair",
		Title: "Your Own Key Pair Manager",
		Display: func() error {
			controller.Display()
			return nil
		},
	}
	icon_png_bytes, err := uiIconPngBytes()
	if err != nil {
		panic("programming error")
	}
	ret.Applications["ownkeypair"].Icon = icon_png_bytes
	return ret
}

type Controller struct {
	iface *appplugsys.AppPlugSysIface
	db    *gorm.DB

	window *UIWindow
}

func NewController(
	iface *appplugsys.AppPlugSysIface,
	db *gorm.DB,
) (*Controller, error) {

	self := new(Controller)
	self.iface = iface
	self.db = db

	if t, err := UIWindowNew(self); err != nil {
		return nil, err
	} else {
		self.window = t
	}

	return self, nil
}

func (self *Controller) Destroy() error {
	return nil
}

func (self *Controller) Display() error {
	self.window.Show()
	return nil
}

var controller *Controller
