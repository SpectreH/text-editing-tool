package converter

import "text-editing-tool/numbers"

func AtoiBase(s string, baseToDefine string) int {
	base := DefineBase(baseToDefine)

	positions := []rune(base)
	numbersToConvert := []rune(s)
	system := len(base)

	var finalResult int = 0

	for i := 0; i < len(numbersToConvert); i++ {
		for k := 0; k < len(positions); k++ {
			if numbersToConvert[i] == positions[k] {
				systemInPower := numbers.Pow(system, ((len(numbersToConvert) - 1) - i))
				finalResult = finalResult + (k * systemInPower)
				break
			}
		}
	}

	return finalResult
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
