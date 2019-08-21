package common

import "sort"

// DeduplicationArrayInt : int型の配列から重複を削除して返す
func DeduplicationArrayInt(array []int) []int {
	dedupMap := make(map[int]struct{})
	for _, elem := range array {
		dedupMap[elem] = struct{}{}
	}

	retArray := []int{}
	for key := range dedupMap {
		retArray = append(retArray, key)
	}

	sort.Ints(retArray)
	return retArray
}
