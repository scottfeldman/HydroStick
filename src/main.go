package main

import (
	"machine"
	"time"

	_ "tinygo.org/x/drivers/netdev"
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
			blynk.event("CONNECT")
		}
		led.Set(!led.Get())
		time.Sleep(ProbeFreq)
	}
}
