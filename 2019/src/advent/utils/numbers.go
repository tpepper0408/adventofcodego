package utils

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func FindMin(n []int) int {
	min := n[0]
	for _, value := range n {
		if value < min {
			min = value
		}
	}
	return min
}
