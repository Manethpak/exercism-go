package elon

import (
	"fmt"
)

// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int

	battery  int
	distance int
}

func (c *Car) Drive() {
	if c.battery < c.batteryDrain {
		return
	}

	c.battery -= c.batteryDrain
	c.distance += c.speed
}

func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

func (c *Car) CanFinish(trackDistant int) bool {
	for c.battery >= c.batteryDrain {
		c.Drive()
		trackDistant -= c.speed
		if trackDistant <= 0 {
			break
		}
	}

	return trackDistant <= 0
}

// NewCar creates a new car with given specifications.
func NewCar(speed, batteryDrain int) *Car {
	return &Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}
