package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text-editing-tool/converter"
	"text-editing-tool/words"
)

type Modificator struct {
	class      string
	lenght     int
	specialNum int
}

var MODSLIST [5]string

func main() {
	MODSLIST = [5]string{"hex", "bin", "up", "low", "cap"}

	if len(os.Args) != 3 {
		fmt.Println("Error! Missing required text files.")

		return
	}

	fileSample := os.Args[1]
	fileResult := os.Args[2]

	file, err := ioutil.ReadFile(fileSample)
	CheckFile(err, fileSample)

	text := TranslateFileToRune(file)

	var wordList [][]rune
	var currentWord []rune
	var punMarkList []rune

	var modValues Modificator
	var modFound bool = false

	for i := 0; i < len(text); i++ {
		if text[i] != 32 && text[i] != 33 && text[i] != 44 && text[i] != 46 && text[i] != 58 && text[i] != 59 && text[i] != 63 {
			currentWord = append(currentWord, text[i])
		} else {

			for furtherIndex := i + 1; furtherIndex < len(text); furtherIndex++ {
				if text[furtherIndex] == 40 {
					modValues = CheckMod(i+1, text)

					if modValues.lenght != 0 {
						delIndex := i
						for charsToDel := modValues.lenght; charsToDel >= 0; charsToDel-- {
							text = RemoveIndex(text, delIndex)
						}

						modFound = true
						break
					} else {
						break
					}
				}
			}

			if currentWord != nil {
				AddWord(&wordList, currentWord)
			}

			if modFound {
				modFound = false
				wordPos := len(wordList) - 1
				for b := modValues.specialNum; b >= 0; b-- {
					wordList[wordPos] = ModWord(wordList[wordPos], modValues)
					wordPos = wordPos - 2
				}
			}

			for m := i; m < len(text); m++ {
				if text[m] == 32 || text[m] == 33 || text[m] == 44 || text[m] == 46 || text[m] == 58 || text[m] == 59 || text[m] == 63 {
					punMarkList = append(punMarkList, text[m])
				} else {
					i = m - 1
					break
				}
			}

			AddWord(&wordList, punMarkList)
			currentWord = nil
			punMarkList = nil

		}
	}

	if currentWord != nil {
		AddWord(&wordList, currentWord)
	}

	for i := 1; i < len(wordList); i = i + 2 {
		var sortedPunctList []rune = nil

		for k := 0; k < len(wordList[i]); k++ {
			if wordList[i][k] == 32 {
				continue
			} else {
				sortedPunctList = append(sortedPunctList, wordList[i][k])
			}
		}

		if i != len(wordList)-1 {
			sortedPunctList = append(sortedPunctList, 32)
		}

		wordList[i] = sortedPunctList
	}

	for i := 0; i < len(wordList); i = i + 2 {
		for k := 0; k < len(wordList[i]); k++ {
			if (wordList[i][k] == 97 || wordList[i][k] == 65) && len(wordList[i]) == 1 {
				wordList[i] = append(wordList[i], 110)
			} else {
				break
			}
		}
	}

	firstQouteFound := false
	secondQouteFound := false

	firstQoutePos := -1
	secondQoutePos := -1

	fmt.Println(wordList)

	for i := 0; i < len(wordList); i++ {
		for k := 0; k < len(wordList[i]); k++ {
			if wordList[i][k] == 39 && len(wordList[i]) == 1 {
				if !firstQouteFound && firstQoutePos == -1 {
					if wordList[i+1][0] == 32 {
						firstQouteFound = true
						firstQoutePos = i
					}
				} else if firstQouteFound {
					if wordList[i-1][0] == 32 {
						secondQouteFound = true
						secondQoutePos = i
						break
					}
				}
			}
		}

		if firstQouteFound && secondQouteFound {
			wordList[firstQoutePos+1] = nil
			wordList[secondQoutePos-1] = nil
			break
		}
	}

	fmt.Println(wordList)

	var test []rune

	for a := 0; a < len(wordList); a++ {
		for b := 0; b < len(wordList[a]); b++ {
			test = append(test, wordList[a][b])
		}
	}

	finalTextInBytes := TranslateToBytes(test)

	errorTwo := ioutil.WriteFile(fileResult, finalTextInBytes, 0)
	CheckFile(errorTwo, fileResult)
}

func TranslateFileToRune(bytes []byte) []rune {
	var text []rune

	for i := range bytes {
		text = append(text, rune(bytes[i]))
	}

	return text
}

func CheckFile(e error, fileName string) {
	if e != nil {
		// TODO Error message
		// errorMessage := "ERROR: open " + fileName + ": no such file or directory\n"
		// PrintBytes([]byte(errorMessage))
		os.Exit(1)
	}
}

func CheckMod(modStartingIndex int, text []rune) Modificator {
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
			//fmt.Println("MOD FOUND")
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
				// TODO Error message
				os.Exit(1)
			}
		}
	} else if text[index] == 41 {
		checkResult.specialNum = 0
	} else {
		// TODO Error message
		os.Exit(1)
	}

	checkResult.specialNum = converter.BasicAtoi(specialNumber)
	checkResult.lenght = checkResult.lenght + 2 + len(specialNumber)

	return checkResult
}

func TranslateToBytes(runes []rune) []byte {
	var text []byte

	for i := range runes {
		text = append(text, byte(runes[i]))
	}

	return text
}

func RemoveIndex(s []rune, index int) []rune {
	return append(s[:index], s[index+1:]...)
}

func AddWord(wordList *[][]rune, wordToAdd []rune) {
	*wordList = append(*wordList, nil)
	for i := len(*wordList) - 1; i < len(*wordList); i++ {
		for k := 0; k < len(wordToAdd); k++ {
			(*wordList)[i] = append((*wordList)[i], wordToAdd[k])
		}
	}
}

func ModWord(word []rune, modValues Modificator) []rune {

	if modValues.class == "hex" || modValues.class == "bin" {
		dataToChange := converter.AtoiBase(word, modValues.class)
		return converter.PrintNbr(dataToChange)
	} else if modValues.class == "up" {
		return words.ToUpper(word)
	} else if modValues.class == "low" {
		return words.ToLower(word)
	} else {
		return words.Cap(word)
	}

}
