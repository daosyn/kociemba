/*
This is the main implementation package.
It divides the implementation into two steps.
It also defines the cube struct.
*/

package kociemba

type firstPhaseCube struct {
	cornerOrientationCoordinate int // 0-2186
	edgeOrientationCoordinate   int // 0-2047
	ySliceCoordinate            int // 0-494
}

/*
phase 1:
    simplify scrambled cube into a state such
    that all edge pieces and corner pieces
    have an orientation of 0
    U and D consist of oriented pieces
*/
func reduce(cube [54]Facelet) [54]Facelet {
	return cube
}

type secondPhaseCube struct {
	cornerPermutationCoordinate int // 0-40319
	edgePermutationCoordinate   int // 0-40319
	ySliceCoordinate            int // 0-23
}

/*
phase 2:
    solve the reduced cube from phase 1
*/
func complete(cube [54]Facelet) string {
	return ""
}
