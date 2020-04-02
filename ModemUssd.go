package go_modemmanager

import "github.com/godbus/dbus/v5"

// This interface provides access to actions based on the USSD protocol.
// This interface will only be available once the modem is ready to be registered in the
// cellular network. 3GPP devices will require a valid unlocked SIM card before any of the features
// in the interface can be used.

const (
	ModemModem3gppInterface = Modem3gppInterface + ".Ussd"

	/* Methods */

	/* Property */

)

type Ussd interface {
	/* METHODS */

	//MarshalJSON() ([]byte, error)
}

func NewUssd(objectPath dbus.ObjectPath) (Ussd, error) {
	var mu ussd
	return &mu, mu.init(ModemManagerInterface, objectPath)
}

type ussd struct {
	dbusBase
}
