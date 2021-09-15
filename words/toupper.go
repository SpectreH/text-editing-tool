package words

func ToUpper(word []rune) []rune {
	wordLength := len(word)

	for i := 0; i < wordLength; i++ {
		if 96 < word[i] && word[i] < 123 {
			word[i] = word[i] - 32
		}
	}

	return word
}
