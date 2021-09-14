package strings

func ToUpper(s string) string {
	sentence := []rune(s)

	finalString := ""

	for i := 0; i < len(sentence); i++ {
		if 96 < sentence[i] && sentence[i] < 123 {
			sentence[i] = sentence[i] - 32
		}

		finalString = finalString + string(sentence[i])
	}

	return finalString
}
