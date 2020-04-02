package go_modemmanager

import "github.com/godbus/dbus/v5"

// The Simple interface allows controlling and querying the status of Modems.
// This interface will only be available once the modem is ready to be registered in the
// cellular network. 3GPP devices will require a valid unlocked SIM card before any of the
// features in the interface can be used.

const (
	ModemSimpleInterface = ModemInterface + ".Simple"

	/* Methods */

	/* Property */

)

type ModemSimple interface {
	/* METHODS */

	//MarshalJSON() ([]byte, error)
}

func NewModemSimple(objectPath dbus.ObjectPath) (ModemSimple, error) {
	var ms modemSimple
	return &ms, ms.init(ModemManagerInterface, objectPath)
}

type modemSimple struct {
	dbusBase
}
