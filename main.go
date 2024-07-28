package main

import (
	grandmas "Schmottky/lib/GrandmasRecipes"
	matplotlib "Schmottky/lib/Receivers/MatPlotLib"
	dfs "Schmottky/lib/Senders/DFS"
	trafo "Schmottky/lib/Trafo"
	"fmt"
)

func main() {
	fmt.Println("Hi <3")
	// data

	//a := trafo.T{1, 0, 0 - 2i, 1}
	//b := trafo.T{1 - 1i, 1, 1, 1 + 1i}

	// p. 238
	//a, b := grandmas.SpecialParabolic(1.91+0.05i, 3+0i)
	// p. 293 - (i)
	//a, b := grandmas.SpecialParabolic(1.9134233 - 0.0362881i, 2 + 0i)
	// p. 293 - (ii)
	//a, b := grandmas.SpecialParabolic(1.8964073 - 0.0487530i, 2 + 0i)
	// p. 293 - (iii)
	a, b := grandmas.SpecialParabolic(1.91-0.05i, 2+0i)
	// p. 293 - (iv)
	//a, b := grandmas.SpecialParabolic(1.90378 - 0.03958i, 2 + 0i)
	// p. 272
	//a, b := grandmas.SpecialParabolic(1.64213876 - 0.76658841i, 2 + 0i)
	// p. 269
	//a, b := grandmas.SpecialParabolic(1.95859 - 0.01128i, 2 + 0i)

	A := a.Inverse()
	B := b.Inverse()

	var gens = [4]trafo.T{a, b, A, B}
	//fmt.Printf("%T", gens)

	/* start dfs */
	eps := 0.0001
	go dfs.Run(gens, eps)
	//go fromFile.Read("/dev/shm/schmooo.data")

	/*start receiver*/
	//pixel.StartDrawing()
	//img.Draw()
	//raw.ToFile("/dev/shm/schmooo.data")
	matplotlib.Plot()

}
