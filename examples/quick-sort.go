package examples

func QuickSort(values []int) []int {
	var pivotValue int
	var minorThantPivot, majorThanPivot []int

	if len(values) <= 1 {
		return values
	}

	pivotValue = values[len(values)/2]

	for _, value := range values {
		if value < pivotValue {
			minorThantPivot = append(minorThantPivot, value)
			continue
		}
		if value > pivotValue {
			majorThanPivot = append(majorThanPivot, value)
		}
	}

	compiledSort := append(QuickSort(minorThantPivot), pivotValue)
	compiledSort = append(compiledSort, QuickSort(majorThanPivot)...)
	return compiledSort
}
