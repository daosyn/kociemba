package kociemba

import "testing"
import "fmt"

func TestPermuteFace(t *testing.T) {
	fmt.Println(permuteFace(faceMove[U], U))
	fmt.Println(permuteFace(faceMove[R], D))
	fmt.Println(permuteFace(faceMove[B], L))
	fmt.Println(permuteFace(faceMove[F], B))
}
