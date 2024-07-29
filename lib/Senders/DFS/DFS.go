package dfs

import (
	common "Schmottky/lib/Senders"
	trafo "Schmottky/lib/Trafo"
	"fmt"

	//"fmt"
	"math/cmplx"
)

const MaxLevel = 1e2

var word [MaxLevel + 2]trafo.T
var tags [MaxLevel + 2]int
var begpt [4]complex128
var endpt [4]complex128
var lev int = 0

func mod4(a int) int {
	return (a%4 + 4) % 4
}

// p. 185
func initPts(gens [4]trafo.T) {
	for i := 0; i < 4; i++ {
		begpt[i] = gens[(i+1)%4].Multiply(gens[(i+2)%4]).Multiply(gens[(i+3)%4]).Multiply(gens[i]).Fix()
		endpt[i] = gens[mod4(i+3)].Multiply(gens[mod4(i+2)]).Multiply(gens[mod4(i+1)]).Multiply(gens[i]).Fix()
	}
}

// p. 148
func goForward(gens [4]trafo.T) {
	lev++
	tags[lev] = mod4(tags[lev-1] + 1)
	word[lev] = word[lev-1].Multiply(gens[tags[lev]])
	//fmt.Println("forward")
}

func availableTurn() bool {
	//fmt.Printf("ther is a turn %v, %v", (tags[lev+1]-1)%4, (tags[lev]+2)%4)

	return !(mod4(tags[lev+1]-1) == mod4(tags[lev]+2))
}

func turnAndGoForward(gens [4]trafo.T) {
	tags[lev+1] = mod4(tags[lev+1] - 1)
	//fmt.Println("turn+forward")
	if lev == -1 {
		word[0] = gens[tags[0]]
		if tags[0] == 0 {
			return
		}
	} else {
		//fmt.Printf("-- l: %v t[l]: %v", lev, tags[lev+1])
		word[lev+1] = word[lev].Multiply(gens[tags[lev+1]])
	}
	lev += 1
}

func branchTermination(oldP complex128, eps float64) (bool, complex128) {
	newP := word[lev].On(endpt[tags[lev]])
	//fmt.Printf("%v\n", lev)
	if cmplx.Abs(oldP-newP) < eps || lev == MaxLevel {
		if common.LevTracking {
			common.LevChannel <- lev
		}
		return true, newP
	}
	return false, oldP
}

func Run(gens [4]trafo.T, eps float64) {
	// init stuff
	word[0] = gens[0]
	tags[0] = 0
	initPts(gens)
	terminate := false
	p := 1000 + 1000i
	count := 0
	fmt.Println("DSF started...")
	for {
		for {
			terminate, p = branchTermination(p, eps)
			if terminate {
				// TODO: do something with p
				common.PointChannel <- p
				count++
				//fmt.Printf("%v: %v\n", count, p)
				break
			}
			goForward(gens)
		}

		for {
			lev-- // goBackward
			if lev == -1 || availableTurn() {
				break
			}
		}
		turnAndGoForward(gens)
		//fmt.Printf("%v %v\n", lev, tags[0])
		if lev == -1 && tags[0] == 0 {
			break
		}
	}
	//fmt.Printf("thinl im done\n")
	close(common.PointChannel)
	if common.LevTracking {
		close(common.LevChannel)
	}
	fmt.Println("Done.")
}
