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

	}
}
func listenToModemPropertiesChanged(modem modemmanager.Modem) {
	c := modem.SubscribePropertiesChanged()
	for v := range c {
		fmt.Println(v)
		interfaceName, changedProperties, invalidatedProperties, err := modem.ParsePropertiesChanged(v)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(interfaceName, changedProperties, invalidatedProperties)
		}

	}
}
func listenToModemStateChanged(modem modemmanager.Modem) {
	c := modem.SubscribeStateChanged()
	for v := range c {
		oldState, newState, reason, err := modem.ParseStateChanged(v)
		if err == nil {
			fmt.Println(oldState, newState, reason)
		}

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
