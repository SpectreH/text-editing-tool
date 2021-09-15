package words

// removes one element from array
func RemoveIndex(s []rune, index int) []rune {
	return append(s[:index], s[index+1:]...)
}

// adds object from 1d rune array to 2d rune array
func AddWord(wordList *[][]rune, wordToAdd []rune) {
	*wordList = append(*wordList, nil)
	for i := len(*wordList) - 1; i < len(*wordList); i++ {
		for k := 0; k < len(wordToAdd); k++ {
			(*wordList)[i] = append((*wordList)[i], wordToAdd[k])
		}
	}
}
