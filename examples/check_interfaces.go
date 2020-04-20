package main

import (
	"fmt"
	"github.com/maltegrosse/go-modemmanager"
	"log"
)

func main() {
	// switch from qmi to at:
	// rmmod qmi_wwan
	// systemctl restart ModemManager
	// switch back
	// modprobe qmi_wwan
	// systemctl restart ModemManager
	tmpByteSlice := []byte("TEST")
	fmt.Println(string(tmpByteSlice))
	mmgr, err := modemmanager.NewModemManager()
	if err != nil {
		log.Fatal(err.Error())
	}
	version, err := mmgr.GetVersion()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("ModemManager Version: ", version)

	tmpByteSlice, err = mmgr.MarshalJSON()
	if err != nil {
		fmt.Println(err)
		fmt.Println("ModemManager not available")
	} else {

		fmt.Println("ModemManager available")
	}

	err = mmgr.ScanDevices()
	if err != nil {
		log.Fatal(err.Error())
	}

	modems, err := mmgr.GetModems()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("found ", len(modems), " modem(s) ")
	for _, modem := range modems {
		fmt.Println("ObjectPath: ", modem.GetObjectPath())
		drivers, err := modem.GetDrivers()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Drivers: ", drivers)

		tmpByteSlice, err = modem.MarshalJSON()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Modem not available")
		} else {
			fmt.Println("Modem available")
		}

		bearers, err := modem.GetBearers()
		if err != nil {
			log.Fatal(err.Error())
		}
		for _, bearer := range bearers {
			fmt.Println("Found bearer at:", bearer.GetObjectPath())
			tmpByteSlice, err = bearer.MarshalJSON()
			if err != nil {
				fmt.Println(err)
				fmt.Println("Bearer not available")
			} else {
				fmt.Println("Bearer available")
			}
		}
		sim, err := modem.GetSim()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(" Found Sim: ", sim.GetObjectPath())
		tmpByteSlice, err = sim.MarshalJSON()
		if err != nil {
			fmt.Println(err)
			fmt.Println("SIM not available")
		} else {
			fmt.Println("SIM available")
		}

		modemSimple, err := modem.GetSimpleModem()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("ModemSimple for: ", modemSimple.GetObjectPath())
		simpleModemStatus, err := modemSimple.GetStatus()
		if err != nil {
			fmt.Println(simpleModemStatus)
			fmt.Println(err)
			fmt.Println("ModemSimple not available")
		} else {
			fmt.Println("ModemSimple available")
		}

		modem3gpp, err := modem.Get3gpp()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Modem3gpp for: ", modem3gpp.GetObjectPath())
		tmpByteSlice, err = modem3gpp.MarshalJSON()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Modem3GPP not available")
		} else {
			fmt.Println("Modem3GPP available")
		}

		ussd, err := modem3gpp.GetUssd()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Ussd for: ", ussd.GetObjectPath())
		tmpByteSlice, err = ussd.MarshalJSON()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Ussd not available")
		} else {
			fmt.Println("Ussd available")
		}

		mCdma, err := modem.GetCdma()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("ModemCdma for: ", mCdma.GetObjectPath())
		tmpByteSlice, err = mCdma.MarshalJSON()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Cdma not available")
		} else {
			fmt.Println("Cdma available")
		}

		messaging, err := modem.GetMessaging()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("ModemMessaging for: ", messaging.GetObjectPath())
		tmpByteSlice, err = messaging.MarshalJSON()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Messaging not available")
		} else {
			fmt.Println("Messaging available")
		}

		modemLocation, err := modem.GetLocation()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("ModemLocation for: ", modemLocation.GetObjectPath())
		tmpByteSlice, err = modemLocation.MarshalJSON()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Location not available")
		} else {
			fmt.Println("Location available")
		}

		modemTime, err := modem.GetTime()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("ModemTime for: ", modemTime.GetObjectPath())
		tmpByteSlice, err = modemTime.MarshalJSON()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Time not available")
		} else {
			fmt.Println("Time available")
		}

		voice, err := modem.GetVoice()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("ModemVoice for", voice.GetObjectPath())
		tmpByteSlice, err = voice.MarshalJSON()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Voice not available")
		} else {
			fmt.Println("Voice available")
		}

		modemFirmware, err := modem.GetFirmware()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("ModemFirmware for: ", modemFirmware.GetObjectPath())
		tmpByteSlice, err = modemFirmware.MarshalJSON()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Firmware not available")
		} else {
			fmt.Println("Firmware available")
		}

		modemSignal, err := modem.GetSignal()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Modem Signal for: ", modemSignal.GetObjectPath())
		tmpByteSlice, err = modemSignal.MarshalJSON()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Signal not available")
		} else {
			fmt.Println("Signal available")
		}

		modemOma, err := modem.GetOma()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("ModemOMA for: ", modemOma.GetObjectPath())
		tmpByteSlice, err = modemOma.MarshalJSON()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Oma not available")
		} else {
			fmt.Println("Oma available")
		}

	}
}
