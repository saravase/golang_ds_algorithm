package searching

func LinearSearch(sortArr []int, searchValue int) (index int) {

	for index := 0; index < len(sortArr); index++ {
		if sortArr[index] == searchValue {
			return index
		}
	}
	return -1
}
