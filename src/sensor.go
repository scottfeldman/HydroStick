package main

import (
	"machine"
	"time"
)

type Sensor struct {
	pin machine.ADC

	beta  float64
	times int
	delay time.Duration
}

func NewSensor(pin machine.ADC) *Sensor {
	return &Sensor{
		pin:   pin,
		beta:  0.05,
		times: 100,
		delay: 1 * time.Millisecond,
	}
}

func (s *Sensor) read() float64 {

	if s.times == 0 {
		s.times = 1
	}

	value := float64(s.pin.Get())
	for i := 0; i < s.times; i++ {
		raw := s.pin.Get()
		value = value*(1-s.beta) + float64(raw)*s.beta
		if s.delay != 0 {
			time.Sleep(s.delay)
		}
	}

	return value
}

func toPercent(value, min, max float64) float64 {
	return ((value - min) / (max - min)) * 100
}
