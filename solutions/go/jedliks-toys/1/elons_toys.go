package elon

import "fmt"

// TODO: define the 'Drive()' method
func (d *Car) Drive() {
	if d.batteryDrain < d.battery {
		d.battery -= d.batteryDrain
		d.distance += d.speed
	}
}

// TODO: define the 'DisplayDistance() string' method
func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c *Car) CanFinish(trackDistance int) bool {
	carTime := float64(trackDistance) / float64(c.speed)

	return carTime*float64(c.batteryDrain) <= float64(c.battery)
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
