package main

import (
	"fmt"
	"github.com/maltegrosse/go-modemmanager"
	"log"
)

func main() {

	mmgr, err := go_modemmanager.NewModemManager()
	if err != nil {
		log.Fatal(err.Error())
	}
	version, err := mmgr.GetVersion()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("ModemManager Version: ", version)
	err = mmgr.ScanDevices()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = mmgr.SetLogging(go_modemmanager.MMLoggingLevelDebug)
	if err != nil {
		log.Fatal(err.Error())
	}

	modems, err := mmgr.ListDevices()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("found ", len(modems), " modem(s) ")
	for _, modem := range modems {
		fmt.Println("ObjectPath: ", modem.GetObjectPath())

		//err = modem.Enable()
		//if err != nil {
		//	log.Fatal(err.Error())
		//}

		sim, err := modem.GetSim()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Found Sim: ", sim.GetObjectPath())
		bearers, err := modem.GetBearers()
		if err != nil {
			log.Fatal(err.Error())
		}
		for _, bearer := range bearers {
			fmt.Println("Found bearer:", bearer.GetObjectPath())
		}
		supportedCapabilites, err := modem.GetSupportedCapabilities()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("SupportedCapabilities: ", supportedCapabilites)
		currentCapabilites, err := modem.GetCurrentCapabilities()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("CurrentCapabilities: ", currentCapabilites)

		maxBearers, err := modem.GetMaxBearers()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Max Bearers:", maxBearers)

		maxActiveBearers, err := modem.GetMaxActiveBearers()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Max Active Bearers:", maxActiveBearers)

		manu, err := modem.GetManufacturer()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Manufacturer: ", manu)

		model, err := modem.GetModel()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Model: ", model)

		rev, err := modem.GetRevision()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Revision: ", rev)

		cConf, err := modem.GetCarrierConfiguration()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Carrier Config: ", cConf)

		cConfRev, err := modem.GetCarrierConfigurationRevision()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Carrier Config Rev: ", cConfRev)

		hRev, err := modem.GetHardwareRevision()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Hardware Rev: ", hRev)

		deviceIdent, err := modem.GetDeviceIdentifier()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Device Ident: ", deviceIdent)

		dev, err := modem.GetDevice()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Device: ", dev)

		drivers, err := modem.GetDriver()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Drivers: ", drivers)

		plugin, err := modem.GetPlugin()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Plugin: ", plugin)

		pPort, err := modem.GetPrimaryPort()
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
		fmt.Println("UnlockRequired: ", unlockReq)

		capabilities, err := modem.GetCurrentCapabilities()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Current Capabilities: ", capabilities)

		unlockRetries, err := modem.GetUnlockRetries()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("UnlockRetries: ", unlockRetries)

		state, err := modem.GetState()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("State: ", state)

		fstate, err := modem.GetStateFailedReason()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("FailedState: ", fstate)

		tecs, err := modem.GetAccessTechnologies()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("AccessTechnologies: ", tecs)

		signalQuality, recent, err := modem.GetSignalQuality()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("SignalQuality: ", signalQuality, recent)

		numbers, err := modem.GetOwnNumbers()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Numbers: ", numbers)

		pState, err := modem.GetPowerState()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("PowerState: ", pState)

		sModes, err := modem.GetSupportedModes()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("SupportedModes: ", sModes)

		cModes, err := modem.GetCurrentModes()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("CurrentModes: ", cModes)

		sbands, err := modem.GetSupportedBands()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("SupportedBands: ", sbands)

		cbands, err := modem.GetCurrentBands()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("CurrentBands: ", cbands)

		ipFams, err := modem.GetSupportedIpFamilies()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("SupportedIpFamilies: ", ipFams)

		/*// listen to modem updates, e.g SignalQuality
		c := modem.Subscribe()
		for v := range c {
			fmt.Println(v)
		}*/

		/*tmpBearer := go_modemmanager.BearerProperty{APN:"test.apn.com"}
		newBearer, err := modem.CreateBearer(tmpBearer)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("New Bearer: ",newBearer)
		*/

		modemSimple, err := modem.GetSimpleModem()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("ModemSimple at: ", modemSimple.GetObjectPath())
		/*
			sProps := go_modemmanager.SimpleProperties{Apn:"test.apn.com"}
			newBearer,err := modemSimple.Connect(sProps)
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println("New Bearer: ",newBearer)
		*/
		status, err := modemSimple.GetStatus()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("SimpleStatus: ", status)
/*
		err = modemSimple.Disconnect(bearers[0])
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Disconnected Bearers: ")
*/
	}

}
