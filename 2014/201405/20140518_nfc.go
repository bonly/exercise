package main

import (
	"fmt"
	"github.com/fuzxxl/nfc/dev/nfc"
	"log"
)

var m = nfc.Modulation{Type: nfc.ISO14443a, BaudRate: nfc.Nbr106}

func main() {
	fmt.Println("Using libnfc", nfc.Version())
	pnd, err := nfc.Open("")
	if err != nil {
		log.Fatalf("Could not open device: %v", err)
	}
	defer pnd.Close()

	if err := pnd.InitiatorInit(); err != nil {
		log.Fatalf("Could not init initiator: %v", err)
	}

	fmt.Println("Opened device", pnd)
	i, err := pnd.Information()
	if err != nil {
		log.Fatalf("Fetching information failed: %v", err)
	}
	fmt.Print(i)

	targets, err := pnd.InitiatorListPassiveTargets(m)

	if err != nil {
		log.Fatalf("Selecting target failed: %v", err)
	}

	for _, t := range targets {
		c, ok := t.(*nfc.ISO14443aTarget)
		if !ok {
			log.Println("Skipping", t)
			continue
		}
		fmt.Println("The following (NFC) ISO14443A tag was found:")
		fmt.Println("    ATQA (SENS_RES):", c.Atqa)
		//fmt.Println("       UID (NFCID%c): ", (nt.nti.nai.abtUid[0] == 0x08 ? '3' : '1'));
		fmt.Println("UID", c.UID)
		fmt.Println("SAK (SEL_RES)", c.Sak)

		if c.AtsLen > 0 {
			fmt.Println("ATS", c.Ats)
		}

	}

}