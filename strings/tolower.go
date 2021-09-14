package strings

func ToLower(s string, charsCounter int) string {
	sentence := []rune(s)

	finalString := ""

	wordLength := len(s)
	charsToChange := len(s)

	if charsCounter != 0 {
		charsToChange = charsCounter
	}

	for i := 0; i < wordLength; i++ {
		if (64 < sentence[i] && sentence[i] < 91) && i < charsToChange {
			sentence[i] = sentence[i] + 32
		}

		finalString = finalString + string(sentence[i])
	}

	return finalString
}
