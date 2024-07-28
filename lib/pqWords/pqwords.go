package pqwords

import (
	grandmas "Schmottky/lib/GrandmasRecipes"
	trafo "Schmottky/lib/Trafo"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func GetTrafos(p, q int) (a, b trafo.T) {

	p_str := strconv.Itoa(p)
	q_str := strconv.Itoa(q)

	out, err := exec.Command("python3", "./lib/pqWords/pq.py", p_str, q_str).Output()
	if err != nil {
		log.Fatal(err)
	}

	c_str := strings.Split(string(out), ";")
	r, _ := strconv.ParseFloat(c_str[0], 64)
	i, _ := strconv.ParseFloat(c_str[1], 64)
	c := complex(r, i)

	a, b = grandmas.SpecialParabolic(c, 2+0i)
	return
}
