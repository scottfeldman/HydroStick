package main

import (
	"machine"
	"time"
)

var Version string

func main() {

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	machine.InitADC()
	pin := machine.ADC{Pin: machine.A0}
	pin.Configure(machine.ADCConfig{})
	sensor := NewSensor(pin)

	blynk := NewBlynk(BlynkToken)

	setupWifi(WifiSsid, WifiPass)
	blynk.event("CONNECT")

	for {
		raw := sensor.read()
		percent := 100 - toPercent(raw, SensorMin, SensorMax)
		if PrintRaw {
			println(raw)
		} else {
			println(percent)
		}
		err := blynk.updateInt("v0", int(percent))
		if err != nil {
			println(err.Error())
			setupWifi(WifiSsid, WifiPass)
			blynk.event("CONNECT")
		}
		led.Set(!led.Get())
		time.Sleep(ProbeFreq)
	}
}
