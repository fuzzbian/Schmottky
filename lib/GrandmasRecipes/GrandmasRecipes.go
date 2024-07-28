package grandmas

import (
	trafo "Schmottky/lib/Trafo"
	"math/cmplx"
)

// p. 229
func SpecialParabolic(ta, tb complex128) (a, b trafo.T) {
	p := -(ta * tb)
	q := (ta * ta) + (tb * tb)
	tab := -(p / 2) - cmplx.Sqrt((p*p)-(4*q))/2

	z0 := ((tab - 2) * tb) / ((tb * tab) - (2 * ta) + ((0 + 2i) * tab))

	b.A = (tb - (0 + 2i)) / 2
	b.B = tb / 2
	b.C = tb / 2
	b.D = (tb + (0 + 2i)) / 2

	var ab trafo.T
	ab.A = tab / 2
	ab.B = (tab - 2) / (2 * z0)
	ab.C = ((tab + 2) * z0) / 2
	ab.D = tab / 2

	a = ab.Multiply(b.Inverse())
	return
}

/*
# Grandma's recipes
###################################################

# p. 229
def special_parabolic(t_a, t_b):
    #(2)
    p = -(t_a * t_b)
    q = (t_a**2) + (t_b**2)
    t_ab = -(p/2) - sqrt((p**2) - (4*q))/2

    #(3)
    z_0 = ((t_ab - 2)*t_b)/((t_b*t_ab)-(2*t_a)+((0 + 2j)*t_ab))

    #(4)
    b_a = (t_b - (0 + 2j))/2
    b_b = t_b/2
    b_c = t_b/2
    b_d = (t_b + (0 + 2j))/2
    b = Trafo([[b_a, b_b], [b_c, b_d]])

    #(5)
    ab_a = t_ab/2
    ab_b = (t_ab - 2)/(2*z_0)
    ab_c = ((t_ab + 2) *z_0)/2
    ab_d = t_ab/2
    ab = Trafo([[ab_a, ab_b], [ab_c, ab_d]])
    a = ab |x| b.get_inv()

    return a, b
*/
