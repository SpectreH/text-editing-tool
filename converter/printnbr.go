package converter

func PrintNbr(n int) []rune {
	// if n < 0 {
	// 	x := []int{}

	// 	var l int = n / -1

	// 	if n == -9223372036854775808 {
	// 		l = 9223372036854775807
	// 	}

	// 	temLB := l
	// 	tempLA := l

	// 	var divideCounter int = 0

	// 	for x := 0; x < 32; x++ {
	// 		tempLA = tempLA / 10
	// 		divideCounter = divideCounter + 1
	// 		if tempLA == 0 {
	// 			break
	// 		}
	// 	}

	// 	for v := 0; v < divideCounter; v++ {
	// 		var tempVar int = temLB % 10
	// 		temLB = temLB / 10
	// 		x = append(x, tempVar)
	// 	}

	// 	z01.PrintRune('-')

	// 	for i := divideCounter - 1; 0 <= i; i-- {
	// 		if i == 0 && n == -9223372036854775808 {
	// 			z01.PrintRune(rune(x[i] + 49))
	// 			break
	// 		}
	// 		z01.PrintRune(rune(x[i] + 48))
	// 	}
	// }

	// if n > 0 {
	// 	x := []int{}

	// 	var l int = n
	// 	temLB := l
	// 	tempLA := l

	// 	var divideCounter int = 0

	// 	for x := 0; x < 32; x++ {
	// 		tempLA = tempLA / 10
	// 		divideCounter = divideCounter + 1
	// 		if tempLA == 0 {
	// 			break
	// 		}
	// 	}

	// 	for v := 0; v < divideCounter; v++ {
	// 		var tempVar int = temLB % 10
	// 		temLB = temLB / 10
	// 		x = append(x, tempVar)
	// 	}

	// 	var result []rune

	// 	for i := divideCounter - 1; 0 <= i; i-- {
	// 		result = append(result, rune(x[i]))
	// 	}

	// 	return result
	// }

	// if n == 0 {
	// 	z01.PrintRune(rune(48))
	// }

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
