package strings

func Cap(s string, charsCounter int) string {
	word := []rune(s)

	finalString := ""

	wordLength := len(s)
	charsToChange := 1

	if charsCounter != 0 {
		charsToChange = charsCounter
	}

	for i := 0; i < wordLength; i++ {
		if (96 < word[i] && word[i] < 123) && i < charsToChange {
			word[i] = word[i] - 32
		}

		finalString = finalString + string(word[i])
	}

	return finalString
}
