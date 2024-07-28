package img

import (
	dfs "Schmottky/lib/DFS"
	"image"
	"image/color"
	"image/png"
	"os"
)

var MaxCoordinate float64 = 1.3
var MaxPixel float64 = 1e4

func point2Pixel(z complex128) (int, int) {
	p_x := ((MaxPixel - 1) / 2) + (real(z)/MaxCoordinate)*((MaxPixel-1)/2)
	p_y := ((MaxPixel - 1) / 2) - (imag(z)/MaxCoordinate)*((MaxPixel-1)/2)

	return int(p_x), int(p_y)
}

func fill(img *image.RGBA, c color.RGBA) {
	for i := 0; i < int(MaxPixel); i++ {
		for j := 0; j < int(MaxPixel); j++ {
			img.SetRGBA(i, j, c)
		}
	}
}

func Draw() {
	myImg := image.NewRGBA(image.Rect(0, 0, int(MaxPixel), int(MaxPixel)))

	fill(myImg, color.RGBA{0, 0, 0, 255})

	for {
		p, ok := <-dfs.PointChannel
		if !ok {
			break
		}
		x, y := point2Pixel(p)
		myImg.SetRGBA(x, y, color.RGBA{255, 255, 255, 255})
	}

	out, _ := os.Create("Schmottky.png")
	png.Encode(out, myImg)
	out.Close()

}
