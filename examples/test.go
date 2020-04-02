package main

import (
	"fmt"
	 "github.com/maltegrosse/go-modemmanager"
	"log"
)

func main() {

	mmgr, err :=go_modemmanager.NewModemManager()
	if err != nil {
		log.Fatal(err.Error())
	}
	version, err := mmgr.GetVersion()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("MM Version: ",version)
	err = mmgr.ScanDevices()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = mmgr.SetLogging(go_modemmanager.MMLoggingLevelDebug)
	if err != nil {
		log.Fatal(err.Error())
	}

	modems,err := mmgr.ListDevices()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(modems)
	for _,modem := range modems {

		//err = modem.Enable()
		//if err != nil {
		//	log.Fatal(err.Error())
		//}

		sim, err := modem.GetSim()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(sim)
		bearers, err := modem.GetBearers()
		if err != nil {
			log.Fatal(err.Error())
		}
		for _,bearer := range bearers {
			fmt.Println(bearer)
		}
		supportedCapabilites, err :=modem.GetSupportedCapabilities()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(supportedCapabilites)
		currentCapabilites, err :=modem.GetCurrentCapabilities()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(currentCapabilites)

		maxBearers, err :=modem.GetMaxBearers()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Max Bearers:", maxBearers)

		maxActiveBearers, err :=modem.GetMaxActiveBearers()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Max Active Bearers:", maxActiveBearers)

		manu, err :=modem.GetManufacturer()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Manufacturer: ", manu)

		model, err :=modem.GetModel()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Model: ", model)

		rev, err :=modem.GetRevision()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Revision: ", rev)

		cConf, err :=modem.GetCarrierConfiguration()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Carrier Config: ", cConf)

		cConfRev, err :=modem.GetCarrierConfigurationRevision()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Carrier Config Rev: ", cConfRev)

		hRev, err :=modem.GetHardwareRevision()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Hardware Rev: ", hRev)

		deviceIdent, err :=modem.GetDeviceIdentifier()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Device Ident: ", deviceIdent)

		dev, err :=modem.GetDevice()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Device: ", dev)

		drivers, err :=modem.GetDriver()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Drivers: ", drivers)

		plugin, err :=modem.GetPlugin()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Plugin: ", plugin)

		pPort, err :=modem.GetPrimaryPort()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Primary Port: ", pPort)

		ports, err := modem.GetPorts()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Ports: ", ports)

		eIdent, err := modem.GetEquipmentIdentifier()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Equipment Ident: ", eIdent)

		unlockReq, err := modem.GetUnlockRequired()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Print("UnlockRequired: ",unlockReq)



	}


}
