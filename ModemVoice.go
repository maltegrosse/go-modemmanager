package go_modemmanager

import "github.com/godbus/dbus/v5"

// The Voice interface handles Calls.
// This interface will only be available once the modem is ready to be registered in the cellular network.
// 3GPP devices will require a valid unlocked SIM card before any of the features in the interface can be used.

const (
	VoiceInterface = ModemInterface + ".Voice"

	/* Methods */

	/* Property */

)

type Voice interface {
	/* METHODS */

	//MarshalJSON() ([]byte, error)
}

func NewVoice(objectPath dbus.ObjectPath) (Voice, error) {
	var vo voice
	return &vo, vo.init(ModemManagerInterface, objectPath)
}

type voice struct {
	dbusBase
}
