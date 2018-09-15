// divides implementation into two steps

package kociemba

// phase 1:
// simplify scrambled cube into a state such
// that all edge pieces and corner pieces
// have an orientation of 0
// U and D consist of oriented pieces
func reduce(cube [54]Facelet) [54]Facelet {
	return cube
}

// phase 2:
// solve the reduced cube from phase 1
func complete(cube [54]Facelet) string {
	return ""
}
