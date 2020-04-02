package go_modemmanager

import "github.com/godbus/dbus/v5"

// This interface provides access to perform different firmware-related operations in the modem,
// including listing the available firmware images in the module and selecting which of them to use.
// This interface does not provide direct access to perform firmware updates in the device. Instead, it
// exposes information about the expected firmware update method as well as method-specific details required for the
// upgrade to happen. The actual firmware upgrade may be performed via the Linux Vendor Firmware Service and the fwupd daemon.
// This interface will always be available as long a the modem is considered valid.

const (
	FirmwareInterface = ModemInterface + ".Firmware"

	/* Methods */

	/* Property */

)

type Firmware interface {
	/* METHODS */

	//MarshalJSON() ([]byte, error)
}

func NewFirmware(objectPath dbus.ObjectPath) (Firmware, error) {
	var fi firmware
	return &fi, fi.init(ModemManagerInterface, objectPath)
}

type firmware struct {
	dbusBase
}
