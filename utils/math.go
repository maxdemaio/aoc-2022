package utils

func Abs(val int) int {
	if val < 0 {
		val = val * -1
	}
	return val
}
