package main

import (
	img "Schmottky/lib/Receivers/Image"
	fromFile "Schmottky/lib/Senders/File"
	trafo "Schmottky/lib/Trafo"
	"fmt"
)

func main() {
	fmt.Println("Hi <3")
	//a := trafo.T{1, 0, 0 - 2i, 1}
	//b := trafo.T{1 - 1i, 1, 1, 1 + 1i}

	//p. 293 - (iii)
	a := trafo.T{0.9550000000000001 - 0.025000000000000022i, 0.025000000000000022 - 0.04499999999999993i, -0.025000000000000022 - 1.9550000000000005i, 0.9550000000000002 - 0.02499999999999991i}
	b := trafo.T{1 - 1i, 1, 1, 1 + 1i}

	A := a.Inverse()
	B := b.Inverse()

	var gens = [4]trafo.T{a, b, A, B}
	fmt.Printf("%v", gens)

	/* start dfs */
	//eps := 0.0001
	//go dfs.Run(gens, eps)
	go fromFile.Read("/dev/shm/schmooo.data")

	/*start receiver*/
	//pixel.StartDrawing()
	img.Draw()
	//raw.ToFile("/dev/shm/schmooo.data")

}
