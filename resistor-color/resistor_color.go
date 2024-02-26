package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
	colorsList := []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
	return colorsList
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	resistanceMap := make(map[string]int)
	resistanceMap["black"] = 0
	resistanceMap["red"] = 2
	resistanceMap["brown"] = 1
	resistanceMap["orange"] = 3
	resistanceMap["yellow"] = 4
	resistanceMap["green"] = 5
	resistanceMap["blue"] = 6
	resistanceMap["violet"] = 7
	resistanceMap["grey"] = 8
	resistanceMap["white"] = 9

	return resistanceMap[color]
}
