package main

import (
	"fmt"
	"github.com/maltegrosse/go-modemmanager"
	"log"
	"reflect"
)

func main() {
	mmgr, err := modemmanager.NewModemManager()
	if err != nil {
		log.Fatal(err.Error())
	}
	version, err := mmgr.GetVersion()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("ModemManager Version: ", version)
	modems, err := mmgr.GetModems()
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, modem := range modems {
		fmt.Println("ObjectPath: ", modem.GetObjectPath())
		listenToModemPropertiesChanged(modem)
		// todo: add PropertiesChanged All objects (Manager, Modems, Bearers, SIMs, SMSs) exported at the org.freedesktop.ModemManager1 bus name implement the standard org.freedesktop.DBus.Properties interface. Objects implementing this interface provide a common way to query for property values and also a generic signal to get notified about changes in those properties.

	}
}
func listenToModemPropertiesChanged(modem modemmanager.Modem) {
	c := modem.SubscribePropertiesChanged()
	for v := range c {
		fmt.Println(v)
		fmt.Println(reflect.TypeOf(v))
		fmt.Println("name", v.Name)
		fmt.Println("path", v.Path)
		fmt.Println("body", v.Body)
		// usually a map with map[property]interface
		fmt.Println("body", reflect.TypeOf(v.Body))
		fmt.Println("sender", v.Sender)
		fmt.Println("-----")
		for _, val := range v.Body {
			// 0 = string interface name
			// 1 = property changed as map[string]dbus.Variant e.g. map[State:11], -> hint parse dbus.Variant with .Value
			// 2 = slice of invalidated properties
			fmt.Println("val",reflect.TypeOf(val), val)
		}
	}
}
func listenToModemStateChanged(modem modemmanager.Modem) {
	c := modem.SubscribeStateChanged()
	for v := range c {
		oldState, newState, reason := modem.ParseStateChanged(v)
		fmt.Println(oldState, newState, reason)
	}

}

func listenToModemVoiceCallAdded(modem modemmanager.Modem) {
	// listen new calls
	voice, err := modem.GetVoice()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(voice.GetObjectPath())
	c := voice.SubscribeCallAdded()
	fmt.Println("start listening ....")
	for v := range c {
		fmt.Println(v)
		fmt.Println(reflect.TypeOf(v))
		fmt.Println("name", v.Name)
		fmt.Println("path", v.Path)
		fmt.Println("body", v.Body)
		fmt.Println("sender", v.Sender)
	}
}
