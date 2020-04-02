package go_modemmanager

import "github.com/godbus/dbus/v5"

// The Messaging interface handles sending SMS messages and notification of new incoming messages.
// This interface will only be available once the modem is ready to be registered in the cellular network.
// 3GPP devices will require a valid unlocked SIM card before any of the features in the interface can
// be used (including listing stored messages).

const (
	MessagingInterface = ModemInterface + ".Messaging"

	/* Methods */

	/* Property */

)

type Messaging interface {
	/* METHODS */

	//MarshalJSON() ([]byte, error)
}

func NewMessaging(objectPath dbus.ObjectPath) (Messaging, error) {
	var me messaging
	return &me, me.init(ModemManagerInterface, objectPath)
}

type  messaging struct {
	dbusBase
}
