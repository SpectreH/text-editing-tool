package main

import (
	"io/ioutil"
	"os"
	"text-editing-tool/converters"
	"text-editing-tool/errors"
	"text-editing-tool/words"
)

// Own struct for modificator from the MODSLIST to save valuable information
type Modificator struct {
	class      string
	lenght     int
	specialNum int
}

var MODSLIST [5]string

func main() {
	MODSLIST = [5]string{"hex", "bin", "up", "low", "cap"}

	if len(os.Args) != 3 {
		errors.PrintError(1)
	}

	fileSample := os.Args[1]
	fileResult := os.Args[2]

	file, err := ioutil.ReadFile(fileSample)
	CheckFile(err, fileSample)

	text := converters.TranslateToRuneSlice(file)

	var wordList [][]rune
	var currentWord []rune
	var punMarkList []rune

	var modValues Modificator
	var modFound bool = false
	var lastChar bool = false

	for i := 0; i < len(text); i++ {
		// If new-line and return in file, start new word
		if text[i] == 10 {
			currentWord = append(currentWord, text[i])
			words.AddWord(&wordList, currentWord)
			currentWord = nil
			punMarkList = nil
			continue
		}
		// Checks only for number, letter or symbol. Excluded punctual marks
		if text[i] != 32 && text[i] != 33 && text[i] != 44 && text[i] != 46 && text[i] != 58 && text[i] != 59 && text[i] != 63 {
			currentWord = append(currentWord, text[i])
		} else {
			FindMod(i, &text, &modValues, &modFound)

			if currentWord != nil {
				words.AddWord(&wordList, currentWord)
			}

			if modFound {
				modFound = false
				UseMod(modValues, &wordList)
			}

			// Adds to the temp list for future appending punctual marks as a one word
			for m := i; m < len(text); m++ {
				if text[m] == 32 || text[m] == 33 || text[m] == 44 || text[m] == 46 || text[m] == 58 || text[m] == 59 || text[m] == 63 {
					punMarkList = append(punMarkList, text[m])
					if m == len(text)-1 {
						lastChar = true
					}
				} else {
					i = m - 1
					break
				}
			}

			// Adds word or punctual marks to the list for future formating
			words.AddWord(&wordList, punMarkList)
			currentWord = nil
			punMarkList = nil

			if lastChar {
				break
			}
		}
	}

	if currentWord != nil {
		words.AddWord(&wordList, currentWord)
	}

	FormatPunctList(&wordList)
	FormatArticle(&wordList)
	FormatQuotes(&wordList)

	SaveFormatedText(&wordList, fileResult)
}

func CheckFile(e error, fileName string) {
	if e != nil {
		errors.PrintError(2)
		os.Exit(1)
	}
}

// Starts checking then '(' found. Checks if mod exists
func CheckMod(modStartingIndex int, text []rune) Modificator {
	// Starts checking then '(' found. Checks if mod exists
	var firstCheckResult Modificator
	var modFound bool

	for k := 0; k < len(MODSLIST); k++ {
		currentMod := []rune(MODSLIST[k])
		modFound = true

		for j := 0; j < len(currentMod); j++ {
			if currentMod[j] == text[modStartingIndex+j+1] {
				continue
			}
			modFound = false
			break
		}

		if modFound {
			firstCheckResult.class = MODSLIST[k]
			firstCheckResult.lenght = len(MODSLIST[k])
			break
		}
	}

	if !modFound {
		return firstCheckResult
	}

	return CheckModSpecNumber(modStartingIndex, text, firstCheckResult)
}

// Checks if optional arguments exists
func CheckModSpecNumber(modStartingIndex int, text []rune, checkResult Modificator) Modificator {
	var specialNumber []rune
	index := modStartingIndex + checkResult.lenght + 1

	if text[index] == 44 && text[index+1] == 32 {
		for m := index + 2; m < len(text); m++ {
			if 47 < text[m] && text[m] < 58 {
				specialNumber = append(specialNumber, text[m])
			} else if text[m] == 41 {
				checkResult.lenght = checkResult.lenght + 2
				break
			} else {
				errors.PrintError(3)
			}
		}
	} else if text[index] == 41 {
		checkResult.specialNum = 0
	} else {
		errors.PrintError(3)
	}

	checkResult.specialNum = converters.BasicAtoi(specialNumber)
	checkResult.lenght = checkResult.lenght + 2 + len(specialNumber)

	return checkResult
}

