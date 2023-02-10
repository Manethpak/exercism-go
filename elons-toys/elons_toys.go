package elon

import "fmt"

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
