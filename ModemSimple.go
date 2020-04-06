package go_modemmanager

import "github.com/godbus/dbus/v5"

// The Simple interface allows controlling and querying the status of Modems.
// This interface will only be available once the modem is ready to be registered in the
// cellular network. 3GPP devices will require a valid unlocked SIM card before any of the
// features in the interface can be used.

const (
	ModemSimpleInterface = ModemInterface + ".Simple"

	/* Methods */
	ModemSimpleConnect = ModemSimpleInterface + ".Connect"
	ModemSimpleDisconnect = ModemSimpleInterface + ".Disconnect"
	ModemSimpleGetStatus = ModemSimpleInterface + ".GetStatus"
	/* Property */

)

type ModemSimple interface {
	/* METHODS */

	// Returns object path
	GetObjectPath() dbus.ObjectPath

	// Do everything needed to connect the modem using the given properties.
	//This method will attempt to find a matching packet data bearer and activate it if necessary,
	//returning the bearer's IP details. If no matching bearer is found, a new bearer will be created and activated, but this operation may fail if no resources are available to complete this connection attempt (ie, if a conflicting bearer is already active).
	//This call may make a large number of changes to modem configuration based on properties passed in.
	//For example, given a PIN-locked, disabled GSM/UMTS modem, this call may unlock the SIM PIN, alter the
	//access technology preference, wait for network registration (or force registration to a specific provider),
	//create a new packet data bearer using the given "apn", and connect that bearer.
	Connect(properties SimpleProperties)(Bearer,error)

	// Disconnect an active packet data connection.
	Disconnect()(Bearer, error)

	// data bearer, while if "/" (ie, no object given) this method will disconnect all active packet data bearers.
	DisconnectAll()error

	// Get the general modem status.
	GetStatus()(SimpleStatus, error)
	MarshalJSON() ([]byte, error)
}
type SimpleProperties struct {
	Pin string `json:"pin"` // SIM-PIN unlock code, given as a string value (signature "s").
	OperatorId string `json:"operator-id"` // ETSI MCC-MNC of a network to force registration with, given as a string value (signature "s").
	Apn string `json:"apn"`	// For GSM/UMTS and LTE devices the APN to use, given as a string value (signature "s").
	IpType MMBearerIpFamily `json:"ip-type"` // For GSM/UMTS and LTE devices the IP addressing type to use, given as a MMBearerIpFamily value (signature "u").
	AllowedAuth MMBearerAllowedAuth `json:"allowed-auth"` // The authentication method to use, given as a MMBearerAllowedAuth value (signature "u"). Optional in 3GPP.
	User string `json:"user"` // User name (if any) required by the network, given as a string value (signature "s"). Optional in 3GPP.
	Password string `json:"password"` // Password (if any) required by the network, given as a string value (signature "s"). Optional in 3GPP.
	Number string `json:"number"` // For POTS devices the number to dial,, given as a string value (signature "s").
	AllowedRoaming bool `json:"allow-roaming"` // FALSE to allow only connections to home networks, given as a boolean value (signature "b").
	RmProtocol MMModemCdmaRmProtocol `json:"rm-protocol"` // For CDMA devices, the protocol of the Rm interface, given as a MMModemCdmaRmProtocol value (signature "u").

}

type SimpleStatus struct {

}

func NewModemSimple(objectPath dbus.ObjectPath) (ModemSimple, error) {
	var ms modemSimple
	return &ms, ms.init(ModemManagerInterface, objectPath)
}

type modemSimple struct {
	dbusBase
}



func (ms modemSimple) Connect(properties SimpleProperties) (Bearer, error) {
	panic("implement me")
}

func (ms modemSimple) Disconnect() (Bearer, error) {
	panic("implement me")
}

func (ms modemSimple) DisconnectAll() error {
	panic("implement me")
}

func (ms modemSimple) GetStatus() (SimpleStatus, error) {
	panic("implement me")
}

func (ms modemSimple) GetObjectPath() dbus.ObjectPath {
	return ms.obj.Path()
}

func (ms modemSimple) MarshalJSON() ([]byte, error) {

	panic("implement me")
}