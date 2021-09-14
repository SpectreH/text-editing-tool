package numbers

func Pow(nb int, power int) int {
	if power == 0 {
		return 1
	} else if power < 0 {
		return 0
	} else if !(power < 63) {
		return 0
	} else {
		return (nb * Pow(nb, power-1))
	}
}
