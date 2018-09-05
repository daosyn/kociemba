// simple cube representation
package kociemba

func permuteFace(cube [54]Facelet, face Face) [54]Facelet {
	var newcube [54]Facelet
	for i, facelet := range FaceMove[face] {
		newcube[facelet] = cube[i]
	}
	return newcube
}
