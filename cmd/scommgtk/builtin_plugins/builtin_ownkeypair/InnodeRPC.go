package builtin_ownkeypair

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

func (self *NodeInternalRPC) GetOwnPrivKey(arg bool, res *string) error {
	if t, err := self.instance.GetOwnPrivKey(); err != nil {
		return err
	} else {
		*res = t
	}
	return nil
}
