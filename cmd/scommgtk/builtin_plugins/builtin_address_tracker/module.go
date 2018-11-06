package builtin_address_tracker

import "github.com/AnimusPEXUS/dnet"

type Module struct {
	name *dnet.ModuleName
}

func (self *Module) Name() *dnet.ModuleName {
	if self.name == nil {
		t, err := dnet.ModuleNameNew("builtin_address_tracker")
		if err != nil {
			panic("this shold not been happen")
		}
		self.name = t
	}

	return self.name
}

func (self *Module) Title() string {
	return "Central DNet Peer Address Tracker"
}

func (self *Module) Description() string {
	return "Stores relation between DNet Addresses and other network addresses"
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
	return NewInstance(self, com)
}
