/*
Package kociemba implements the famous two-phase algorithm.
*/
package kociemba

import (
	"math/rand"
)

// GetNewScramble returns scramble, the sequence of moves that will transform a solved cube to the returned layout.
func GetNewScramble() (scramble []string, layout string) {
	var mod = [3]string{"", "'", "2"}
	var x = [2]string{"R", "L"}
	var y = [2]string{"U", "D"}
	var z = [2]string{"F", "B"}
	curr := -1
	for i := 0; i < 25; i++ {
		next := rand.Intn(3)
		if next != curr {
			switch next {
			case 0:
				scramble = append(scramble, x[rand.Intn(2)]+mod[rand.Intn(3)])
			case 1:
				scramble = append(scramble, y[rand.Intn(2)]+mod[rand.Intn(3)])
			case 2:
				scramble = append(scramble, z[rand.Intn(2)]+mod[rand.Intn(3)])
			}
			curr = next
		} else {
			i--
		}
	}
	layout = "UUUUUUUUURRRRRRRRRFFFFFFFFFDDDDDDDDDLLLLLLLLLBBBBBBBBB"
	return
}

// Solve returns a string of moves that solves the provided scramble.
func Solve(scramble string) string {
	var cube = convert(scramble)
	cube = reduce(cube)
	return complete(cube)
}
