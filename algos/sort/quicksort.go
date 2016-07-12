package sort

func QuickSort(slice []int) []int {
	if len(slice) == 1 {
		return slice
	}

	midIndex := len(slice) / 2
	midPoint := slice[midIndex]
	var left []int
	var right []int

	for index, elem := range slice {

		if elem < midPoint {
			left = append(left, elem)
		} else if elem == midPoint && index != midIndex {
			right = prependAndGrow(right, elem)
		} else if elem > midPoint {
			right = append(right, elem)
		}

	}

	if len(left) == 0 {
		return append([]int{midPoint}, QuickSort(right)...)
	}
	if len(right) == 0 {
		return append(QuickSort(left), midPoint)
	}

	left = append(QuickSort(left), midPoint)
	return append(left, QuickSort(right)...)
}

func prependAndGrow(slice []int, elem int) []int {
	if cap(slice) == 0 {
		return []int{elem}
	}
	tempSlice := make([]int, cap(slice)+1)
	tempSlice[0] = elem

	for index, elem := range slice[0:] {
		tempSlice[index+1] = elem
	}
	//	tempSlice = append(tempSlice, slice[0:len(slice)-1]

	return tempSlice
}
