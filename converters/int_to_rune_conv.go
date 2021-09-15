package converter

// // converts from int number to rune array number
func ConvertIntToRune(n int) []rune {
	x := []int{}

	var l int = n
	temLB := l
	tempLA := l

	var divideCounter int = 0

	for x := 0; x < 32; x++ {
		tempLA = tempLA / 10
		divideCounter = divideCounter + 1
		if tempLA == 0 {
			break
		}
	}

	for v := 0; v < divideCounter; v++ {
		var tempVar int = temLB % 10
		temLB = temLB / 10
		x = append(x, tempVar)
	}

	var result []rune

	for i := divideCounter - 1; 0 <= i; i-- {
		result = append(result, rune(x[i]+48))
	}

	return result
}
