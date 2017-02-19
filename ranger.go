package ranger

// UpTo returns a slice of struct{} with len == max.
// It is intended for use in a for-range loop:
// 	for i := range ranger.UpTo(10) {
//		fmt.Println(i)
//	}
// This will print the numbers 0 to 9, one on each line
func UpTo(max int) []struct{} {
	return make([]struct{}, max)
}
