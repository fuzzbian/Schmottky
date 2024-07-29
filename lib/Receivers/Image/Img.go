package img

import (
	common "Schmottky/lib/Senders"
	dfs "Schmottky/lib/Senders/DFS"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
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

func getColor(p complex128, lev int) color.RGBA {

	r := uint8(255)
	g := uint8((cmplx.Abs(0-p) / MaxCoordinate) * float64(255))
	b := uint8((float64(lev) / float64(dfs.MaxLevel)) * float64(255))
	//fmt.Printf("%v\n", b)
	a := uint8(255)

	return color.RGBA{r, g, b, a}
}

func Draw() {
	myImg := image.NewRGBA(image.Rect(0, 0, int(MaxPixel), int(MaxPixel)))

	fill(myImg, color.RGBA{0, 0, 0, 255})
	var lev int = -1
	for {
		if common.LevTracking {
			lev = <-common.LevChannel
		}
		p, ok := <-common.PointChannel
		if !ok {
			break
		}
		x, y := point2Pixel(p)
		color := getColor(p, lev)
		myImg.SetRGBA(x, y, color)
	}

	out, _ := os.Create("Schmottky.png")
	png.Encode(out, myImg)
	out.Close()

}
