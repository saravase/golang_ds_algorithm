package sorting

func SelectionSort(unSortArr []int) (sortArr []int) {

	length := len(unSortArr)
	for outIndex := 0; outIndex < length; outIndex++ {
		minValueIndex := outIndex
		for inIndex := outIndex; inIndex < length; inIndex++ {
			if unSortArr[minValueIndex] > unSortArr[inIndex] {
				minValueIndex = inIndex
			}
		}
		unSortArr[outIndex], unSortArr[minValueIndex] = unSortArr[minValueIndex], unSortArr[outIndex]
	}

	return unSortArr
}
