package go_modemmanager

import "github.com/godbus/dbus/v5"

// This interface allows clients to handle device management operations as specified by the Open Mobile Alliance (OMA).
// Device management sessions are either on-demand (client-initiated), or automatically initiated by either the device
// itself or the network.
// This interface will only be available once the modem is ready to be registered in the cellular network.
// 3GPP devices will require a valid unlocked SIM card before any of the features in the interface can be used.

const (
	OmaInterface = ModemInterface + ".Oma"

	/* Methods */

	/* Property */

)

type Oma interface {
	/* METHODS */

	//MarshalJSON() ([]byte, error)
}

func NewOma(objectPath dbus.ObjectPath) (Oma, error) {
	var om oma
	return &om, om.init(ModemManagerInterface, objectPath)
}

type oma struct {
	dbusBase
}
