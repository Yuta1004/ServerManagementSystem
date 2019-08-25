package common

// ConvToInterfaceSlice : スライスを[]interfaceに変換する
func ConvToInterfaceSlice(sSlice []string) []interface{} {
	retSlice := make([]interface{}, len(sSlice))
	for idx, item := range sSlice {
		retSlice[idx] = item
	}
	return retSlice
}