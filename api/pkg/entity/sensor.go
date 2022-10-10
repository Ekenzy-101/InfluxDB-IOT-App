package entity

import (
	"math/rand"
	"time"
)

type Sensor struct{}

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

func NewSensor() *Sensor {
	rand.Seed(time.Now().UnixNano())
	return &Sensor{}
}

func (s *Sensor) Name() string {
	return "virtual_bme280"
}

func (s *Sensor) Temperature() int {
	return int(s.random(0, 100))
}

func (s *Sensor) Humidity() int {
	return int(s.random(0, 100))
}

func (s *Sensor) Pressure() int {
	return int(s.random(0, 100))
}

func (s *Sensor) Geo() *Coordinate {
	return &Coordinate{
		Latitude:  s.random(-100, 100),
		Longitude: s.random(-100, 100),
	}
}

func (s *Sensor) random(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
