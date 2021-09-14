package piscine

func ToLower(s string) string {
	sentence := []rune(s)

	finalString := ""

	for i := 0; i < len(sentence); i++ {
		if 64 < sentence[i] && sentence[i] < 91 {
			sentence[i] = sentence[i] + 32
		}

		finalString = finalString + string(sentence[i])
	}

	return finalString
}