// Tries to find mod until new word starts
func FindMod(i int, text *[]rune, modValues *Modificator, modFound *bool) {
	for furtherIndex := i + 1; furtherIndex < len(*text); furtherIndex++ {
		if (*text)[furtherIndex] == 40 {
			*modValues = CheckMod(i+1, *text)

			if modValues.lenght != 0 {
				delIndex := i
				for charsToDel := modValues.lenght; charsToDel >= 0; charsToDel-- {
					*text = words.RemoveIndex(*text, delIndex)
				}
				*modFound = true
				break
			} else {
				break
			}
		}
	}
}

// Shows what mod to use for word(s)
func ChooseMod(word []rune, modValues Modificator) []rune {
	if modValues.class == "hex" || modValues.class == "bin" {
		dataToChange := converters.AtoiBase(word, modValues.class)
		return converters.ConvertIntToRune(dataToChange)
	} else if modValues.class == "up" {
		return words.ToUpper(word)
	} else if modValues.class == "low" {
		return words.ToLower(word)
	} else {
		return words.Cap(word)
	}
}

// Uses mod to word(s)
func UseMod(modValues Modificator, wordList *[][]rune) {
	wordPos := len(*wordList) - 1

	if modValues.specialNum == 0 {
		modValues.specialNum = modValues.specialNum + 1
	}

	for b := modValues.specialNum; b > 0; b-- {
		(*wordList)[wordPos] = ChooseMod((*wordList)[wordPos], modValues)
		wordPos = wordPos - 2
	}
}

// Formats punctual marks as per assignment
func FormatPunctList(wordList *[][]rune) {
	for i := 0; i < len(*wordList); i++ {
		var sortedPunctList []rune = nil
		notPunctElement := false

		for k := 0; k < len((*wordList)[i]); k++ {
			if (*wordList)[i][k] == 32 || (*wordList)[i][k] == 33 || (*wordList)[i][k] == 44 || (*wordList)[i][k] == 46 || (*wordList)[i][k] == 58 || (*wordList)[i][k] == 59 || (*wordList)[i][k] == 63 {
				continue
			} else {
				notPunctElement = true
				break
			}
		}

		if notPunctElement {
			continue
		}

		for k := 0; k < len((*wordList)[i]); k++ {
			if (*wordList)[i][k] == 32 {
				continue
			} else {
				sortedPunctList = append(sortedPunctList, (*wordList)[i][k])
			}
		}

		if i != len(*wordList)-1 {
			sortedPunctList = append(sortedPunctList, 32)
		}

		(*wordList)[i] = sortedPunctList
	}
}

// Formats article 'A' and 'a'
func FormatArticle(wordList *[][]rune) {
	for i := 0; i < len(*wordList); i++ {
		for k := 0; k < len((*wordList)[i]); k++ {
			if ((*wordList)[i][k] == 97 || (*wordList)[i][k] == 65) && len((*wordList)[i]) == 1 {
				(*wordList)[i] = append((*wordList)[i], 110)
			}
		}
	}
}

// Formats quotes as per assignment
func FormatQuotes(wordList *[][]rune) {
	firstQuoteFound := false
	secondQuoteFound := false
	firstQuotePos := -1
	secondQuotePos := -1

	for i := 0; i < len(*wordList); i++ {
		for k := 0; k < len((*wordList)[i]); k++ {
			if (*wordList)[i][k] == 39 {
				if !firstQuoteFound && firstQuotePos == -1 && len((*wordList)[i]) == 1 {
					firstQuoteFound = true
					firstQuotePos = i
				} else if firstQuoteFound && len((*wordList)[i-1]) == 1 && (*wordList)[i][0] == 39 {
					secondQuoteFound = true
					secondQuotePos = i
				}
			}

			if firstQuoteFound && secondQuoteFound {
				(*wordList)[firstQuotePos+1] = nil
				(*wordList)[secondQuotePos-1] = nil
				firstQuoteFound = false
				secondQuoteFound = false
				firstQuotePos = -1
				secondQuotePos = -1
			}
		}
	}
}

// Saves formated list to result.txt
func SaveFormatedText(wordList *[][]rune, fileResult string) {
	var wordListInBytes []rune

	for a := 0; a < len(*wordList); a++ {
		for b := 0; b < len((*wordList)[a]); b++ {
			wordListInBytes = append(wordListInBytes, (*wordList)[a][b])
		}
	}

	finalTextInBytes := converters.TranslateToByteSlice(wordListInBytes)

	errorTwo := ioutil.WriteFile(fileResult, finalTextInBytes, 0)
	CheckFile(errorTwo, fileResult)
}
