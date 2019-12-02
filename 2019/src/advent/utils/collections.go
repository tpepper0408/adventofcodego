package utils

func DeepCopyInt(original []int) []int {
	copy := make([]int, len(original))
	for i, originalValue := range original {
		copy[i] = originalValue
	}
	return copy
}
