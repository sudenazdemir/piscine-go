package piscine

func RecursivePower(nb int, power int) int {
	result := nb
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		return result * RecursivePower(nb, power-1)
	}
}
