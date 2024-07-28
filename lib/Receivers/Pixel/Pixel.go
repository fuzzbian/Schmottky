package pixel

import (
	common "Schmottky/lib/Senders"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var MaxCoordinate float64 = 1.3
var MaxPixel float64 = 1024

func point2Pixel(z complex128) pixel.Vec {
	p_x := (MaxPixel / 2) + (real(z)/MaxCoordinate)*(MaxPixel/2)
	p_y := (MaxPixel / 2) + (imag(z)/MaxCoordinate)*(MaxPixel/2)

	return pixel.V(p_x, p_y)
}

func artist() {
	cfg := pixelgl.WindowConfig{
		Title:  "Schmottky",
		Bounds: pixel.R(0, 0, MaxPixel, MaxPixel),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.Clear(colornames.Black)
	started := false
	var firstP, oldP pixel.Vec
	for !win.Closed() {
		imd := imdraw.New(nil)

		point, ok := <-common.PointChannel
		p := point2Pixel(point)
		if ok {
			if !started {
				firstP = p
				oldP = p
				started = true
				continue
			}

			imd.Color = colornames.Blueviolet
			imd.EndShape = imdraw.RoundEndShape
			imd.Push(oldP, p)
			imd.Line(1)
			oldP = p
		} else {
			p = firstP
		}

		imd.Color = colornames.Blueviolet
		imd.EndShape = imdraw.RoundEndShape
		imd.Push(oldP, p)
		imd.Line(1)
		oldP = p

		imd.Draw(win)
		win.Update()

	}
}

func StartDrawing() {
	pixelgl.Run(artist)
}
