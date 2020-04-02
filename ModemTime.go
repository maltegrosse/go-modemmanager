package go_modemmanager

import "github.com/godbus/dbus/v5"

// This interface allows clients to receive network time and timezone updates broadcast by mobile networks.
// This interface will only be available once the modem is ready to be registered in the cellular network.
// 3GPP devices will require a valid unlocked SIM card before any of the features in the interface can be used.

const (
	TimeInterface = ModemInterface + ".Time"

	/* Methods */

	/* Property */

)

type MmTime interface {
	/* METHODS */

	//MarshalJSON() ([]byte, error)
}

func NewMmTime(objectPath dbus.ObjectPath) (MmTime, error) {
	var ti mmTime
	return &ti, ti.init(ModemManagerInterface, objectPath)
}

type mmTime struct {
	dbusBase
}
