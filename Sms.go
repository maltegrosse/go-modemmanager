package go_modemmanager

import "github.com/godbus/dbus/v5"

// The SMS interface Defines operations and properties of a single SMS message.

const (
	SmsInterface = ModemManagerInterface + ".Sms"

	SMSsObjectPath = modemManagerMainObjectPath + "SMSs"
	/* Methods */

	/* Property */

)

type Sms interface {
	/* METHODS */

	//MarshalJSON() ([]byte, error)
}

func NewSms(objectPath dbus.ObjectPath) (Sms, error) {
	var ss sms
	return &ss, ss.init(SimInterface, objectPath)
}

type  sms struct {
	dbusBase
}
