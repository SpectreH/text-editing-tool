package converter

import "text-editing-tool/numbers"

func AtoiBase(s string, base string) int {
	if CheckBaseTwo(base) {
		return 0
	}

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

func CheckBaseTwo(baseForCheck string) bool {
	base := []rune(baseForCheck)
	var checkResult bool = false

	if len(baseForCheck) < 2 {
		checkResult = true
	}

	for i := 0; i < len(base); i++ {
		for k := 0; k < len(base); k++ {
			if base[i] == 45 || base[i] == 43 {
				checkResult = true
			} else if base[i] > base[k] || base[i] < base[k] {
				continue
			} else if i == k {
				continue
			} else if base[i] == base[k] {
				checkResult = true
			}
		}
	}

	return checkResult
}
