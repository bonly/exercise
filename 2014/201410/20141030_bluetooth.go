package main

import (
	"log"
	"github.com/paypal/gatt"
)

var done = make(chan struct{})

var uartServiceId = gatt.MustParseUUID("6e400001-b5a3-f393-e0a9-e50e24dcca9e")
var uartServiceRXCharId = gatt.MustParseUUID("6e400002-b5a3-f393-e0a9-e50e24dcca9e")
var uartServiceTXCharId = gatt.MustParseUUID("6e400003-b5a3-f393-e0a9-e50e24dcca9e")

func onStateChanged(d gatt.Device, s gatt.State) {
	log.Println("State:", s)
	switch s {
	case gatt.StatePoweredOn:
		log.Println("scanning...")
		d.Scan([]gatt.UUID{}, false)
		return
	default:
		d.StopScanning()
	}
}

func onPeriphDiscovered(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {
	if (a.LocalName == "UART") {
		log.Printf("Preipheral Discovered: %s \n", p.Name())
		p.Device().StopScanning()
		p.Device().Connect(p);
	}
}

func onPeriphConnected(p gatt.Peripheral, err error) {
	log.Printf("Peripheral connected\n")

	services, err := p.DiscoverServices(nil)
	if err != nil {
		log.Printf("Failed to discover services, err: %s\n", err)
		return
	}

	for _, service := range services {

		if (service.UUID().Equal(uartServiceId)) {
			log.Printf("Service Found %s\n", service.Name())

			cs, _ := p.DiscoverCharacteristics(nil, service)

			for _, c := range cs {
				if (c.UUID().Equal(uartServiceTXCharId)) {
					log.Println("TX Characteristic Found")

					p.DiscoverDescriptors(nil, c)

					p.SetNotifyValue(c, func(c *gatt.Characteristic, b []byte, e error) {
							log.Printf("Got back %s\n", string(b))
					})
				}
			}
			for _, c := range cs {
					if (c.UUID().Equal(uartServiceRXCharId)) {
						log.Println("RX Characteristic Found")
						p.WriteCharacteristic(c, []byte{0x74}, true)
						log.Printf("Wrote %s\n", string([]byte{0x74}))
					}
			}
		}
	}
}

func onPeriphDisconnected(p gatt.Peripheral, err error) {
	log.Println("Disconnected")
}

func main() {
	var DefaultClientOptions = []gatt.Option{
		gatt.LnxMaxConnections(1),
		gatt.LnxDeviceID(-1, false),
	}

	d, err := gatt.NewDevice(DefaultClientOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s\n", err)
		return
	}

	d.Handle(
		gatt.PeripheralDiscovered(onPeriphDiscovered),
		gatt.PeripheralConnected(onPeriphConnected),
		gatt.PeripheralDisconnected(onPeriphDisconnected),
	)
	d.Init(onStateChanged)
	<-done
	log.Println("Done")
}

//http://christopherbird.co.uk/posts/talking-to-the-nrf8001-ble-chip-via-golang/