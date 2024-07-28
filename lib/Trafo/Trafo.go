package trafo

import "math/cmplx"

type T struct {
	/*
	   A B
	   C D
	*/
	A complex128
	B complex128
	C complex128
	D complex128
}

func (t T) On(z complex128) (tz complex128) {
	tz = (t.A*z + t.B) / (t.C*z + t.D)
	return
}

func (t T) det() (d complex128) {
	d = (t.A * t.D) - (t.B * t.C)
	return
}

func (t T) Inverse() T {
	det := t.det()
	r := T{t.D / det, -(t.B / det), -(t.C / det), t.A / det}
	return r
}

func (t T) Multiply(x T) T {
	var r T
	r.A = t.A*x.A + t.B*x.C
	r.B = t.A*x.B + t.B*x.D
	r.C = t.C*x.A + t.D*x.C
	r.D = t.C*x.B + t.D*x.D
	return r
}

func (t T) Fix() (f complex128) {
	f = t.A - t.D + cmplx.Sqrt(cmplx.Pow(t.D-t.A, 2)-4)/2*t.C
	return
}
