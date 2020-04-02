package go_modemmanager

import "github.com/godbus/dbus/v5"

// This interface provides access to specific actions that may be performed on available bearers.

const (
	BearerInterface = ModemManagerInterface + ".Bearer"

	BearersObjectPath = modemManagerMainObjectPath + "Bearers"
	/* Methods */

	/* Property */

)

type Bearer interface {
	/* METHODS */

	//MarshalJSON() ([]byte, error)
}

func NewBearer(objectPath dbus.ObjectPath) (Bearer, error) {
	var be bearer
	return &be, be.init(BearerInterface, objectPath)
}

type  bearer struct {
	dbusBase
}

type BearerProperty struct {
	APN string `json:"apn"` // Access Point Name, given as a string value (signature "s"). Required in 3GPP.
	IPType MMBearerIpFamily `json:"ip-type"` // Addressing type, given as a MMBearerIpFamily value (signature "u"). Optional in 3GPP and CDMA.
	AllowedAuth MMBearerAllowedAuth `json:"allowed-auth"` //  The authentication method to use, given as a MMBearerAllowedAuth value (signature "u"). Optional in 3GPP.
	User string `json:"string"` // User name (if any) required by the network, given as a string value (signature "s"). Optional in 3GPP.
	Password string `json:"password"` // Password (if any) required by the network, given as a string value (signature "s"). Optional in 3GPP.
	AllowRoaming bool `json:"allow-roaming"` // Flag to tell whether connection is allowed during roaming, given as a boolean value (signature "b"). Optional in 3GPP.
	RMProtocol MMModemCdmaRmProtocol `json:"rm-protocol"` // Protocol of the Rm interface, given as a MMModemCdmaRmProtocol value (signature "u"). Optional in CDMA.
	Number string `json:"number"` // Telephone number to dial, given as a string value (signature "s"). Required in POTS.
}
