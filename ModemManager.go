package modemmanager

import (
	"fmt"
	"reflect"
)

// Paths of methods and properties
const (
	ModemManagerInterface = "org.freedesktop.ModemManager1"

	ModemManagerObjectPath     = "/org/freedesktop/ModemManager1"
	modemManagerMainObjectPath = "/org/freedesktop/ModemManager/"

	/* Methods */
	ModemManagerScanDevices       = ModemManagerInterface + ".ScanDevices"
	ModemManagerSetLogging        = ModemManagerInterface + ".SetLogging"
	ModemManagerReportKernelEvent = ModemManagerInterface + ".ReportKernelEvent"
	ModemManagerInhibitDevice     = ModemManagerInterface + ".InhibitDevice"

	/* Property */
	ModemManagerPropertyVersion = ModemManagerInterface + ".Version" // readable   s

)

// The ModemManager interface allows controlling and querying the status of the ModemManager daemon.
type ModemManager interface {
	/* METHODS */

	// Start a new scan for connected modem devices.
	ScanDevices() error

	// List modem devices. renamed from ListDevices to GetModems
	GetModems() ([]Modem, error)

	// Set logging verbosity.
	SetLogging(level MMLoggingLevel) error

	// Event Properties.
	// Reports a kernel event to ModemManager.
	// This method is only available if udev is not being used to report kernel events.
	// The properties dictionary is composed of key/value string pairs. The possible keys are:
	// see EventProperty and MMKernelPropertyAction
	ReportKernelEvent(EventProperties) error

	// org.freedesktop.ModemManager1.Modem:Device property. inhibit: TRUE to inhibit the modem and FALSE to uninhibit it.
	// Inhibit or uninhibit the device.
	// When the modem is inhibited ModemManager will close all its ports and unexport it from the bus, so that users of the interface are no longer able to operate with it.
	// This operation binds the inhibition request to the existence of the caller in the DBus bus. If the caller disappears from the bus, the inhibition will automatically removed.
	// 		IN s uid: the unique ID of the physical device, given in the
	// 		IN b inhibit:
	InhibitDevice(uid string, inhibit bool) error

	// The runtime version of the ModemManager daemon.
	GetVersion() (string, error)

	MarshalJSON() ([]byte, error)
}

// Returns new ModemManager Interface
func NewModemManager() (ModemManager, error) {
	var mm modemManager
	return &mm, mm.init(ModemManagerInterface, ModemManagerObjectPath)
}

type modemManager struct {
	dbusBase
}

// EventProperties  defines the properties which should be reported to the kernel
type EventProperties struct {
	Action    MMKernelPropertyAction `json:"action"`    // The type of action, given as a string value (signature "s"). This parameter is MANDATORY.
	Name      string                 `json:"name"`      // The device name, given as a string value (signature "s"). This parameter is MANDATORY.
	Subsystem string                 `json:"subsystem"` // The device subsystem, given as a string value (signature "s"). This parameter is MANDATORY.
	Uid       string                 `json:"uid"`       // The unique ID of the physical device, given as a string value (signature "s"). This parameter is OPTIONAL, if not given the sysfs path of the physical device will be used. This parameter must be the same for all devices exposed by the same physical device.
}

func (mm modemManager) GetModems() (modems []Modem, err error) {
	fmt.Println("####->", ModemManagerInterface, ModemManagerObjectPath)
	devPaths, err := mm.getManagedObjects(ModemManagerInterface, ModemManagerObjectPath)
	if err != nil {
		return nil, err
	}
	fmt.Println(devPaths)
	for idx := range devPaths {
		modem, err := NewModem(devPaths[idx])
		if err != nil {
			return nil, err
		}
		modems = append(modems, modem)
	}
	return
}

func (mm modemManager) ScanDevices() error {
	err := mm.call(ModemManagerScanDevices)
	return err
}

func (mm modemManager) SetLogging(level MMLoggingLevel) error {
	err := mm.call(ModemManagerSetLogging, &level)
	return err
}

func (mm modemManager) ReportKernelEvent(properties EventProperties) error {
	// untested
	v := reflect.ValueOf(properties)
	st := reflect.TypeOf(properties)
	type dynMap interface{}
	var myMap map[string]dynMap
	myMap = make(map[string]dynMap)
	for i := 0; i < v.NumField(); i++ {
		field := st.Field(i)
		tag := field.Tag.Get("json")
		value := v.Field(i).Interface()
		if v.Field(i).IsZero() {
			continue
		}
		myMap[tag] = value
	}
	return mm.call(ModemManagerReportKernelEvent, &myMap)
}

func (mm modemManager) InhibitDevice(uid string, inhibit bool) error {
	// untested
	err := mm.call(ModemManagerInhibitDevice, &uid, &inhibit)
	return err
}

func (mm modemManager) GetVersion() (string, error) {
	v, err := mm.getStringProperty(ModemManagerPropertyVersion)
	return v, err
}

func (mm modemManager) MarshalJSON() ([]byte, error) {
	// todo: not implemented yet
	panic("implement me")
}
