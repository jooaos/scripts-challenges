package examples

func RecursiveSum(values []int) int {
	if len(values) <= 1 {
		return values[0]
	}
	total := values[0] + RecursiveSum(values[0+1:])
	return total
}
