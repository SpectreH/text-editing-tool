package strings

func ToLower(s string, charsCounter int) string {
	sentence := []rune(s)

	finalString := ""

	charsToChange := len(s)

	if charsCounter != 0 {
		charsToChange = charsCounter
	}

	for i := 0; i < charsToChange; i++ {
		if 64 < sentence[i] && sentence[i] < 91 {
			sentence[i] = sentence[i] + 32
		}

		finalString = finalString + string(sentence[i])
	}

	return finalString
}
