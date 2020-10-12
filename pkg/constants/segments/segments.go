package segments

// A-G and DP represent individual segments
// You may combine them with binary OR to form complex shapes
const (
	A = 1 << iota
	B
	C
	D
	E
	F
	G
	DP = 0b10000000
)
