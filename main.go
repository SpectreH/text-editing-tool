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

	var currentWord []rune
	//var wordStartIndex []int

	fileSample := os.Args[1]
	// fileResult := os.Args[2]

	file, err := ioutil.ReadFile(fileSample)

	CheckFiles(err, fileSample)

	text := TranslateFileText(file)

	var writingWordDown bool
	var modValues Modificator

	for i := 0; i < len(text); i++ {
		if (64 < text[i] && text[i] < 91) || (96 < text[i] && text[i] < 123) {
			writingWordDown = true
		} else if text[i] == 32 {
			writingWordDown = false
		}

		if text[i] == 40 {
			modValues = CheckMod(i, text)

			if modValues.lenght != 0 {
				i = i + modValues.lenght
			} else {
				writingWordDown = true
			}
		}

		if writingWordDown {
			currentWord = append(currentWord, text[i])
		}
	}
}

func TranslateFileText(bytes []byte) []rune {
	var text []rune

	for i := range bytes {
		text = append(text, rune(bytes[i]))
	}

	return text
}

func CheckFiles(e error, fileName string) {
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
			fmt.Println("Mod Found")
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

	return checkResult
}
