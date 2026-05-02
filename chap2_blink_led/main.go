package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

func main() {
	pin := machine.GPIO16
	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	led := ws2812.New(pin)

	for {
		led.WriteColors([]color.RGBA{
			// 検証したところ、しうやら G, R, Bの順番だと思われる
			{R: 0, G: 122, B: 0, A: 255},
		})
		time.Sleep(time.Millisecond * 500)

		led.WriteColors([]color.RGBA{
			{R: 0, G: 0, B: 0, A: 255},
		})
		time.Sleep(time.Millisecond * 500)
	}
}
