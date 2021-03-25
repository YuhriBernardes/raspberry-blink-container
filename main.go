package main

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio"
)

var (
// Use mcu pin 10, corresponds to physical pin 19 on the pi
)

func main() {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		logrus.WithError(err).Error("Failed to open RPIO device")
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pin to output mode
	pinNumber := 10
	pin := rpio.Pin(pinNumber)
	logrus.Infof("Setting pin %d as output", pinNumber)
	pin.Output()

	// Toggle pin 20 times
	for x := 0; x < 20; x++ {
		pin.Toggle()
		logrus.WithField("iteraction", x).Info("Pin toggled")
		time.Sleep(500 * time.Millisecond)
	}
}
