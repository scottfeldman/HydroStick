package main

import (
	"machine"
	"time"
)

var Version string

func main() {

	machine.InitADC()

	pin := machine.ADC{Pin: machine.A0}
	pin.Configure(machine.ADCConfig{})
	sensor := Sensor{beta: 0.05, lastValue: 0, adc: pin, scale: 0.01}

	setupWifi(WifiSsid, WifiPass)
	blynk := NewBlynk(BlynkToken)

	for {
		value := sensor.read(100, 10*time.Millisecond)
		println(int(value))
		blynk.updateInt("v0", int(value))
		time.Sleep(1 * time.Minute)
	}
}
