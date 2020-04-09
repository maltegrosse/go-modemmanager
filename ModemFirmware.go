package go_modemmanager

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"reflect"
)

// This interface provides access to perform different firmware-related operations in the modem,
// including listing the available firmware images in the module and selecting which of them to use.
// This interface does not provide direct access to perform firmware updates in the device. Instead, it
// exposes information about the expected firmware update method as well as method-specific details required for the
// upgrade to happen. The actual firmware upgrade may be performed via the Linux Vendor Firmware Service and the fwupd daemon.
// This interface will always be available as long a the modem is considered valid.

const (
	ModemFirmwareInterface = ModemInterface + ".Firmware"

	/* Methods */
	ModemFirmwareList = ModemFirmwareInterface + ".List"
	ModemFirmwareSelect = ModemFirmwareInterface + ".Select"

	/* Property */
	ModemFirmwarePropertyUpdateSettings = ModemFirmwareInterface + ".UpdateSettings" // readable   (ua{sv})


)

type ModemFirmware interface {
	/* METHODS */

	// Returns object path
	GetObjectPath() dbus.ObjectPath

	// List installed firmware images.
	// Firmware slots and firmware images are identified by arbitrary opaque strings.
	// 		List (OUT s      selected, OUT aa{sv} installed);

	List()(string, []FirmwareProperty, error)

	// Selects a different firmware image to use, and immediately resets the modem so that it begins using the new firmware image.
	// The method will fail if the identifier does not match any of the names returned by List(), or if the image could not be selected for some reason.
	// Installed images can be selected non-destructively.
	// 		IN s uniqueid: The unique ID of the firmware image to select.
	Select(string) error

	MarshalJSON() ([]byte, error)

	/* PROPERTIES */

	GetUpdateSettings()(interface{}, error)
}

func NewModemFirmware(objectPath dbus.ObjectPath) (ModemFirmware, error) {
	var fi modemFirmware
	return &fi, fi.init(ModemManagerInterface, objectPath)
}

type modemFirmware struct {
	dbusBase
}

type FirmwareProperty struct {
	ImageType MMFirmwareImageType `json:"image-type"`    // (Required) Type of the firmware image, given as a MMFirmwareImageType value (signature "u"). Firmware images of type MM_FIRMWARE_IMAGE_TYPE_GENERIC will only expose only the mandatory properties.
	UniqueId string `json:"unique-id"`    // (Required) A user-readable unique ID for the firmware image, given as a string value (signature "s").
	GobiPriVersion string `json:"gobi-pri-version"`    // (Optional) The version of the PRI firmware image, in images of type MM_FIRMWARE_IMAGE_TYPE_GOBI, given as a string value (signature "s").
	GobiPriInfo string `json:"gobi-pri-info"`    // (Optional) Additional information of the PRI image, in images of type MM_FIRMWARE_IMAGE_TYPE_GOBI, given as a string value (signature "s").
	GobiBootVersion string `json:"gobi-boot-version"`    // (Optional) The boot version of the PRI firmware image, in images of type MM_FIRMWARE_IMAGE_TYPE_GOBI, given as a string value (signature "s").
	GobiPriUniqueId string `json:"gobi-pri-unique-id"`    // (Optional) The unique ID of the PRI firmware image, in images of type MM_FIRMWARE_IMAGE_TYPE_GOBI, given as a string value (signature "s").
	GobiModemUniqueId string `json:"gobi-modem-unique-id"`    // (Optional) The unique ID of the Modem firmware image, in images of type MM_FIRMWARE_IMAGE_TYPE_GOBI, given as a string value (signature "s").

}
func (fp FirmwareProperty) String() string {
	return "ImageType: " + fmt.Sprint(fp.ImageType) +
		", UniqueId: " + fp.UniqueId +
		", GobiPriVersion: " + fp.GobiPriVersion +
		", GobiPriInfo: " + fp.GobiPriInfo +
		", GobiBootVersion: " + fp.GobiBootVersion +
		", GobiPriUniqueId: " + fp.GobiPriUniqueId +
		", GobiModemUniqueId: " + fp.GobiModemUniqueId
}

func (fi modemFirmware) GetObjectPath() dbus.ObjectPath {
	return fi.obj.Path()
}

func (fi modemFirmware) Select(uid string) error {
	return fi.call(ModemFirmwareSelect, uid)
}

func (fi modemFirmware) List() (string, []FirmwareProperty, error) {
	//todo double check if call is correct, empty result
	var myMap []map[string]dbus.Variant
	var tmpString string

	err := fi.callWithReturn2(&tmpString,&myMap,ModemFirmwareList)
	if err != nil {
		return "",nil, err
	}
	fmt.Println(tmpString)
	fmt.Println(myMap)
	return "",nil, nil
}





func (fi modemFirmware) GetUpdateSettings() (interface{}, error) {
	//todo finish
	res, err := fi.getInterfaceProperty(ModemFirmwarePropertyUpdateSettings)
	if err != nil {
		return nil, err
	}
	values, ok := res.([]interface{})
	if !ok {
		fmt.Println("wrong format")
	}
	for idx, v := range values {
		fmt.Println("idx",idx)
		fmt.Println(v)
		fmt.Println(reflect.TypeOf(v))
		fmt.Println("-----")
	}
	fmt.Println(reflect.TypeOf(res))
	return res, nil

}

func (fi modemFirmware) MarshalJSON() ([]byte, error) {
	panic("implement me")
}

