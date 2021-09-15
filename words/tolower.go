package words

func ToLower(word []rune) []rune {
	wordLength := len(word)

	for i := 0; i < wordLength; i++ {
		if 64 < word[i] && word[i] < 91 {
			word[i] = word[i] + 32
		}
	}

	return word
}
