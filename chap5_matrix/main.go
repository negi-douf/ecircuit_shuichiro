package main

import (
	"machine"
	"time"
)

const (
	regDecodeMode  = 0x09
	regIntensity   = 0x0A
	regScanLimit   = 0x0B
	regShutdown    = 0x0C
	regDisplayTest = 0x0F
	// Digitレジスタは 0x01-0z08 (行0-7に対応)
	regDigit0 = 0x01
)

type MAX7219 struct {
	spi machine.SPI
	cs  machine.Pin
}

func NewMAX7219(spi machine.SPI, cs machine.Pin) MAX7219 {
	cs.Configure(machine.PinConfig{Mode: machine.PinOutput})
	cs.High()
	return MAX7219{spi: spi, cs: cs}
}

func (d MAX7219) writeCommand(register, data byte) {
	d.cs.Low()
	d.spi.Transfer(register) // 送りたいだけなので、戻り値は使わない
	d.spi.Transfer(data)
	d.cs.High()
}

func (d MAX7219) configure() {
	d.writeCommand(regShutdown, 0x01)
	d.writeCommand(regDisplayTest, 0x00)
	d.writeCommand(regDecodeMode, 0x00)
	d.writeCommand(regScanLimit, 0x07)
	d.writeCommand(regIntensity, 0x04)
}

// "A" のビットパターン
var letterA = [8]byte{
	0b00011000,
	0b00111100,
	0b01100110,
	0b01100110,
	0b01111110,
	0b01100110,
	0b01100110,
	0b00000000,
}

func main() {
	machine.SPI1.Configure(machine.SPIConfig{
		Frequency: 1_000_000,      // 1MHz
		SCK:       machine.GPIO10, // CLK
		SDO:       machine.GPIO11, // データ出力
		SDI:       machine.GPIO12, // データ入力は今回使用しない
		Mode:      0,
	})

	display := NewMAX7219(*machine.SPI1, machine.GPIO9)
	display.configure()

	for i, pattern := range letterA {
		display.writeCommand(byte(regDigit0+i), pattern)
	}

	for {
		time.Sleep(time.Hour)
	}
}
