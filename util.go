package kociemba

func convert(cubestring string) [54]Facelet {
	var cube = [54]Facelet{}
	for i := 0; i < 54; i++ {
		cube[i] = Facelet(i)
	}
	return cube
}
