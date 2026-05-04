package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.GPIO15
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		// 電圧を High 3.3Vにする
		led.High()
		time.Sleep(time.Millisecond * 500)
		led.Low()
		time.Sleep(time.Millisecond * 500)
	}
}
