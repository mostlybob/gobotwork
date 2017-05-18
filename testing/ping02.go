// copied this from answer posted by @billettc
// https://github.com/hybridgroup/gobot/issues/246

package main

import (
	"fmt"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
	"time"
)

func main() {
	//r := raspi.NewAdaptor()  // original code
	r := firmata.NewAdaptor("/dev/ttyUSB0")

	trigPin := gpio.NewDirectPinDriver(r, "11")
	echoPin := gpio.NewDirectPinDriver(r, "12")
	/*  - this parts a bit different from the arduino code I was using
	    - I have the trigger & echo pins wired together and the code
	      reads/writes to the same pin
	    - this is how the reference sensor works (by Parallax)
	      - i.e. the trigger & echo come off the same pin
	    - the sensors I got have separate pins for echo & trigger
	*/

	// led := gpio.NewLedDriver(r, "7")  //original code
	led := gpio.NewLedDriver(r, "13")

	work := func() {

		gobot.Every(1*time.Second, func() {

			println("Starting probing ")
			led.Toggle()

			trigPin.DigitalWrite(byte(0))
			time.Sleep(2 * time.Microsecond)

			trigPin.DigitalWrite(byte(1))
			time.Sleep(10 * time.Microsecond)

			trigPin.DigitalWrite(byte(0))
			start := time.Now()
			end := time.Now()

			for {
				val, err := echoPin.DigitalRead()
				start = time.Now()

				if err != nil {
					println(err)
					break
				}

				if val == 0 {
					continue
				}

				break
			}

			for {
				val, err := echoPin.DigitalRead()
				end = time.Now()
				if err != nil {
					println(err)
					break
				}

				if val == 1 {
					continue
				}

				break
			}

			duration := end.Sub(start)
			durationAsInt64 := int64(duration)
			distance := duration.Seconds() * 34300
			distance = distance / 2 //one way travel time
			fmt.Printf("Duration : %v %v %v \n", distance, duration.Seconds(), durationAsInt64)
		})
	}

	robot := gobot.NewRobot("makeyBot",
		[]gobot.Connection{r},
		[]gobot.Device{trigPin, echoPin, led},
		work,
	)

	robot.Start()
}
