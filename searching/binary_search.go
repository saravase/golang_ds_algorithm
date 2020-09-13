package searching

func BinarySearch(sortArr []int, searchValue int) (index int) {
	min := 0
	max := len(sortArr)
	for index := min; index < max; index++ {
		mid := (max + min) / 2
		if sortArr[mid] == searchValue {
			return mid
		} else if sortArr[mid] > searchValue {
			max = mid
		} else {
			min = mid
		}
	}
	return -1
}
