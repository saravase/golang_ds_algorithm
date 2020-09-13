package sorting

func BubbleSort(unSortArr []int) (sortArr []int) {

	length := len(unSortArr)
	for outIndex := 0; outIndex < length; outIndex++ {
		for inIndex := outIndex + 1; inIndex < length; inIndex++ {
			if unSortArr[outIndex] > unSortArr[inIndex] {
				unSortArr[outIndex], unSortArr[inIndex] = unSortArr[inIndex], unSortArr[outIndex]
			}
		}
	}

	return unSortArr
}
