package utils

func AbsDiffInt(x int, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
