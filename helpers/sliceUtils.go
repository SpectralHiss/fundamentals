package helpers

func PrependNotGrow(slice []interface{}, elem interface{}) []interface{} {

	if cap(slice) == 0 {
		return []interface{}{elem}
	}
	tempSlice := make([]interface{}, cap(slice))
	tempSlice[0] = elem

	for index, elem := range slice[0 : len(slice)-1] {
		tempSlice[index+1] = elem
	}

	return tempSlice
}

func PrependAndGrow(slice []int, elem int) []int {
	if cap(slice) == 0 {
		return []int{elem}
	}
	tempSlice := make([]int, cap(slice)+1)
	tempSlice[0] = elem

	for index, elem := range slice[0:] {
		tempSlice[index+1] = elem
	}

	return tempSlice
}
