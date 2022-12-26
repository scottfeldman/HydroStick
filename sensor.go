package main

import (
	"machine"
	"time"
)

const sensorMax = 500
const sensorMin = 250
const sensorMid = (sensorMax + sensorMin) / 2

type Sensor struct {
	adc       machine.ADC
	lastValue float64
	beta      float64
	scale     float64
}

func (s *Sensor) read(times int, delay time.Duration) float64 {

	if times == 0 {
		times = 1
	}

	for i := 0; i < times; i++ {
		raw := s.adc.Get()
		if i == 0 {
			s.lastValue = float64(raw)
		} else {
			s.lastValue = s.lastValue*(1-s.beta) + float64(raw)*s.beta
		}
		if delay != 0 {
			time.Sleep(delay)
		}
	}

	return s.lastValue * s.scale
}
