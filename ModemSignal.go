package go_modemmanager

import "github.com/godbus/dbus/v5"

// This interface provides access to extended signal quality information.
// This interface will only be available once the modem is ready to be registered in the cellular network.
// 3GPP devices will require a valid unlocked SIM card before any of the features in the interface can be used.

const (
	SignalInterface = ModemInterface + ".Signal"

	/* Methods */

	/* Property */

)

type Signal interface {
	/* METHODS */

	//MarshalJSON() ([]byte, error)
}

func NewSignal(objectPath dbus.ObjectPath) (Signal, error) {
	var si signal
	return &si, si.init(ModemManagerInterface, objectPath)
}

type signal struct {
	dbusBase
}
