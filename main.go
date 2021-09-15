package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text-editing-tool/converter"
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

	var wordsList [][]rune
	var currentWord []rune

	fileSample := os.Args[1]
	fileResult := os.Args[2]

	file, err := ioutil.ReadFile(fileSample)
	CheckFile(err, fileSample)

	text := TranslateFileToRune(file)

	var modValues Modificator
	// var textStart bool = true
	var wordWrited bool = false
	var saveIndex int = 0

	for i := 0; i < len(text); i++ {
		if (64 < text[i] && text[i] < 91) || (96 < text[i] && text[i] < 123) {
			wordWrited = false
			currentWord = append(currentWord, text[i])
		} else {
			saveIndex = i
			for furtherIndex := i - 1; furtherIndex < len(text); furtherIndex++ {
				if text[furtherIndex] == 40 {
					modValues = CheckMod(i, text)

					if modValues.lenght != 0 {
						if modValues.specialNum != 0 {
							modValues.lenght = modValues.lenght + 2
						}

						delIndex := i

						for charsToDel := modValues.lenght; charsToDel > 0; charsToDel-- {
							text = RemoveIndex(text, delIndex)
						}
						break
					}

				} else if (64 < text[furtherIndex] && text[furtherIndex] < 91) || (96 < text[furtherIndex] && text[furtherIndex] < 123) {
					wordWrited = true
					break
				}
			}

			wordToAdd(&wordsList, currentWord)
			currentWord = nil
		}

		if wordWrited {
			for furtherIndex := saveIndex; furtherIndex < len(text); furtherIndex++ {
				if !(64 < text[furtherIndex] && text[furtherIndex] < 91) || (96 < text[furtherIndex] && text[furtherIndex] < 123) {
					currentWord = append(currentWord, text[furtherIndex])
				} else {
					i = furtherIndex - 1
					wordToAdd(&wordsList, currentWord)
					currentWord = nil
					break
				}
			}
		}
	}

	finalTextInBytes := TranslateToBytes(text)

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
			fmt.Println("MOD FOUND")
			firstCheckResult.class = MODSLIST[k]
			firstCheckResult.lenght = len(MODSLIST[k])
			break
		}
	}

	if !modFound {
		return firstCheckResult
	}

	finalCheckResult := CheckModSpecNumber(modStartingIndex, text, firstCheckResult)

	return finalCheckResult
}

func CheckModSpecNumber(modStartingIndex int, text []rune, checkResult Modificator) Modificator {
	var specialNumber []rune
	index := modStartingIndex + checkResult.lenght + 1

	if text[index] == 44 && text[index+1] == 32 {

		for m := index + 2; m < len(text); m++ {
			if 47 < text[m] && text[m] < 58 {
				specialNumber = append(specialNumber, text[m])
			} else if text[m] == 41 {
				break
			} else {
				// TODO Error message
				checkResult.class = "NULL"
				checkResult.lenght = 0
				checkResult.specialNum = 0
				return checkResult
			}
		}

	} else if text[index] == 41 {
		checkResult.specialNum = 0
	} else {
		// TODO Error message
		checkResult.class = "NULL"
		checkResult.lenght = 0
		checkResult.specialNum = 0
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

func wordToAdd(wordsList *[][]rune, wordToAdd []rune) {
	*wordsList = append(*wordsList, nil)
	for i := len(*wordsList) - 1; i < len(*wordsList); i++ {
		for k := 0; k < len(wordToAdd); k++ {
			(*wordsList)[i] = append((*wordsList)[i], wordToAdd[k])
		}
	}
}
