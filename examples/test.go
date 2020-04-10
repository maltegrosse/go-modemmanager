package main

import (
	"fmt"
	"github.com/maltegrosse/go-modemmanager"
	"log"
	"time"
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
		modem3gpp, err := modem.GetModem3gpp()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Modem3gpp: ", modem3gpp.GetObjectPath())

		/*		// takes around 1 min
				networks, err := modem3gpp.Scan()
				if err != nil {
					log.Fatal(err.Error())
				}
				fmt.Println("------- ")
				fmt.Println("Scanned Networks: ", networks)
				modem3gpp.RequestScan() //async, takes ~1min
				fmt.Println("-----")
				networkRes, err := modem3gpp.GetScanResults()
				if err != nil {
					log.Fatal(err.Error())
				}
				fmt.Println(networkRes)
				time.Sleep(1*time.Minute)
				fmt.Println("----- sleep 1 min ------")
				networkRes2, err := modem3gpp.GetScanResults()
				if err != nil {
					log.Fatal(err.Error())
				}
				fmt.Println(networkRes2)
		*/
		imei, err := modem3gpp.GetImei()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Imei: ", imei)

		regState, err := modem3gpp.GetRegistrationState()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("RegistrationState: ", regState)

		opCode, err := modem3gpp.GetOperatorCode()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("OperatorCode: ", opCode)

		mcc, err := modem3gpp.GetMcc()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Mcc: ", mcc)

		mnc, err := modem3gpp.GetMnc()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Mnc: ", mnc)

		opName, err := modem3gpp.GetOperatorName()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("OperatorName: ", opName)

		facLocks, err := modem3gpp.GetEnabledFacilityLocks()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("EnabledFacilityLocks: ", facLocks)

		epsMode, err := modem3gpp.GetEpsUeModeOperation()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("EpsUeModeOperation: ", epsMode)

		pco, err := modem3gpp.GetPco()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Pco: ", pco)

		// ussd untested as not available via qmi
		ussd, err := modem3gpp.GetUssd()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Ussd for: ", ussd.GetObjectPath())



		// cdma untested as not available via qmi
		mCdma, err := modem.GetModemCdma()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("ModemCdma for: ", mCdma.GetObjectPath())

		modemTime, err := modem.GetModemTime()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("ModemTime for: ", modemTime.GetObjectPath())

		modemNTime, err := modemTime.GetNetworkTime()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Current Network Time: ", modemNTime)

		modemNTimeZone, err := modemTime.GetNetworkTimezone()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Current Network Time Zone: ", modemNTimeZone)

		modemFirmware, err := modem.GetModemFirmware()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("ModemFirmware for: ", modemFirmware.GetObjectPath())

		// functionality untested as my modem returns empty results
		usedFirmware,err := modemFirmware.List()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("ModemFirmware: ", usedFirmware)

		updateSettings, err := modemFirmware.GetUpdateSettings()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("ModemFirmware UpdateSettings: ", updateSettings)

		modemSignal, err := modem.GetModemSignal()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Modem Signal for: ", modemSignal.GetObjectPath())

		err = modemSignal.Setup(1)
		if err != nil {
			log.Fatal(err.Error())
		}

		time.Sleep(2*time.Second)

		mSignalRate, err := modemSignal.GetRate()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Signal Ratea: ", mSignalRate)

		mSignalCdma, err := modemSignal.GetCdma()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Signal Cdma: ", mSignalCdma)

		mSignalEvdo, err := modemSignal.GetEvdo()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Signal Evdo: ", mSignalEvdo)

		mSignalGsm, err := modemSignal.GetGsm()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Signal Gsm: ", mSignalGsm)

		mSignalUmts, err := modemSignal.GetUmts()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Signal Umts: ", mSignalUmts)

		mSignalLte, err := modemSignal.GetLte()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Signal Lte: ", mSignalLte)
		// todo test by setting up rate




	}

}
