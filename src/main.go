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
	sensor := Sensor{beta: 0.05, lastValue: 0, adc: pin, scale: 0.01}

	blynk := NewBlynk(BlynkToken)

	setupWifi(WifiSsid, WifiPass)
	blynk.event("CONNECT")

	for {
		println("main: start ", time.Now().Unix())
		value := sensor.read(100, 10*time.Millisecond)
		println(int(value))
		err := blynk.updateInt("v0", int(value))
		println("main: err")
		if err != nil {
			println(err.Error())
			setupWifi(WifiSsid, WifiPass)
			blynk.event("CONNECT")
		}
		led.Set(!led.Get())
		println("main: end", time.Now().Unix())
		time.Sleep(1 * time.Minute)
	}
}
