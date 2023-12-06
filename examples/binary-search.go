package examples

func BinarySearch(values []int, correctValue int) int {
	minorAttempt := 0
	maxAttempt := len(values) - 1
	for minorAttempt <= maxAttempt {
		mid := int((float64(minorAttempt) + float64(maxAttempt)) / 2)
		attempt := values[mid]

		if attempt == correctValue {
			return mid
		}

		if attempt > correctValue {
			maxAttempt = mid - 1
		} else {
			minorAttempt = mid + 1
		}
	}
	return -1
}
