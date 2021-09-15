package words

func Cap(word []rune) []rune {
	for i := 0; i < 1; i++ {
		if 96 < word[i] && word[i] < 123 {
			word[i] = word[i] - 32
		}
	}

	return word
}
