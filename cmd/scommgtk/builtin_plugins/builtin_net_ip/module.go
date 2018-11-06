package builtin_net_ip

import "github.com/AnimusPEXUS/dnet"

var (
	MULTICAST_IP = "224.0.0.1"
)

type Module struct {
	name *dnet.ModuleName
}

func (self *Module) Name() *dnet.ModuleName {
	if self.name == nil {
		t, err := dnet.ModuleNameNew("builtin_net_ip")
		if err != nil {
			panic("this shold not been happen")
		}
		self.name = t
	}

	return self.name
}

func (self *Module) Title() string {
	return "net_ip"
}

func (self *Module) Description() string {
	return ("IP transport mechanism")
}

func (self *Module) DependsOn() []string {
	return []string{}
}

func (self *Module) IsWorker() bool {
	return true
}

func (self *Module) IsNetwork() bool {
	return true
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
