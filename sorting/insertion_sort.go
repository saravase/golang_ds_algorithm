package sorting

func InsertionSort(unSortArr []int) (sortArr []int) {

	length := len(unSortArr)
	for outIndex := 1; outIndex < length; outIndex++ {
		currentVal := unSortArr[outIndex]
		inIndex := outIndex - 1
		for inIndex >= 0 && unSortArr[inIndex] > currentVal {
			unSortArr[inIndex+1] = unSortArr[inIndex]
			inIndex--
		}
		unSortArr[inIndex+1] = currentVal
	}

	return unSortArr
}
