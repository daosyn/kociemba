package kociemba

import "fmt"

var faceMove = [6][54]Facelet{
	{U3, U6, U9, U2, U5, U8, U1, U4, U7,
		F1, F2, F3, R4, R5, R6, R7, R8, R9,
		L1, L2, L3, F4, F5, F6, F7, F8, F9,
		D1, D2, D3, D4, D5, D6, D7, D8, D9,
		B1, B2, B3, L4, L5, L6, L7, L8, L9,
		R1, R2, R3, B4, B5, B6, B7, B8, B9},
	{U3, U6, U9, U2, U5, U8, U1, U4, U7,
		F1, F2, F3, R4, R5, R6, R7, R8, R9,
		L1, L2, L3, F4, F5, F6, F7, F8, F9,
		D1, D2, D3, D4, D5, D6, D7, D8, D9,
		B1, B2, B3, L4, L5, L6, L7, L8, L9,
		R1, R2, R3, B4, B5, B6, B7, B8, B9},
	{U3, U6, U9, U2, U5, U8, U1, U4, U7,
		F1, F2, F3, R4, R5, R6, R7, R8, R9,
		L1, L2, L3, F4, F5, F6, F7, F8, F9,
		D1, D2, D3, D4, D5, D6, D7, D8, D9,
		B1, B2, B3, L4, L5, L6, L7, L8, L9,
		R1, R2, R3, B4, B5, B6, B7, B8, B9},
	{U3, U6, U9, U2, U5, U8, U1, U4, U7,
		F1, F2, F3, R4, R5, R6, R7, R8, R9,
		L1, L2, L3, F4, F5, F6, F7, F8, F9,
		D1, D2, D3, D4, D5, D6, D7, D8, D9,
		B1, B2, B3, L4, L5, L6, L7, L8, L9,
		R1, R2, R3, B4, B5, B6, B7, B8, B9},
	{U3, U6, U9, U2, U5, U8, U1, U4, U7,
		F1, F2, F3, R4, R5, R6, R7, R8, R9,
		L1, L2, L3, F4, F5, F6, F7, F8, F9,
		D1, D2, D3, D4, D5, D6, D7, D8, D9,
		B1, B2, B3, L4, L5, L6, L7, L8, L9,
		R1, R2, R3, B4, B5, B6, B7, B8, B9},
	{U3, U6, U9, U2, U5, U8, U1, U4, U7,
		F1, F2, F3, R4, R5, R6, R7, R8, R9,
		L1, L2, L3, F4, F5, F6, F7, F8, F9,
		D1, D2, D3, D4, D5, D6, D7, D8, D9,
		B1, B2, B3, L4, L5, L6, L7, L8, L9,
		R1, R2, R3, B4, B5, B6, B7, B8, B9},
}

func permuteFace(face []Facelet) []Facelet {
	for _, color := range face {
		fmt.Println(color)
	}
	return face
}
