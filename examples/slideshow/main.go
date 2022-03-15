package main

import (
	"image/color"
	"log"

	"time"

	pngle "github.com/sago35/go-pngle"
	"github.com/sago35/tinydisplay/examples/initdisplay"
)

var (
	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
	green = color.RGBA{0, 255, 0, 255}
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	display := initdisplay.InitDisplay()
	pngle.SetCallback(func(x, y, w, h uint16, r, g, b, a uint8) {
		//fmt.Printf("%d %d %d %d %02X %02X %02X %02X\n", x, y, w, h, r, g, b, a)
		c := color.RGBA{
			R: r,
			G: g,
			B: b,
			A: a,
		}
		display.SetPixel(int16(x), int16(y), c)
	})

	b := []byte(out_org_png)
	pngle.DecodeFromBytes(b)

	for {
		time.Sleep(500 * time.Millisecond)
	}

	return nil
}
