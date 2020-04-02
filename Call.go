package go_modemmanager

import "github.com/godbus/dbus/v5"

// The Call interface Defines operations and properties of a single Call.

const (
	CallInterface = ModemManagerInterface + ".Call"

	CallsObjectPath = modemManagerMainObjectPath + "Calls"
	/* Methods */

	/* Property */

)

type Call interface {
	/* METHODS */

	//MarshalJSON() ([]byte, error)
}

func NewCall(objectPath dbus.ObjectPath) (Call, error) {
	var ca call
	return &ca, ca.init(CallInterface, objectPath)
}

type  call struct {
	dbusBase
}
