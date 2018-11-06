package builtin_owntlscert

import (
	"fmt"
	"sync"

	"github.com/AnimusPEXUS/dnet"
	"github.com/AnimusPEXUS/utils/worker"
)

type Module struct {
	name *dnet.ModuleName
}

func (self *Module) Name() *dnet.ModuleName {
	if self.name == nil {
		t, err := dnet.ModuleNameNew("builtin_owntlscert")
		if err != nil {
			panic("this shold not been happen")
		}
		self.name = t
	}

	return self.name
}

func (self *Module) Title() string {
	return "Your Own TLS Certificate Editor"
}

func (self *Module) Description() string {
	return "Use this to create or load existing TLS Certificate" +
		" to use with untrusted networks"
}

func (self *Module) DependsOn() []string {
	return []string{}
}

func (self *Module) IsWorker() bool {
	return false
}

func (self *Module) IsNetwork() bool {
	return false
}

func (self *Module) HaveUI() bool {
	return true
}

func (self *Module) Instantiate(com dnet.ApplicationCommunicatorI) (
	dnet.ApplicationModuleInstanceI,
	error,
) {
	ret := &Instance{}
	ret.com = com
	ret.db = &DB{db: com.GetDBConnection()}
	ret.mod = self
	ret.window_show_sync = new(sync.Mutex)

	if !ret.db.db.HasTable(&OwnData{}) {
		if err := ret.db.db.CreateTable(&OwnData{}).Error; err != nil {
			fmt.Println("builtin_owntlscert:", "Can't create table:", err.Error())
		}
	}

	ret.Worker = worker.New(ret.threadWorker)

	return ret, nil
}
