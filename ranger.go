package ranger

// UpTo returns a slice of struct{} with len == max.
// It is intended for use in a for-range loop:
// 	for i := range ranger.UpTo(10) {
//		fmt.Println(i)
//	}
// This will print the numbers 0 to 9, one on each line
// If max < 0, you'll get a nil slice (which is like having a slice of length 0)
func UpTo(max int) []struct{} {
	if max < 0 {
		return nil
	}
	return make([]struct{}, max)
}
