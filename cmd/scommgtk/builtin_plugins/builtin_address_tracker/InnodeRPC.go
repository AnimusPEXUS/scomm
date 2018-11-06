package builtin_address_tracker

import (
	"time"
)

type NodeInternalRPC struct {
	instance         *Instance
	calling_app_name string
}

func NewNodeInternalRPC(
	instance *Instance,
	calling_app_name string,
) *NodeInternalRPC {
	ret := new(NodeInternalRPC)
	ret.instance = instance
	ret.calling_app_name = calling_app_name
	return ret
}

func (self *NodeInternalRPC) LastTimeWithNetNameAndNetAddr(
	arg *struct {
		NetName string
		NetAddr string
	},
	ret *time.Time,
) error {
	res, err := self.instance.db.LastTimeWithNetNameAndNetAddr(
		arg.NetName,
		arg.NetAddr,
	)
	ret = res
	return err
}

func (self *NodeInternalRPC) ApplyDNetAddressToNetAddrTime(
	arg *struct {
		NetName  string
		NetAddr  string
		Time     time.Time
		DNetAddr string
	},
	ret *bool,
) error {
	res, err := self.instance.db.ApplyDNetAddressToNetAddrTime(
		arg.NetName,
		arg.NetAddr,
		arg.Time,
		arg.DNetAddr,
	)
	*ret = res
	return err
}
