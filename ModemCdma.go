package go_modemmanager

import "github.com/godbus/dbus/v5"

// This interface provides access to specific actions that may be performed in modems with CDMA capabilities.
// This interface will only be available once the modem is ready to be registered in the cellular network.
// Mixed 3GPP+3GPP2 devices will require a valid unlocked SIM card before any of the features in the interface can be used.

const (
	ModemCdmaInterface = ModemInterface + ".ModemCdma"

	/* Methods */

	/* Property */

)

type ModemCdma interface {
	/* METHODS */

	//MarshalJSON() ([]byte, error)
}

func NewModemCdma(objectPath dbus.ObjectPath) (ModemCdma, error) {
	var mc modemCdma
	return &mc, mc.init(ModemManagerInterface, objectPath)
}

type  modemCdma struct {
	dbusBase
}
