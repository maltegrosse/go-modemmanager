package go_modemmanager

import "github.com/godbus/dbus/v5"

// The Location interface allows devices to provide location information to client applications.
// Not all devices can provide this information, or even if they do, they may not be able to provide it while
// a data session is active.
// This interface will only be available once the modem is ready to be registered in the cellular network.
// 3GPP devices will require a valid unlocked SIM card before any of the features in the interface can be used
// (including GNSS module management).

const (
	LocationInterface = ModemInterface + ".Location"

	/* Methods */

	/* Property */

)

type Location interface {
	/* METHODS */

	//MarshalJSON() ([]byte, error)
}

func NewLocation(objectPath dbus.ObjectPath) (Location, error) {
	var lo location
	return &lo, lo.init(ModemManagerInterface, objectPath)
}

type  location struct {
	dbusBase
}
