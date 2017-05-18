package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor("/dev/ttyUSB0")
	led := gpio.NewLedDriver(firmataAdaptor, "13")
	ping := gpio.NewDirectPinDriver(firmataAdaptor, "2")

	work := func() {
		gobot.Every(2*time.Second, func() {
			led.Toggle()

			ping.DigitalWrite(0)
			time.Sleep(2 * time.Microsecond)
			ping.DigitalWrite(1)
			time.Sleep(5 * time.Microsecond)
			ping.DigitalWrite(0)

			ping.DigitalRead()
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led},
		work,
	)

	robot.Start()
}
