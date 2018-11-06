package builtin_address_tracker

import (
	"sync"
	"time"

	"github.com/jinzhu/gorm"
)

type DB struct {
	db *gorm.DB

	s sync.Mutex
}

type AddressHistory struct {
	Key            uint64 `gorm:"primary_key"`
	DnetAddrStr    string
	DiscoveredTime *time.Time
	NetworkName    string
	NetworkAddress string
}

func (self *DB) LastTimeWithNetNameAndNetAddr(
	net_name string,
	net_addr string,
) (*time.Time, error) {
	self.s.Lock()
	defer self.s.Unlock()

	var res AddressHistory
	var ret *time.Time
	err := self.db.
		Order("discovered_time desc").
		Where(
			"network_name = ? AND network_address = ?",
			net_name,
			net_addr,
		).
		Limit(1).
		Find(&res).
		Error
	if err != nil {
	} else {
		t := res.DiscoveredTime
		ret = t
	}
	return ret, err
}

func (self *DB) ApplyDNetAddressToNetAddrTime(
	net_name string,
	net_addr string,
	discovered_time time.Time,
	dnet_addr string,
) (bool, error) {
	self.s.Lock()
	defer self.s.Unlock()

	var res AddressHistory
	err := self.db.
		Order("discovered_time desc").
		Where(
			"network_name = ? AND network_address = ? AND discovered_time =?",
			net_name,
			net_addr,
			discovered_time,
		).
		Limit(1).
		Find(&res).
		Error
	if err != nil {
		return false, err
	} else {
		res.DiscoveredTime = &discovered_time
		self.db.Update(res)
	}
	return true, nil
}
