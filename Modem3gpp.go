package go_modemmanager

import "github.com/godbus/dbus/v5"

// This interface provides access to specific actions that may be performed in modems with 3GPP capabilities.
// This interface will only be available once the modem is ready to be registered in the cellular network.
// 3GPP devices will require a valid unlocked SIM card before any of the features in the interface can be used.

const (
	Modem3gppInterface = ModemInterface + ".Modem3gpp"

	/* Methods */

	/* Property */

)

type Modem3gpp interface {
	/* METHODS */

	//MarshalJSON() ([]byte, error)
}

func NewModem3gpp(objectPath dbus.ObjectPath) (Modem3gpp, error) {
	var m3gpp modem3gpp
	return &m3gpp, m3gpp.init(ModemManagerInterface, objectPath)
}

type  modem3gpp struct {
	dbusBase
}
