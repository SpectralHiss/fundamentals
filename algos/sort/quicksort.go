package sort

func QuickSort(slice []int) []int {
	if len(slice) == 1 {
		return slice
	}

	midPoint := slice[len(slice)/2]
	var left []int
	var right []int

	for _, elem := range slice {

		if elem < midPoint {
			left = append(left, elem)
		} else if elem == midPoint {
			right = prepend(right, elem)
		} else {
			right = append(right, elem)
		}

	}

	if len(left) == 0 {
		return QuickSort(right)
	}
	if len(right) == 0 {
		return QuickSort(left)
	}

	return append(QuickSort(left), QuickSort(right)...)
}

func prepend(slice []int, elem int) []int {
	if cap(slice) == 0 {
		return []int{elem}
	}
	tempSlice := make([]int, cap(slice)+1)
	tempSlice[0] = elem

	for index, elem := range slice[0:len(slice)] {
		tempSlice[index+1] = elem
	}
	//	tempSlice = append(tempSlice, slice[0:len(slice)-1]

	return tempSlice
}
