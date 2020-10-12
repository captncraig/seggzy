package grid

type Digit byte

// Grid is a rectangular grid of seven-segment displays
// Ordered from top left corner [x][y]
type Grid [][]Digit

// New creates a segment array with the given width and height
// All segments will be initialized to zero
func New(w, h int) Grid {
	arr := make(Grid, w)
	for i := range arr {
		arr[i] = make([]Digit, h)
	}
	return arr
}
