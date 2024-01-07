package arrayhashing

func ContainsDuplicates(arr []int8) bool {
	elmentwithcount := make(map[int8]int8)
	for _, v := range arr {
		_, ok := elmentwithcount[v]
		if ok {
			return true
		} else {
			elmentwithcount[v] = 1
		}
	}
	return false
}
