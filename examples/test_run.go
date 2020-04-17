package main

import (
	"fmt"
	"github.com/maltegrosse/go-modemmanager"
	"log"
	"time"
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
	err = mmgr.ScanDevices()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = mmgr.SetLogging(modemmanager.MMLoggingLevelDebug)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("### Modems Start ###")
	modems, err := mmgr.GetModems()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("found ", len(modems), " modem(s) ")
	for _, modem := range modems {
		fmt.Println("ObjectPath: ", modem.GetObjectPath())

		// test gps
		gpsAvailable := false

		//err = modem.Enable()
		//if err != nil {
		//	log.Fatal(err.Error())
		//}
		fmt.Println("### START Sim ####")
		sim, err := modem.GetSim()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(" Found Sim: ", sim.GetObjectPath())

		simIdent, err := sim.GetSimIdentifier()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("-  SimIdentifier", simIdent)

		simImsi, err := sim.GetImsi()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("-  Imsi", simImsi)

		simOpIdent, err := sim.GetOperatorIdentifier()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("-  OperatorIdentifier", simOpIdent)

		simOp, err := sim.GetOperatorName()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("-  OperatorName", simOp)

		simEm, err := sim.GetEmergencyNumbers()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("-  EmergencyNumbers", simEm)

		fmt.Println("### END Sim ###")

		bearers, err := modem.GetBearers()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("### START Bearer ####")
		for _, bearer := range bearers {
			fmt.Println("Found bearer at:", bearer.GetObjectPath())
			bearerInterface, err := bearer.GetInterface()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println("- Interface:", bearerInterface)

			bearerConnected, err := bearer.GetConnected()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println("- Connected:", bearerConnected)

			bearerSuspended, err := bearer.GetSuspended()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println("- Suspended:", bearerSuspended)

			bearerip4, err := bearer.GetIp4Config()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println("- Ip4Config:", bearerip4)

			bearerip6, err := bearer.GetIp6Config()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println("- Ip6Config:", bearerip6)

			bearerStats, err := bearer.GetStats()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println("- Stats:", bearerStats)

			beareripTimeout, err := bearer.GetIpTimeout()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println("- IpTimeout:", beareripTimeout)

			bearerType, err := bearer.GetBearerType()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println("- BearerType:", bearerType)

			bearerProperties, err := bearer.GetProperties()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println("- Properties:", bearerProperties)

		}

		fmt.Println("### END Bearer ###")

		fmt.Println("### START Modem ####")
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
		fmt.Println("### START Simple Modem ####")
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
		fmt.Println("### END Simple Modem ####")

		fmt.Println("### START 3GPP ####")
		modem3gpp, err := modem.Get3gpp()
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

		fmt.Println("### END 3GPP ####")
		// cdma untested as not available via qmi

		mCdma, err := modem.GetCdma()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("ModemCdma for: ", mCdma.GetObjectPath())

		fmt.Println("### START Time ####")
		modemTime, err := modem.GetTime()
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
		fmt.Println("### END Time ####")

		fmt.Println("### START Firmware ####")
		modemFirmware, err := modem.GetFirmware()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("ModemFirmware for: ", modemFirmware.GetObjectPath())

		// functionality untested as my modem returns empty results
		usedFirmware, err := modemFirmware.List()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("ModemFirmware: ", usedFirmware)

		updateSettings, err := modemFirmware.GetUpdateSettings()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("ModemFirmware UpdateSettings: ", updateSettings)
		fmt.Println("### END Firmware ####")

		fmt.Println("### START Signal ####")
		modemSignal, err := modem.GetSignal()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Modem Signal for: ", modemSignal.GetObjectPath())

		err = modemSignal.Setup(1)
		if err != nil {
			log.Fatal(err.Error())
		}
		//	fmt.Println("Set Signal rate to 1 sec, now wait 2 seconds....")
		//	time.Sleep(2*time.Second)

		mSignalRate, err := modemSignal.GetRate()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Signal Rate: ", mSignalRate)

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

		currentSignal, err := modemSignal.GetCurrentSignals()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Current Signal: ", currentSignal)

		err = modemSignal.Setup(0)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Set Signal rate to 0 sec, disabled...")
		fmt.Println("### END Signal ####")
		// OMA untested as not available via qmi
		modemOma, err := modem.GetOma()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("ModemOMA at: ", modemOma.GetObjectPath())

		fmt.Println("### START Location ####")
		modemLocation, err := modem.GetLocation()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("ModemLocation at: ", modemOma.GetObjectPath())

		mlCap, err := modemLocation.GetCapabilities()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Capabilities: ", mlCap)

		var tmpLocCap []modemmanager.MMModemLocationSource
		tmpLocCap = append(tmpLocCap, modemmanager.MmModemLocationSource3gppLacCi)
		err = modemLocation.Setup(tmpLocCap, true)
		if err != nil {
			log.Fatal(err.Error())
		}

		mloc, err := modemLocation.GetCurrentLocation()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Location 3gppLacCi: ", mloc.ThreeGppLacCi)

		mlocEnabledSources, err := modemLocation.GetEnabledLocationSources()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Location Enabled Sources: ", mlocEnabledSources)

		if gpsAvailable {
			//tmpLocCap = append(tmpLocCap, go_modemmanager.MmModemLocationSourceGpsRaw)
			//err = modemLocation.Setup(tmpLocCap,true)
			//if err != nil {
			//	log.Fatal(err.Error())
			//}
			//fmt.Println("Wait two seconds until gps signal is ready")
			//time.Sleep(2*time.Second)
			//mloc, err = modemLocation.GetCurrentLocation()
			//if err != nil {
			//	log.Fatal(err.Error())
			//}
			//fmt.Println("Location GpsRaw: ", mloc.GpsRaw)

			tmpLocCap = append(tmpLocCap, modemmanager.MmModemLocationSourceGpsNmea)
			err = modemLocation.Setup(tmpLocCap, true)
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println("Wait four seconds until gps signal is ready")
			time.Sleep(4 * time.Second)

			mlocEnabledSources, err := modemLocation.GetEnabledLocationSources()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println("Location Enabled Sources: ", mlocEnabledSources)

			err = modemLocation.SetGpsRefreshRate(1)
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println("Location Set Refresh rate to 1 sec: ")

			fmt.Println("Wait four seconds until gps signal is ready")
			time.Sleep(4 * time.Second)
			mloc, err = modemLocation.GetCurrentLocation()
			if err != nil {
				log.Fatal(err.Error())
			}
			// use e.g. github.com/adrianmo/go-nmea to parse
			fmt.Println("Location GpsNmea: ", mloc.GpsNmea)

			mlocsdata, err := modemLocation.GetSupportedAssistanceData()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println("Location SupportedAssistanceData: ", mlocsdata)

			mlocssignal, err := modemLocation.GetSignalsLocation()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println("Location SignalsLocation: ", mlocssignal)

			mloc2, err := modemLocation.GetLocation()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println("Location Property: ", mloc2)

			mlocSuplS, err := modemLocation.GetSuplServer()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println("Location SuplServer: ", mlocSuplS)

			mlocAssDs, err := modemLocation.GetAssistanceDataServers()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println("Location AssistanceDataServers: ", mlocAssDs)

			mlocrefreshr, err := modemLocation.GetGpsRefreshRate()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println("Location GpsRefreshRate: ", mlocrefreshr)

		}
		fmt.Println("### END Location ####")

		fmt.Println("### START Messaging ####")
		// sms
		messaging, err := modem.GetMessaging()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("ModemMessaging at: ", messaging.GetObjectPath())

		smss, err := messaging.List()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Found ", len(smss), " SMS")

		// See comments regarding method and property
		//smss2, err := messaging.GetMessages()
		//if err != nil {
		//	log.Fatal(err.Error())
		//}
		//fmt.Println("Found ", len(smss2), " SMS")

		messSS, err := messaging.GetSupportedStorages()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("SMS Supported Storages ", messSS)

		messDS, err := messaging.GetDefaultStorage()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("SMS Default Storage ", messDS)

		for _, sms := range smss {
			fmt.Println(" SMS at ", sms.GetObjectPath())

			smsState, err := sms.GetState()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(" - State: ", smsState)

			smsPdu, err := sms.GetPduType()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(" - PduType: ", smsPdu)

			smsNo, err := sms.GetNumber()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(" - Number: ", smsNo)

			smsTxt, err := sms.GetText()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(" - Text: ", smsTxt)

			smsData, err := sms.GetData()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(" - Data: ", smsData)

			smsSmsc, err := sms.GetSMSC()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(" - SMSC: ", smsSmsc)

			smsVal, err := sms.GetValidity()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(" - Validity: ", smsVal)

			smsClass, err := sms.GetClass()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(" - Class: ", smsClass)

			smsTid, err := sms.GetTeleserviceId()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(" - TeleserviceId: ", smsTid)

			smsSc, err := sms.GetServiceCategory()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(" - ServiceCategory: ", smsSc)

			smsDRR, err := sms.GetDeliveryReportRequest()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(" - DeliveryReportRequest: ", smsDRR)

			smsMR, err := sms.GetMessageReference()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(" - MessageReference: ", smsMR)

			smsTS, err := sms.GetTimestamp()
			if err != nil {
				fmt.Println(" - Timestamp: ", err.Error())

			} else {
				fmt.Println(" - Timestamp: ", smsTS)
			}

			smsDTS, err := sms.GetDischargeTimestamp()
			if err != nil {
				fmt.Println(" - DischargeTimestamp: ", err.Error())
			} else {
				fmt.Println(" - DischargeTimestamp: ", smsDTS)
			}

			smsDS, err := sms.GetDeliveryState()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(" - DeliveryState: ", smsDS)

			smsStor, err := sms.GetStorage()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(" - Storage: ", smsStor)
		}
		//err = messaging.Delete(smss[1])
		//if err != nil {
		//	log.Fatal(err.Error())
		//}
		//fmt.Println("Deleted SMS at position 1")

		// create and send sms
		//singleSms, err := messaging.CreateSms("0172xxxx","greetings from modem manager")
		//if err != nil {
		//	log.Fatal(err.Error())
		//} else {
		//	fmt.Println(" - New sms at: ", singleSms.GetObjectPath())
		//	err = singleSms.Send()
		//	if err != nil {
		//		log.Fatal(err.Error())
		//	}
		//	fmt.Println(" - SMS sent: ", singleSms.GetObjectPath())
		//}

		fmt.Println("### END Messaging ####")

		enableVoiceTest := false
		if enableVoiceTest {
			fmt.Println("### START Voice ####")
			voice, err := modem.GetVoice()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(" voice at", voice.GetObjectPath())

			call, err := voice.CreateCall("0173xxxx")
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(" -starting call", call.GetObjectPath())
			err = call.Start()
			if err != nil {
				log.Fatal(err.Error())
			}
			time.Sleep(3 * time.Second)
			callState, err := call.GetState()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(" -State", callState)

			time.Sleep(3 * time.Second)
			callStateReason, err := call.GetStateReason()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(" -StateReason", callStateReason)

			time.Sleep(3 * time.Second)
			callDirection, err := call.GetDirection()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(" -Direction", callDirection)

			time.Sleep(3 * time.Second)
			callMulti, err := call.GetMultiparty()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(" -Multiparty", callMulti)

			time.Sleep(5 * time.Second)
			err = call.SendDtmf("123454567")
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(" -sent dtmf")

			time.Sleep(5 * time.Second)
			callAudioPort, err := call.GetAudioPort()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(" -AudioPort", callAudioPort)

			time.Sleep(3 * time.Second)
			callFormat, err := call.GetAudioFormat()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(" -AudioFormat", callFormat)

			time.Sleep(10 * time.Second)

			// only works if accepted
			err = call.Hangup()
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(" -hangup call", call.GetObjectPath())

			err = voice.HangupAll()
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(" -hangup all", call.GetObjectPath())

			fmt.Println("### END Voice ####")
		}

	}

}
