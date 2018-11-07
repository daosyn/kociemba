package kociemba

/*
 	N_TWIST = 2187  # 3^7 possible corner orientations
    N_FLIP = 2048   # 2^11 possible edge flips
    N_SLICE1 = 495  # 12 choose 4 possible positions of FR,FL,BL,BR edges
    N_SLICE2 = 24   # 4! permutations of FR,FL,BL,BR edges in phase2
    N_PARITY = 2    # 2 possible corner parities
    N_URFtoDLF = 20160  # 8!/(8-6)! permutation of URF,UFL,ULB,UBR,DFR,DLF corners
    N_FRtoBR = 11880    # 12!/(12-4)! permutation of FR,FL,BL,BR edges
    N_URtoUL = 1320     # 12!/(12-3)! permutation of UR,UF,UL edges
    N_UBtoDF = 1320     # 12!/(12-3)! permutation of UB,DR,DF edges
    N_URtoDF = 20160    # 8!/(8-6)! permutation of UR,UF,UL,UB,DR,DF edges in phase2

    N_URFtoDLB = 40320  # 8! permutations of the corners
    N_URtoBR = 479001600    # 8! permutations of the corners

    N_MOVE = 18
*/
var maxTwists = 2187
var maxFlips = 2048
var moves = 18

type (
	// Corner represents one of the eight corners of the cube.
	Corner int

	// CornerOrientation represents whether the corner has been rotated from its original position.
	CornerOrientation int

	// Edge represents one of the twelve edge pieces of the cube.
	Edge int

	// EdgeOrientation represents whether the edge has been flipped.
	EdgeOrientation int

	// Face represents one of the six sides of the cube.
	Face int

	// Facelet represents each of the nine faces of each side of the cube.
	Facelet int
)

// URF represents the corner that is shared by the U, R, and F face.
const (
	URF Corner = iota
	UFL
	ULB
	UBR
	DFR
	DLF
	DBL
	DRB
)

// Corner can be oriented, rotated clockwise, or rotated counter-clockwise.
const (
	CornerOriented CornerOrientation = iota
	CornerCW
	CornerCCW
)

// UR represents the edge that has a U color on one side and R color on the other.
const (
	UR Edge = iota
	UF
	UL
	UB
	DR
	DF
	DL
	DB
	FR
	FL
	BL
	BR
)

// Edge can either be oriented or flipped.
const (
	EdgeOriented EdgeOrientation = iota
	EdgeFlipped
)

// Each face is used to identify which side should be turned or which color a facelet is set to.
const (
	U Face = iota
	R
	F
	D
	L
	B
)

// These are static and represent from top left to bottom right which particular facelet is being referenced.
const (
	U1 Facelet = iota
	U2
	U3
	U4
	U5
	U6
	U7
	U8
	U9
	R1
	R2
	R3
	R4
	R5
	R6
	R7
	R8
	R9
	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	F9
	D1
	D2
	D3
	D4
	D5
	D6
	D7
	D8
	D9
	L1
	L2
	L3
	L4
	L5
	L6
	L7
	L8
	L9
	B1
	B2
	B3
	B4
	B5
	B6
	B7
	B8
	B9
)
