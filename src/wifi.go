package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/net"
	"tinygo.org/x/drivers/wifinina"
)

var (
	spi     = machine.NINA_SPI
	adaptor *wifinina.Device
)

func setupWifi(ssid, pass string) {
	configureWifinina()
	connectToAP(ssid, pass)
}

func configureWifinina() {
	spi.Configure(machine.SPIConfig{
		Frequency: 8 * 1e6,
		SDO:       machine.NINA_SDO,
		SDI:       machine.NINA_SDI,
		SCK:       machine.NINA_SCK,
	})

	adaptor = wifinina.New(spi,
		machine.NINA_CS,
		machine.NINA_ACK,
		machine.NINA_GPIO0,
		machine.NINA_RESETN)

	net.ActiveDevice = nil
	adaptor.Configure()
}

func connectToAP(ssid, pass string) {
	time.Sleep(2 * time.Second)
	trace("Connecting to " + ssid)
	err := adaptor.ConnectToAccessPoint(ssid, pass, 10*time.Second)
	if err != nil {
		for {
			println(err)
			time.Sleep(1 * time.Second)
		}
	}

	trace("Connected.")

	time.Sleep(2 * time.Second)
	ip, _, _, err := adaptor.GetIP()
	for ; err != nil; ip, _, _, err = adaptor.GetIP() {
		println(err.Error())
		time.Sleep(1 * time.Second)
	}
	trace(ip.String())
}
