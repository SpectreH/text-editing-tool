package converter

import (
	"text-editing-tool/numbers"
)

func AtoiBase(numbersToConvert []rune, baseToDefine string) int {
	base := DefineBase(baseToDefine)

	positions := []rune(base)
	system := len(base)

	var finalResultInInt int = 0

	for i := 0; i < len(numbersToConvert); i++ {
		for k := 0; k < len(positions); k++ {
			if numbersToConvert[i] == positions[k] {
				systemInPower := numbers.Pow(system, ((len(numbersToConvert) - 1) - i))
				finalResultInInt = finalResultInInt + (k * systemInPower)
				break
			}
		}
	}

	return finalResultInInt
}

func DefineBase(base string) string {
	baseForCheck := []rune(base)
	finalBase := "null"

	if baseForCheck[0] == 104 {
		finalBase = "0123456789ABCDEF"
	} else {
		finalBase = "01"
	}

	return finalBase
}
