package digits

import . "github.com/captncraig/seggzy/pkg/constants/segments"

// D0 - D9 are single digits
const (
	D0 = A | B | C | D | E | F
	D1 = B | C
	D2 = A | B | D | E | G
	D3 = A | B | C | D | G
	D4 = B | C | F | G
	D5 = A | C | D | F | G
	D6 = A | C | D | E | F | G
	D7 = A | B | C
	D8 = A | B | C | D | E | F | G
	D9 = A | B | C | D | F | G
)
