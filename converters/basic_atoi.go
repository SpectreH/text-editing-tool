package converters

// converts from rune array number to int number
func BasicAtoi(s_rune []rune) int {
	var result int = 0
	var runeLength int = len(s_rune) - 1
	var multiplier int = 1
	for i := runeLength; 0 <= i; i-- {
		result = result + ((int(uint(s_rune[i])) - 48) * multiplier)
		multiplier = multiplier * 10
	}
	return result
}
