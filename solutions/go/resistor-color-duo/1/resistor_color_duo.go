package resistorcolorduo
//black: 0 brown: 1 red: 2 orange: 3 yellow: 4 green: 5 blue: 6 violet: 7 grey: 8 white: 9

var colorMap = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}
// Value calculates the two-digit integer value from the first two resistor color bands.
func Value(colors []string) int {
	return colorMap[colors[0]]*10 + colorMap[colors[1]]
}
