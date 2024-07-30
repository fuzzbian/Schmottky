package img

import (
	common "Schmottky/lib/Senders"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"

	lua "github.com/yuin/gopher-lua"
)

var MaxCoordinate float64 = 1.03
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

var px, py, l float64

func GetPointX(L *lua.LState) int {
	L.Push(lua.LNumber(px))
	return 1
}
func GetPointY(L *lua.LState) int {
	L.Push(lua.LNumber(py))
	return 1
}
func GetLev(L *lua.LState) int {
	L.Push(lua.LNumber(l))
	return 1
}
func GetMax(L *lua.LState) int {
	L.Push(lua.LNumber(MaxCoordinate))
	return 1
}

func getColorLua(p complex128, lev int) color.RGBA {
	fileName := "./lib/Receivers/Image/color.lua"
	px = real(p)
	py = imag(p)
	l = float64(lev)

	// set up
	L := lua.NewState()
	defer L.Close()

	// export functions as external API
	L.SetGlobal("getX", L.NewFunction(GetPointX))
	L.SetGlobal("getY", L.NewFunction(GetPointY))
	L.SetGlobal("getL", L.NewFunction(GetLev))
	L.SetGlobal("getMax", L.NewFunction(GetMax))

	// run lua script
	if err := L.DoFile(fileName); err != nil {
		fmt.Println(err)
	}

	// get return value
	ret := L.Get(-1)
	L.Pop(1)

	r, _ := strconv.ParseFloat(ret.(*lua.LTable).RawGet(lua.LNumber(1)).String(), 64)
	g, _ := strconv.ParseFloat(ret.(*lua.LTable).RawGet(lua.LNumber(2)).String(), 64)
	b, _ := strconv.ParseFloat(ret.(*lua.LTable).RawGet(lua.LNumber(3)).String(), 64)
	a, _ := strconv.ParseFloat(ret.(*lua.LTable).RawGet(lua.LNumber(4)).String(), 64)

	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}

}

func getColor(p complex128, lev int) (c color.RGBA) {

	// use external lua script for color calculation
	c = getColorLua(p, lev)

	// use internal calculation, much faster, less dynamic
	/*
		r := uint8(255)
		g := uint8((float64(lev) / float64(dfs.MaxLevel)) * float64(255))
		b := uint8((cmplx.Abs(0-p) / MaxCoordinate) * float64(255))
		//fmt.Printf("%v\n", b)
		a := uint8(255)

		c = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
	*/

	return

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
